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
	mode                 = flag.String("mode", "faucet", "")
	addrToSeed           = flag.String("seed-addr", "0x809E16D8fa815839cF453FB4A935860e06a499c8", "")
	dataHashesReaderAddr = flag.String("helper-contract-addr", "0x6006f425ac71535452a5e23BCF73E26A115f281d", "")
	privHex              = "45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8"
	addr                 = "http://localhost:8545"
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
	case "faucet":
		seedFunds(ctx, client, sender, eipSigner)
	case "deploy":
		deployContracts(ctx, client, sender, eipSigner)
	case "submit":
		sendBlobTx(ctx, client, sender, dankSigner)
	default:
		panic("no valid mode set")
	}
}

func deployContracts(ctx context.Context, client *ethclient.Client, sender common.Address, signerFn bind.SignerFn) {
	key, err := crypto.HexToECDSA(priv)
	if err != nil {
		panic(err)
	}
	sequencerAddr, tx, instance, err := bindings.DeployToySequencerInbox(
		&bind.TransactOpts{Signer: signerFn},
		client,
		common.HexToAddress(*dataHashesReaderAddr),
	)
	if err != nil {
		panic(err)
	}
}

func sendBlobTx(ctx context.Context, client *ethclient.Client, sender common.Address, signerFn bind.SignerFn) {
	nonce, err := client.PendingNonceAt(ctx, crypto.PubkeyToAddress(key.PublicKey))
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

	to := common.HexToAddress(*dataHashesReaderAddr)
	calldata, _ := hexutil.Decode("0xe83a2d820000000000000000000000000000000000000000000000000000000000000000")
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
	log.Printf("Transaction submitted. hash=%v", tx.Hash())
}

func seedFunds(ctx context.Context, client *ethclient.Client, sender common.Address, signerFn bind.SignerFn) {
	nonce, err := client.PendingNonceAt(ctx, sender)
	if err != nil {
		panic(err)
	}
	value := big.NewInt(5000000000000000000) // in wei (5 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}
	toAddress := common.HexToAddress(*addrToSeed)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := signerFn(sender, tx)
	if err != nil {
		panic(err)
	}
	if err = client.SendTransaction(ctx, signedTx); err != nil {
		panic(err)
	}
}
