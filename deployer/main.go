package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	priv := "45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8"
	addr := "http://localhost:8545"

	chainId := big.NewInt(1)
	signer := types.NewEIP155Signer(chainId)

	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, addr)
	if err != nil {
		panic(err)
	}

	key, err := crypto.HexToECDSA(priv)
	if err != nil {
		panic(err)
	}
	nonce, err := client.PendingNonceAt(ctx, crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		panic(err)
	}

	value := big.NewInt(5000000000000000000) // in wei (5 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	toAddress := common.HexToAddress("0x93e188d4cdD4C661d22DBA6Bb7Bd445BA0fCc039")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, signer, key)
	if err != nil {
		panic(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}
}
