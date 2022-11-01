package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/Inphi/eip4844-interop/shared"
	"github.com/Inphi/eip4844-interop/upload/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"github.com/protolambda/ztyp/view"
)

var (
	mode          = flag.String("mode", "faucet", "")
	sequencerAddr = flag.String("sequencer-addr", "", "")
	privHex       = "45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8"
	addr          = "http://localhost:8545"
)

func main() {
	flag.Parse()
	chainId := big.NewInt(1)
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, addr)
	if err != nil {
		panic(err)
	}
	privKey, err := crypto.HexToECDSA(privHex)
	if err != nil {
		panic(err)
	}
	sender := crypto.PubkeyToAddress(privKey.PublicKey)

	eipSigner := func(sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(chainId)
		return types.SignTx(tx, signer, privKey)
	}
	dankSigner := func(sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewDankSigner(chainId)
		return types.SignTx(tx, signer, privKey)
	}
	switch *mode {
	case "deploy":
		deployContracts(ctx, client, sender, eipSigner)
	case "read-events":
		readEvents(ctx, client)
	case "submit":
		sendBlobTx(ctx, client, sender, dankSigner)
	default:
		panic("no valid mode set")
	}
}

func deployContracts(ctx context.Context, client *ethclient.Client, sender common.Address, signerFn bind.SignerFn) {
	gp, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}
	nonce, err := client.PendingNonceAt(ctx, sender)
	if err != nil {
		panic(err)
	}
	opts := &bind.TransactOpts{
		From: sender, Signer: signerFn, GasPrice: gp, Value: big.NewInt(0), Nonce: big.NewInt(int64(nonce)),
	}
	dataHashesReaderAddr, tx1, _, err := bindings.DeployDataHashesReader(opts, client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deployed at addr %#x and sent tx1: %+v\n", dataHashesReaderAddr, tx1)
	nonce, err = client.PendingNonceAt(ctx, sender)
	if err != nil {
		panic(err)
	}
	opts.Nonce = big.NewInt(int64(nonce))
	deployedSequencerAddr, tx2, _, err := bindings.DeployToySequencerInbox(
		opts,
		client,
		dataHashesReaderAddr,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deployed at addr %#x and sent tx2: %+v\n", deployedSequencerAddr, tx2)
}

func readEvents(ctx context.Context, client *ethclient.Client) {
	filterer, err := bindings.NewToySequencerInboxFilterer(common.HexToAddress(*sequencerAddr), client)
	if err != nil {
		panic(err)
	}
	num, err := client.BlockNumber(ctx)
	if err != nil {
		panic(err)
	}
	it, err := filterer.FilterDataHashesParsed(&bind.FilterOpts{
		Start: 0,
		End:   &num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Initial %+v\n", it.Event)
	for it.Next() {
		fmt.Printf("Event %+v\n", it.Event)
	}
}

func sendBlobTx(ctx context.Context, client *ethclient.Client, sender common.Address, signerFn bind.SignerFn) {
	nonce, err := client.PendingNonceAt(ctx, sender)
	if err != nil {
		log.Fatalf("Error getting nonce: %v", err)
	}
	items := []byte("foobar")
	blobs := shared.EncodeBlobs(items)
	commitments, versionedHashes, aggregatedProof, err := blobs.ComputeCommitmentsAndAggregatedProof()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Num versioned hashes %d\n", len(versionedHashes))

	to := common.HexToAddress(*sequencerAddr)
	calldata, _ := hexutil.Decode("0x461153830000000000000000000000000000000000000000000000000000000000000000")
	txData := types.SignedBlobTx{
		Message: types.BlobTxMessage{
			ChainID:             view.Uint256View(*uint256.NewInt(1)),
			Nonce:               view.Uint64View(nonce),
			Gas:                 210000,
			GasFeeCap:           view.Uint256View(*uint256.NewInt(5000000000)),
			GasTipCap:           view.Uint256View(*uint256.NewInt(5000000000)),
			MaxFeePerDataGas:    view.Uint256View(*uint256.NewInt(300000000000)), // needs to be at least the min fee
			Value:               view.Uint256View(*uint256.NewInt(0)),
			Data:                types.TxDataView(calldata),
			To:                  types.AddressOptionalSSZ{Address: (*types.AddressSSZ)(&to)},
			BlobVersionedHashes: versionedHashes,
		},
	}

	wrapData := types.BlobTxWrapData{
		BlobKzgs:           commitments,
		Blobs:              blobs,
		KzgAggregatedProof: aggregatedProof,
	}
	tx := types.NewTx(&txData, types.WithTxWrapData(&wrapData))
	signedTx, err := signerFn(sender, tx)
	if err != nil {
		panic(err)
	}
	if err = client.SendTransaction(ctx, signedTx); err != nil {
		panic(err)
	}
	log.Printf("Blob tx submitted %v", tx.Hash())
}
