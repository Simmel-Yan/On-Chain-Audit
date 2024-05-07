/*
generate an Ethereum keystore by two means:
1) generate random private key (see func generateRandomKeyStore())
2) generate according to a given privatekey (see func main())

In this project we use the second method, in which the privatekey belongs to the first account in Ganache desktop
According to the given private key, func ks.ImportECDSA() generates a new keystore file (saved in directory "wallets").

The content of keystore file "UTC-xxx" are copied to deployVRF.go and deployTestVRF.go (line 46) , to generate transaction signer "auth"
*/

package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func generateRandomKeyStore() {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "123456"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

func main() {
	// generate a keystore object
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	// load private key (the input key is the first account's private key from Ganache desktop)
	privateKey, err := crypto.HexToECDSA("d421a02a518b5915d5bed4e86d202cc7b72ecc59d045f57ed567d466ec47bd53")
	password := "111111"
	if err != nil {
		log.Fatal(err)
	}

	// generate keystore file according to the given private key
	ks2, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Keystore file generated for address: %s\nURL: %s\n", ks2.Address, ks2.URL)
}
