// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VRF

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VrfMetaData contains all meta data concerning the Vrf contract.
var VrfMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101f5610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c806303a507be146100715780635727dc5c1461008f5780637a308a4c146100ad5780639870e2a4146100cb578063997da8d4146100e9578063eeeac01e14610107575b600080fd5b610079610125565b6040518082815260200191505060405180910390f35b610097610149565b6040518082815260200191505060405180910390f35b6100b561014e565b6040518082815260200191505060405180910390f35b6100d3610172565b6040518082815260200191505060405180910390f35b6100f1610196565b6040518082815260200191505060405180910390f35b61010f61019b565b6040518082815260200191505060405180910390f35b7f79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f8179881565b600781565b7f483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b881565b7ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036414181565b600081565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f8156fea26469706673582212200edddb4e3fbc85d837ba3b965e9ae92a257fe1237bd68aa5f04a267c725de41764736f6c634300060c0033",
}

// VrfABI is the input ABI used to generate the binding from.
// Deprecated: Use VrfMetaData.ABI instead.
var VrfABI = VrfMetaData.ABI

// VrfBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VrfMetaData.Bin instead.
var VrfBin = VrfMetaData.Bin

// DeployVrf deploys a new Ethereum contract, binding an instance of Vrf to it.
func DeployVrf(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vrf, error) {
	parsed, err := VrfMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VrfBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vrf{VrfCaller: VrfCaller{contract: contract}, VrfTransactor: VrfTransactor{contract: contract}, VrfFilterer: VrfFilterer{contract: contract}}, nil
}

// Vrf is an auto generated Go binding around an Ethereum contract.
type Vrf struct {
	VrfCaller     // Read-only binding to the contract
	VrfTransactor // Write-only binding to the contract
	VrfFilterer   // Log filterer for contract events
}

