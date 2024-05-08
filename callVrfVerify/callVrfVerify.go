package main

import (
	"callEthereum/contract_testVRF"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Case struct {
	Pkx    string `json:"pkx"`
	Pky    string `json:"pky"`
	Gammax string `json:"gammax"`
	Gammay string `json:"gammay"`
	C      string `json:"c"`
	S      string `json:"s"`
	Msg    string `json:"msg"`
}

func readCases(fileName string) ([]Case, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err2 := io.ReadAll(jsonFile)
	if err2 != nil {
		return nil, err2
	}

	var cases = make([]Case, 0)
	err3 := json.Unmarshal(byteValue, &cases)
	if err3 != nil {
		return cases, err3
	}

	return cases, nil
}

func main() {
	// connect to Ganache RPC server
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Panic("failed to Dail", err)
	}
	fmt.Println("we have a connection")

	// obtain account address (we use the first account in Ganache)
	accountAddress := common.HexToAddress("0xc9f0f90f52b635c5fb4ede9e439a28bdea2c587c")
	if err != nil {
		log.Panic(err)
	}

	// creat a contract object
	address := common.HexToAddress("0x244974cD3F614038B7240905518f0d943F922a07") // contract address
	instance, err := testVRF.NewTestVRF(address, client)
	if err != nil {
		log.Fatal(err)
	}

	var cases, _ = readCases("./input.json")
	for _, ca := range cases {
		// generate publicKey
		publicX := new(big.Int)
		publicX.SetString(ca.Pkx, 16)
		publicY := new(big.Int)
		publicY.SetString(ca.Pky, 16)
		publicKey := [2]*big.Int{publicX, publicY}

		// generate vrf proof pi
		gammaX := new(big.Int)
		gammaX.SetString(ca.Gammax, 16)
		gammaY := new(big.Int)
		gammaY.SetString(ca.Gammay, 16)
		c := new(big.Int)
		c.SetString(ca.C, 16)
		s := new(big.Int)
		s.SetString(ca.S, 16)
		pi := [4]*big.Int{gammaX, gammaY, c, s}

		message, _ := hex.DecodeString(ca.Msg)

		// call the contract
		res, err := instance.Verify(&bind.CallOpts{
			Pending:     false,
			From:        accountAddress,
			BlockNumber: nil,
			Context:     nil,
		}, publicKey, pi, message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("verify result=", res)
	}
}
