Operating system: macOS Sonoma 14.4.1

## Generate Keystore File
A [keystore](https://goethereumbook.org/keystore/) is a file recording encrypted Ethereum wallet private key. You should run `./keystore/keystore.go` to generate a keystore file, which will be saved to `./keystore/wallets`.
The keystore will be useful when we deploy contracts.

## Compile Contracts
use [remix online IDE](https://remix.ethereum.org/) to compile contracts `./contracts/VRF.sol` and `./contracts/testVRF.sol`, you'll get `VRF.abi`, `VRF.bin` (we save them in `./contract_VRF/`), and `testVRF.abi`, `testVRF.bin` (saved in `./contract_testVRF/`).

Then open terminal and generate go files using following commands:
```
> cd contract_VRF
> abigen --bin=VRF.bin --abi=VRF.abi --pkg=vrf --out=VRF.go
> cd ../contract_testVRF
> abigen --bin=testVRF.bin --abi=testVRF.abi --pkg=testVRF --out=testVRF.go  
```
These commands generate `VRF.go` and `testVRF.go`, which are used to deploy and call contracts.

## Deploy Contracts

## Call Contracts
