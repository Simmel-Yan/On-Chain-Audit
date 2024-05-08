package main

import (
	"On-Chain-Audit/contract_testVRF"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func main() {
	// connect to ethclient
	client, err := ethclient.Dial("http://localhost:7545") // method 1: connect to Ganache RPC server
	//client, err := ethclient.Dial("https://rinkeby.infura.io") // method 2: connect to infura
	if err != nil {
		log.Fatal(err)
	}

	// load private key (the first account's private key from Ganache)
	privateKey, err := crypto.HexToECDSA("d421a02a518b5915d5bed4e86d202cc7b72ecc59d045f57ed567d466ec47bd53")
	if err != nil {
		log.Fatal(err)
	}

	// generate public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// generate publickey address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)                 // generate address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress) // nonce for transaction
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("public key: %s\n", fromAddress)

	// generate auth (keyinfo is pasted from keystore file ../keystore/wallets/UTC--2024-04-26T07-47-55.353681000Z--c9f0f90f52b635c5fb4ede9e439a28bdea2c587c)
	var keyinfo = `{"address":"c9f0f90f52b635c5fb4ede9e439a28bdea2c587c","crypto":{"cipher":"aes-128-ctr","ciphertext":"7bd99ece7ef284c3ead1b06aa122054e5d69ade4c955398e03068520b34d70e0","cipherparams":{"iv":"3c084d54c74317d5df907c12f2d2cef3"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"25f854357fc16efe300115743b6bdef76d379d7f410975b316f5e0bf9be899ba"},"mac":"4fc00cfb7f9ae8098dc1ab5826f3ea8c6cb17f89696f7954496edc3573f028b6"},"id":"2ef7d432-0102-4012-a515-ce162d3623cc","version":3}`
	keyin := strings.NewReader(keyinfo)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Panic(err)
	}

	auth, err := bind.NewTransactorWithChainID(keyin, "111111", chainID) // the passphrase should be identical to "password" in keystore.go
	if err != nil {
		log.Panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	fmt.Printf("auth: %s\n", auth)

	// deploy contract testVRF
	address, tx, instance, err := testVRF.DeployTestVRF(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("contract testVRF address: %s, ", address.Hex())
	fmt.Printf("tx hash: %s\n", tx.Hash().Hex())
	_ = instance // contract instance
}
