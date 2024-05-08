Operating system: macOS Sonoma 14.4.1

## Generate Keystore File
A [keystore](https://goethereumbook.org/keystore/) is a file recording encrypted Ethereum wallet private key. You should run `./keystore/keystore.go` to generate a keystore file, which will be saved to `./keystore/wallets`.
The keystore will be useful when we deploy contracts.

## Compile Contracts
To verify the vrf output on-chain, we have two contracts: contract VRF (saved in `./contracts/VRF.sol`) and contract testVRF (saved in `./contracts/testVRF.sol`). The latter one calls the functions in the former one, and the Golang client only calls the latter contract during practical use.

use [remix online IDE](https://remix.ethereum.org/) to compile these two contracts, you'll get `VRF.abi`, `VRF.bin` (we save them in `./contract_VRF/`), and `testVRF.abi`, `testVRF.bin` (saved in `./contract_testVRF/`).

Then open terminal and generate go files using following commands:
```
> cd contract_VRF
> abigen --bin=VRF.bin --abi=VRF.abi --pkg=vrf --out=VRF.go
> cd ../contract_testVRF
> abigen --bin=testVRF.bin --abi=testVRF.abi --pkg=testVRF --out=testVRF.go  
```
These commands generate `VRF.go` and `testVRF.go`, which are used to deploy and call contracts.

## Deploy Contracts
Before deploying the contracts, make sure that you have initiates the Ethereum test chain. In this project we use [Ganache](https://archive.trufflesuite.com/ganache/), because it is easy to use. The default port of Ganache is 7545, which is useful when we deploy a contract. (See `./deployVRF/deployVRF.go` and `./deployTestVRF/deployTestVRF.go` line 18.)

In this part, the generated keystore is used, to. (See `deployVRF.go` and `deployTestVRF.go` line 46.) 

To deploy contract VRF and testVRF, run `deployVRF.go` and `deployTestVRF.go`, as follows. 
```
> cd deployVRF
> go run deployVRF.go
> cd ../deployTestVRF
> go run deployTestVRF.go
```
We can check the deployed contracts on Ganache desktop, and the contract address will be used when we call the contract. 

## Call Contracts