// VrfCaller is an auto generated read-only Go binding around an Ethereum contract.
type VrfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VrfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VrfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VrfSession struct {
	Contract     *Vrf              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VrfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VrfCallerSession struct {
	Contract *VrfCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VrfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VrfTransactorSession struct {
	Contract     *VrfTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VrfRaw is an auto generated low-level Go binding around an Ethereum contract.
type VrfRaw struct {
	Contract *Vrf // Generic contract binding to access the raw methods on
}

// VrfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VrfCallerRaw struct {
	Contract *VrfCaller // Generic read-only contract binding to access the raw methods on
}

// VrfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VrfTransactorRaw struct {
	Contract *VrfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVrf creates a new instance of Vrf, bound to a specific deployed contract.
func NewVrf(address common.Address, backend bind.ContractBackend) (*Vrf, error) {
	contract, err := bindVrf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vrf{VrfCaller: VrfCaller{contract: contract}, VrfTransactor: VrfTransactor{contract: contract}, VrfFilterer: VrfFilterer{contract: contract}}, nil
}

// NewVrfCaller creates a new read-only instance of Vrf, bound to a specific deployed contract.
func NewVrfCaller(address common.Address, caller bind.ContractCaller) (*VrfCaller, error) {
	contract, err := bindVrf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VrfCaller{contract: contract}, nil
}

// NewVrfTransactor creates a new write-only instance of Vrf, bound to a specific deployed contract.
func NewVrfTransactor(address common.Address, transactor bind.ContractTransactor) (*VrfTransactor, error) {
	contract, err := bindVrf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VrfTransactor{contract: contract}, nil
}

// NewVrfFilterer creates a new log filterer instance of Vrf, bound to a specific deployed contract.
func NewVrfFilterer(address common.Address, filterer bind.ContractFilterer) (*VrfFilterer, error) {
	contract, err := bindVrf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VrfFilterer{contract: contract}, nil
}

// bindVrf binds a generic wrapper to an already deployed contract.
func bindVrf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VrfMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vrf *VrfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vrf.Contract.VrfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vrf *VrfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrf.Contract.VrfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vrf *VrfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vrf.Contract.VrfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vrf *VrfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vrf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vrf *VrfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vrf *VrfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vrf.Contract.contract.Transact(opts, method, params...)
}

// AA is a free data retrieval call binding the contract method 0x997da8d4.
//
// Solidity: function AA() view returns(uint256)
func (_Vrf *VrfCaller) AA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "AA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AA is a free data retrieval call binding the contract method 0x997da8d4.
//
// Solidity: function AA() view returns(uint256)
func (_Vrf *VrfSession) AA() (*big.Int, error) {
	return _Vrf.Contract.AA(&_Vrf.CallOpts)
}

// AA is a free data retrieval call binding the contract method 0x997da8d4.
//
// Solidity: function AA() view returns(uint256)
func (_Vrf *VrfCallerSession) AA() (*big.Int, error) {
	return _Vrf.Contract.AA(&_Vrf.CallOpts)
}

// BB is a free data retrieval call binding the contract method 0x5727dc5c.
//
// Solidity: function BB() view returns(uint256)
func (_Vrf *VrfCaller) BB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "BB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BB is a free data retrieval call binding the contract method 0x5727dc5c.
//
// Solidity: function BB() view returns(uint256)
func (_Vrf *VrfSession) BB() (*big.Int, error) {
	return _Vrf.Contract.BB(&_Vrf.CallOpts)
}

// BB is a free data retrieval call binding the contract method 0x5727dc5c.
//
// Solidity: function BB() view returns(uint256)
func (_Vrf *VrfCallerSession) BB() (*big.Int, error) {
	return _Vrf.Contract.BB(&_Vrf.CallOpts)
}

// GX is a free data retrieval call binding the contract method 0x03a507be.
//
// Solidity: function GX() view returns(uint256)
func (_Vrf *VrfCaller) GX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "GX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GX is a free data retrieval call binding the contract method 0x03a507be.
//
// Solidity: function GX() view returns(uint256)
func (_Vrf *VrfSession) GX() (*big.Int, error) {
	return _Vrf.Contract.GX(&_Vrf.CallOpts)
}

// GX is a free data retrieval call binding the contract method 0x03a507be.
//
// Solidity: function GX() view returns(uint256)
func (_Vrf *VrfCallerSession) GX() (*big.Int, error) {
	return _Vrf.Contract.GX(&_Vrf.CallOpts)
}

// GY is a free data retrieval call binding the contract method 0x7a308a4c.
//
// Solidity: function GY() view returns(uint256)
func (_Vrf *VrfCaller) GY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "GY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GY is a free data retrieval call binding the contract method 0x7a308a4c.
//
// Solidity: function GY() view returns(uint256)
func (_Vrf *VrfSession) GY() (*big.Int, error) {
	return _Vrf.Contract.GY(&_Vrf.CallOpts)
}

// GY is a free data retrieval call binding the contract method 0x7a308a4c.
//
// Solidity: function GY() view returns(uint256)
func (_Vrf *VrfCallerSession) GY() (*big.Int, error) {
	return _Vrf.Contract.GY(&_Vrf.CallOpts)
}

// NN is a free data retrieval call binding the contract method 0x9870e2a4.
//
// Solidity: function NN() view returns(uint256)
func (_Vrf *VrfCaller) NN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "NN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NN is a free data retrieval call binding the contract method 0x9870e2a4.
//
// Solidity: function NN() view returns(uint256)
func (_Vrf *VrfSession) NN() (*big.Int, error) {
	return _Vrf.Contract.NN(&_Vrf.CallOpts)
}

// NN is a free data retrieval call binding the contract method 0x9870e2a4.
//
// Solidity: function NN() view returns(uint256)
func (_Vrf *VrfCallerSession) NN() (*big.Int, error) {
	return _Vrf.Contract.NN(&_Vrf.CallOpts)
}

// PP is a free data retrieval call binding the contract method 0xeeeac01e.
//
// Solidity: function PP() view returns(uint256)
func (_Vrf *VrfCaller) PP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "PP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PP is a free data retrieval call binding the contract method 0xeeeac01e.
//
// Solidity: function PP() view returns(uint256)
func (_Vrf *VrfSession) PP() (*big.Int, error) {
	return _Vrf.Contract.PP(&_Vrf.CallOpts)
}

// PP is a free data retrieval call binding the contract method 0xeeeac01e.
//
// Solidity: function PP() view returns(uint256)
func (_Vrf *VrfCallerSession) PP() (*big.Int, error) {
	return _Vrf.Contract.PP(&_Vrf.CallOpts)
}
