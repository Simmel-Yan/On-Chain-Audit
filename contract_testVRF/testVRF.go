// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testVRF

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

// TestVRFMetaData contains all meta data concerning the TestVRF contract.
var TestVRFMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_publicKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_proof\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"computeFastVerifyParams\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_point\",\"type\":\"bytes\"}],\"name\":\"decodePoint\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"decodeProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_publicKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_proof\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"_uPoint\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_vComponents\",\"type\":\"uint256[4]\"}],\"name\":\"fastVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gammaX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gammaY\",\"type\":\"uint256\"}],\"name\":\"gammaToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_publicKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_proof\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061265a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806345899df41461006757806351cdea5f1461023d5780635f09834114610334578063bbe95b4c14610380578063c95ac47a146104d3578063c9d4ef1f1461067b575b600080fd5b61022560048036036101a081101561007e57600080fd5b8101908080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019064010000000081111561011d57600080fd5b82018360208201111561012f57600080fd5b8035906020019184600183028401116401000000008311171561015157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f8201169050808301925050505050509192919290505050610772565b60405180821515815260200191505060405180910390f35b6102f66004803603602081101561025357600080fd5b810190808035906020019064010000000081111561027057600080fd5b82018360208201111561028257600080fd5b803590602001918460018302840111640100000000831117156102a457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061078c565b6040518082600260200280838360005b83811015610321578082015181840152602081019050610306565b5050505090500191505060405180910390f35b61036a6004803603604081101561034a57600080fd5b8101908080359060200190929190803590602001909291905050506107a4565b6040518082815260200191505060405180910390f35b6104bb600480360360e081101561039657600080fd5b8101908080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019064010000000081111561043557600080fd5b82018360208201111561044757600080fd5b8035906020019184600183028401116401000000008311171561046957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107b8565b60405180821515815260200191505060405180910390f35b61060e600480360360e08110156104e957600080fd5b8101908080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019064010000000081111561058857600080fd5b82018360208201111561059a57600080fd5b803590602001918460018302840111640100000000831117156105bc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107ce565b6040518083600260200280838360005b8381101561063957808201518184015260208101905061061e565b5050505090500182600460200280838360005b8381101561066757808201518184015260208101905061064c565b505050509050019250505060405180910390f35b6107346004803603602081101561069157600080fd5b81019080803590602001906401000000008111156106ae57600080fd5b8201836020820111156106c057600080fd5b803590602001918460018302840111640100000000831117156106e257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107f5565b6040518082600460200280838360005b8381101561075f578082015181840152602081019050610744565b5050505090500191505060405180910390f35b6000610781868686868661080d565b905095945050505050565b6107946125be565b61079d82610a63565b9050919050565b60006107b08383610b20565b905092915050565b60006107c5848484610c5f565b90509392505050565b6107d66125be565b6107de6125e0565b6107e9858585610dd3565b91509150935093915050565b6107fd6125e0565b61080682610f3f565b9050919050565b600080600061081c888761102d565b915091506108948760036004811061083057fe5b60200201518860026004811061084257fe5b60200201518a60006002811061085457fe5b60200201518b60016002811061086657fe5b60200201518960006002811061087857fe5b60200201518a60016002811061088a57fe5b6020020151611346565b15806108dd57506108db876003600481106108ab57fe5b60200201518383876000600481106108bf57fe5b6020020151886001600481106108d157fe5b6020020151611512565b155b806109475750610945876002600481106108f357fe5b60200201518860006004811061090557fe5b60200201518960016004811061091757fe5b60200201518760026004811061092957fe5b60200201518860036004811061093b57fe5b6020020151611512565b155b1561095757600092505050610a5a565b6000806109cd8660006004811061096a57fe5b60200201518760016004811061097c57fe5b60200201518860026004811061098e57fe5b6020020151896003600481106109a057fe5b602002015160007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f61160e565b915091506000610a2785858c6000600481106109e557fe5b60200201518d6001600481106109f757fe5b60200201518c600060028110610a0957fe5b60200201518d600160028110610a1b57fe5b60200201518989611642565b905089600260048110610a3657fe5b60200201518160801c6fffffffffffffffffffffffffffffffff1614955050505050505b95945050505050565b610a6b6125be565b6021825114610ae2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f4d616c666f726d656420636f6d7072657373656420454320706f696e7400000081525060200191505060405180910390fd5b60008060018401519150602184015190506000610aff83836118b5565b90506040518060400160405280838152602001828152509350505050919050565b6000606060fe6003610b3286866118ee565b604051602001808460ff1660f81b81526001018360ff1660f81b815260010182805190602001908083835b60208310610b805780518252602082019150602081019050602083039250610b5d565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405290506002816040518082805190602001908083835b60208310610bee5780518252602082019150602081019050602083039250610bcb565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015610c30573d6000803e3d6000fd5b5050506040513d6020811015610c4557600080fd5b810190808051906020019092919050505091505092915050565b6000806000610c6e868561102d565b91509150600080610d0787600360048110610c8557fe5b60200201517f79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f817987f483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b88a600260048110610cd957fe5b60200201518c600060028110610ceb57fe5b60200201518d600160028110610cfd57fe5b6020020151611939565b91509150600080610d6089600360048110610d1e57fe5b602002015187878c600260048110610d3257fe5b60200201518d600060048110610d4457fe5b60200201518e600160048110610d5657fe5b6020020151611939565b915091506000610d9887878c600060048110610d7857fe5b60200201518d600160048110610d8a57fe5b602002015189898989611642565b905089600260048110610da757fe5b60200201518160801c6fffffffffffffffffffffffffffffffff16149750505050505050509392505050565b610ddb6125be565b610de36125e0565b600080610df0878661102d565b91509150600080610e8988600360048110610e0757fe5b60200201517f79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f817987f483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b88b600260048110610e5b57fe5b60200201518d600060028110610e6d57fe5b60200201518e600160028110610e7f57fe5b6020020151611939565b91509150600080610eac8a600360048110610ea057fe5b602002015187876119ad565b91509150600080610ef18c600260048110610ec357fe5b60200201518d600060048110610ed557fe5b60200201518e600160048110610ee757fe5b60200201516119ad565b91509150604051806040016040528087815260200186815250604051806080016040528086815260200185815260200184815260200183815250995099505050505050505050935093915050565b610f476125e0565b6051825114610fbe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f4d616c666f726d6564205652462070726f6f660000000000000000000000000081525060200191505060405180910390fd5b600080600080600186015193506021860151925060318601519150605186015190506000610fec85856118b5565b90506040518060800160405280858152602001828152602001846fffffffffffffffffffffffffffffffff1681526020018381525095505050505050919050565b600080606060fe60016110628760006002811061104657fe5b60200201518860016002811061105857fe5b60200201516118ee565b86604051602001808560ff1660f81b81526001018460ff1660f81b815260010183805190602001908083835b602083106110b1578051825260208201915060208101905060208303925061108e565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b6020831061110257805182526020820191506020810190506020830392506110df565b6001836020036101000a038019825116818451168082178552505050505050905001945050505050604051602081830303815290604052905060005b6101008160ff1610156112d0576000600283836040516020018083805190602001908083835b602083106111875780518252602082019150602081019050602083039250611164565b6001836020036101000a0380198251168184511680821785525050505050509050018260ff1660f81b8152600101925050506040516020818303038152906040526040518082805190602001908083835b602083106111fb57805182526020820191506020810190506020830392506111d8565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa15801561123d573d6000803e3d6000fd5b5050506040513d602081101561125257600080fd5b8101908080519060200190929190505050905060008160001c9050600061127a6002836118b5565b90506112ab8282600060077ffffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f6119ea565b156112c057818196509650505050505061133f565b505050808060010191505061113e565b506040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f4e6f2076616c696420706f696e742077617320666f756e64000000000000000081525060200191505060405180910390fd5b9250929050565b6000807ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141887ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141038161139457fe5b0690507ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141806113bf57fe5b868209905060007ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141887ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141038161141157fe5b069050600060018360001b600060028a8161142857fe5b06141561143657601b611439565b601c5b8a60001b7ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03641418061146557fe5b8c870960001b60405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156114bf573d6000803e3d6000fd5b5050506020604051035190508073ffffffffffffffffffffffffffffffffffffffff166114ec8787611a8e565b73ffffffffffffffffffffffffffffffffffffffff161493505050509695505050505050565b60008060016000806002888161152457fe5b06141561153257601b611535565b601c5b8860001b7ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03641418061156157fe5b8a8c0960001b60405160008152602001604052604051808560001b81526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156115be573d6000803e3d6000fd5b5050506020604051035190508073ffffffffffffffffffffffffffffffffffffffff166115eb8585611a8e565b73ffffffffffffffffffffffffffffffffffffffff161491505095945050505050565b60008060008061161f888887611ae0565b915091506116318a8a84848a8a611afc565b935093505050965096945050505050565b6000606060fe60026116548c8c6118ee565b61165e8b8b6118ee565b6116688a8a6118ee565b61167289896118ee565b604051602001808760ff1660f81b81526001018660ff1660f81b815260010185805190602001908083835b602083106116c0578051825260208201915060208101905060208303925061169d565b6001836020036101000a03801982511681845116808217855250505050505090500184805190602001908083835b6020831061171157805182526020820191506020810190506020830392506116ee565b6001836020036101000a03801982511681845116808217855250505050505090500183805190602001908083835b60208310611762578051825260208201915060208101905060208303925061173f565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b602083106117b35780518252602082019150602081019050602083039250611790565b6001836020036101000a0380198251168184511680821785525050505050509050019650505050505050604051602081830303815290604052905060006002826040518082805190602001908083835b602083106118265780518252602082019150602081019050602083039250611803565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015611868573d6000803e3d6000fd5b5050506040513d602081101561187d57600080fd5b810190808051906020019092919050505090506000604051826000820152600081015191505080935050505098975050505050505050565b60006118e68383600060077ffffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f611b8a565b905092915050565b60606000600283816118fc57fe5b0660020190508084604051602001808360ff1660f81b81526001018281526020019250505060405160208183030381529060405291505092915050565b60008060008061194a8a8a8a6119ad565b9150915060008061195c8989896119ad565b915091506000806119928686868660007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f61160e565b91509150818197509750505050505050965096945050505050565b6000806119de85858560007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f611c79565b91509150935093915050565b600085600014806119fb5750818610155b80611a065750846000145b80611a115750818510155b15611a1f5760009050611a85565b60008280611a2957fe5b868709905060008380611a3857fe5b888580611a4157fe5b8a8b0909905060008614611a67578380611a5757fe5b8480611a5f57fe5b878a09820890505b60008514611a7d578380611a7757fe5b85820890505b808214925050505b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff838360405160200180838152602001828152602001925050506040516020818303038152906040528051906020012060001c16905092915050565b600080848385850381611aef57fe5b0691509150935093915050565b6000806000806000888b1415611b4d5760008680611b1657fe5b898c081415611b2e5760008094509450505050611b7f565b611b3c8b8b60018a8a611cb3565b809350819450829550505050611b6b565b611b5e8b8b60018c8c60018c611dd2565b8093508194508295505050505b611b77838383896122a3565b945094505050505b965096945050505050565b600060028660ff161480611ba1575060038660ff16145b611bf6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806126036022913960400191505060405180910390fd5b60008280611c0057fe5b8380611c0857fe5b858580611c1157fe5b888a09088480611c1d57fe5b8580611c2557fe5b898a098909089050611c458160046001860181611c3e57fe5b04856122fe565b905060008060028960ff16840181611c5957fe5b0614611c6757818403611c69565b815b9050809250505095945050505050565b6000806000806000611c908a8a8a60018b8b61242c565b925092509250611ca2838383896122a3565b945094505050509550959350505050565b600080600080861415611cce57878787925092509250611dc7565b60008480611cd857fe5b898a09905060008580611ce757fe5b898a09905060008680611cf657fe5b898a09905060008780611d0557fe5b8880611d0d57fe5b848e09600409905060008880611d1f57fe5b8980611d2757fe5b8a80611d2f57fe5b8586098c098a80611d3c57fe5b876003090890508880611d4b57fe5b8980611d5357fe5b8384088a038a80611d6057fe5b8384090894508880611d6e57fe5b8980611d7657fe5b8a80611d7e57fe5b8687096008098a038a80611d8e57fe5b8b80611d9657fe5b888d03860884090893508880611da857fe5b8980611db057fe5b8c8e09600209925084848497509750975050505050505b955095509592505050565b6000806000808a148015611de65750600089145b15611df957868686925092509250612296565b600087148015611e095750600086145b15611e1c57898989925092509250612296565b611e246125e0565b8480611e2c57fe5b898a0981600060048110611e3c57fe5b6020020181815250508480611e4d57fe5b81600060048110611e5a57fe5b60200201518a0981600160048110611e6e57fe5b6020020181815250508480611e7f57fe5b86870981600260048110611e8f57fe5b6020020181815250508480611ea057fe5b81600260048110611ead57fe5b6020020151870981600360048110611ec157fe5b60200201818152505060405180608001604052808680611edd57fe5b83600260048110611eea57fe5b60200201518e0981526020018680611efe57fe5b83600360048110611f0b57fe5b60200201518d0981526020018680611f1f57fe5b83600060048110611f2c57fe5b60200201518b0981526020018680611f4057fe5b83600160048110611f4d57fe5b60200201518a09815250905080600260048110611f6657fe5b602002015181600060048110611f7857fe5b6020020151141580611fac575080600360048110611f9257fe5b602002015181600160048110611fa457fe5b602002015114155b61201e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f557365206a6163446f75626c652066756e6374696f6e20696e7374656164000081525060200191505060405180910390fd5b6120266125e0565b858061202e57fe5b8260006004811061203b57fe5b602002015187038360026004811061204f57fe5b6020020151088160006004811061206257fe5b602002018181525050858061207357fe5b8260016004811061208057fe5b602002015187038360036004811061209457fe5b602002015108816001600481106120a757fe5b60200201818152505085806120b857fe5b816000600481106120c557fe5b6020020151826000600481106120d757fe5b602002015109816002600481106120ea57fe5b60200201818152505085806120fb57fe5b8160006004811061210857fe5b60200201518260026004811061211a57fe5b6020020151098160036004811061212d57fe5b6020020181815250506000868061214057fe5b8260036004811061214d57fe5b60200201518803888061215c57fe5b8460016004811061216957fe5b60200201518560016004811061217b57fe5b602002015109089050868061218c57fe5b878061219457fe5b888061219c57fe5b846002600481106121a957fe5b6020020151866000600481106121bb57fe5b602002015109600209880382089050600087806121d457fe5b88806121dc57fe5b838a038a806121e757fe5b866002600481106121f457fe5b60200201518860006004811061220657fe5b602002015109088460016004811061221a57fe5b6020020151099050878061222a57fe5b888061223257fe5b8460036004811061223f57fe5b60200201518660016004811061225157fe5b6020020151098903820890506000888061226757fe5b898061226f57fe5b8b8f098560006004811061227f57fe5b602002015109905082828297509750975050505050505b9750975097945050505050565b60008060006122b285856124ca565b9050600084806122be57fe5b8283099050600085806122cd57fe5b828a099050600086806122dc57fe5b87806122e457fe5b8486098a0990508181955095505050505094509492505050565b600080821415612376576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4d6f64756c7573206973207a65726f000000000000000000000000000000000081525060200191505060405180910390fd5b60008414156123885760009050612425565b600083141561239a5760019050612425565b60006001905060007f800000000000000000000000000000000000000000000000000000000000000090505b600081111561241f57838186161515870a85848509099150836002820486161515870a85848509099150836004820486161515870a85848509099150836008820486161515870a858485090991506010810490506123c6565b81925050505b9392505050565b600080600080891415612447578787879250925092506124be565b60008990506000806000600190505b600084146124b05760006001851614612485576124788383838f8f8f8e611dd2565b8093508194508295505050505b6002848161248f57fe5b04935061249f8c8c8c8c8c611cb3565b809c50819d50829e50505050612456565b828282965096509650505050505b96509650969350505050565b60008083141580156124dc5750818314155b80156124e9575060008214155b61255b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f496e76616c6964206e756d62657200000000000000000000000000000000000081525060200191505060405180910390fd5b60008060019050600084905060005b600087146125b15786828161257b57fe5b04905082868061258757fe5b878061258f57fe5b858409880386088094508195505050868782028303809850819350505061256a565b8394505050505092915050565b6040518060400160405280600290602082028036833780820191505090505090565b604051806080016040528060049060208202803683378082019150509050509056fe496e76616c696420636f6d7072657373656420454320706f696e7420707265666978a2646970667358221220b367b882d286f4d16ba91637d13833864226f6244f7ea0bbd67caaf69350f5ea64736f6c634300060c0033",
}

// TestVRFABI is the input ABI used to generate the binding from.
// Deprecated: Use TestVRFMetaData.ABI instead.
var TestVRFABI = TestVRFMetaData.ABI

// TestVRFBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestVRFMetaData.Bin instead.
var TestVRFBin = TestVRFMetaData.Bin

// DeployTestVRF deploys a new Ethereum contract, binding an instance of TestVRF to it.
func DeployTestVRF(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestVRF, error) {
	parsed, err := TestVRFMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestVRFBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestVRF{TestVRFCaller: TestVRFCaller{contract: contract}, TestVRFTransactor: TestVRFTransactor{contract: contract}, TestVRFFilterer: TestVRFFilterer{contract: contract}}, nil
}

// TestVRF is an auto generated Go binding around an Ethereum contract.
type TestVRF struct {
	TestVRFCaller     // Read-only binding to the contract
	TestVRFTransactor // Write-only binding to the contract
	TestVRFFilterer   // Log filterer for contract events
}

// TestVRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestVRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestVRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestVRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestVRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestVRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestVRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestVRFSession struct {
	Contract     *TestVRF          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestVRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestVRFCallerSession struct {
	Contract *TestVRFCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TestVRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestVRFTransactorSession struct {
	Contract     *TestVRFTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestVRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestVRFRaw struct {
	Contract *TestVRF // Generic contract binding to access the raw methods on
}

// TestVRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestVRFCallerRaw struct {
	Contract *TestVRFCaller // Generic read-only contract binding to access the raw methods on
}

// TestVRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestVRFTransactorRaw struct {
	Contract *TestVRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestVRF creates a new instance of TestVRF, bound to a specific deployed contract.
func NewTestVRF(address common.Address, backend bind.ContractBackend) (*TestVRF, error) {
	contract, err := bindTestVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestVRF{TestVRFCaller: TestVRFCaller{contract: contract}, TestVRFTransactor: TestVRFTransactor{contract: contract}, TestVRFFilterer: TestVRFFilterer{contract: contract}}, nil
}

// NewTestVRFCaller creates a new read-only instance of TestVRF, bound to a specific deployed contract.
func NewTestVRFCaller(address common.Address, caller bind.ContractCaller) (*TestVRFCaller, error) {
	contract, err := bindTestVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestVRFCaller{contract: contract}, nil
}

// NewTestVRFTransactor creates a new write-only instance of TestVRF, bound to a specific deployed contract.
func NewTestVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*TestVRFTransactor, error) {
	contract, err := bindTestVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestVRFTransactor{contract: contract}, nil
}

// NewTestVRFFilterer creates a new log filterer instance of TestVRF, bound to a specific deployed contract.
func NewTestVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*TestVRFFilterer, error) {
	contract, err := bindTestVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestVRFFilterer{contract: contract}, nil
}

// bindTestVRF binds a generic wrapper to an already deployed contract.
func bindTestVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestVRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestVRF *TestVRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestVRF.Contract.TestVRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestVRF *TestVRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestVRF.Contract.TestVRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestVRF *TestVRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestVRF.Contract.TestVRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestVRF *TestVRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestVRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestVRF *TestVRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestVRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestVRF *TestVRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestVRF.Contract.contract.Transact(opts, method, params...)
}

// ComputeFastVerifyParams is a free data retrieval call binding the contract method 0xc95ac47a.
//
// Solidity: function computeFastVerifyParams(uint256[2] _publicKey, uint256[4] _proof, bytes _message) pure returns(uint256[2], uint256[4])
func (_TestVRF *TestVRFCaller) ComputeFastVerifyParams(opts *bind.CallOpts, _publicKey [2]*big.Int, _proof [4]*big.Int, _message []byte) ([2]*big.Int, [4]*big.Int, error) {
	var out []interface{}
	err := _TestVRF.contract.Call(opts, &out, "computeFastVerifyParams", _publicKey, _proof, _message)

	if err != nil {
		return *new([2]*big.Int), *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)
	out1 := *abi.ConvertType(out[1], new([4]*big.Int)).(*[4]*big.Int)

	return out0, out1, err

}

// ComputeFastVerifyParams is a free data retrieval call binding the contract method 0xc95ac47a.
//
// Solidity: function computeFastVerifyParams(uint256[2] _publicKey, uint256[4] _proof, bytes _message) pure returns(uint256[2], uint256[4])
func (_TestVRF *TestVRFSession) ComputeFastVerifyParams(_publicKey [2]*big.Int, _proof [4]*big.Int, _message []byte) ([2]*big.Int, [4]*big.Int, error) {
	return _TestVRF.Contract.ComputeFastVerifyParams(&_TestVRF.CallOpts, _publicKey, _proof, _message)
}

// ComputeFastVerifyParams is a free data retrieval call binding the contract method 0xc95ac47a.
//
// Solidity: function computeFastVerifyParams(uint256[2] _publicKey, uint256[4] _proof, bytes _message) pure returns(uint256[2], uint256[4])
func (_TestVRF *TestVRFCallerSession) ComputeFastVerifyParams(_publicKey [2]*big.Int, _proof [4]*big.Int, _message []byte) ([2]*big.Int, [4]*big.Int, error) {
	return _TestVRF.Contract.ComputeFastVerifyParams(&_TestVRF.CallOpts, _publicKey, _proof, _message)
}

// DecodePoint is a free data retrieval call binding the contract method 0x51cdea5f.
//
// Solidity: function decodePoint(bytes _point) pure returns(uint256[2])
func (_TestVRF *TestVRFCaller) DecodePoint(opts *bind.CallOpts, _point []byte) ([2]*big.Int, error) {
	var out []interface{}
	err := _TestVRF.contract.Call(opts, &out, "decodePoint", _point)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// DecodePoint is a free data retrieval call binding the contract method 0x51cdea5f.
//
// Solidity: function decodePoint(bytes _point) pure returns(uint256[2])
func (_TestVRF *TestVRFSession) DecodePoint(_point []byte) ([2]*big.Int, error) {
	return _TestVRF.Contract.DecodePoint(&_TestVRF.CallOpts, _point)
}

// DecodePoint is a free data retrieval call binding the contract method 0x51cdea5f.
//
// Solidity: function decodePoint(bytes _point) pure returns(uint256[2])
func (_TestVRF *TestVRFCallerSession) DecodePoint(_point []byte) ([2]*big.Int, error) {
	return _TestVRF.Contract.DecodePoint(&_TestVRF.CallOpts, _point)
}

// DecodeProof is a free data retrieval call binding the contract method 0xc9d4ef1f.
//
// Solidity: function decodeProof(bytes _proof) pure returns(uint256[4])
func (_TestVRF *TestVRFCaller) DecodeProof(opts *bind.CallOpts, _proof []byte) ([4]*big.Int, error) {
	var out []interface{}
	err := _TestVRF.contract.Call(opts, &out, "decodeProof", _proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// DecodeProof is a free data retrieval call binding the contract method 0xc9d4ef1f.
//
// Solidity: function decodeProof(bytes _proof) pure returns(uint256[4])
func (_TestVRF *TestVRFSession) DecodeProof(_proof []byte) ([4]*big.Int, error) {
	return _TestVRF.Contract.DecodeProof(&_TestVRF.CallOpts, _proof)
}

// DecodeProof is a free data retrieval call binding the contract method 0xc9d4ef1f.
//
// Solidity: function decodeProof(bytes _proof) pure returns(uint256[4])
func (_TestVRF *TestVRFCallerSession) DecodeProof(_proof []byte) ([4]*big.Int, error) {
	return _TestVRF.Contract.DecodeProof(&_TestVRF.CallOpts, _proof)
}

// FastVerify is a free data retrieval call binding the contract method 0x45899df4.
//
// Solidity: function fastVerify(uint256[2] _publicKey, uint256[4] _proof, bytes _message, uint256[2] _uPoint, uint256[4] _vComponents) pure returns(bool)
func (_TestVRF *TestVRFCaller) FastVerify(opts *bind.CallOpts, _publicKey [2]*big.Int, _proof [4]*big.Int, _message []byte, _uPoint [2]*big.Int, _vComponents [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _TestVRF.contract.Call(opts, &out, "fastVerify", _publicKey, _proof, _message, _uPoint, _vComponents)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FastVerify is a free data retrieval call binding the contract method 0x45899df4.
//
// Solidity: function fastVerify(uint256[2] _publicKey, uint256[4] _proof, bytes _message, uint256[2] _uPoint, uint256[4] _vComponents) pure returns(bool)
func (_TestVRF *TestVRFSession) FastVerify(_publicKey [2]*big.Int, _proof [4]*big.Int, _message []byte, _uPoint [2]*big.Int, _vComponents [4]*big.Int) (bool, error) {
	return _TestVRF.Contract.FastVerify(&_TestVRF.CallOpts, _publicKey, _proof, _message, _uPoint, _vComponents)
}

// FastVerify is a free data retrieval call binding the contract method 0x45899df4.
//
// Solidity: function fastVerify(uint256[2] _publicKey, uint256[4] _proof, bytes _message, uint256[2] _uPoint, uint256[4] _vComponents) pure returns(bool)
func (_TestVRF *TestVRFCallerSession) FastVerify(_publicKey [2]*big.Int, _proof [4]*big.Int, _message []byte, _uPoint [2]*big.Int, _vComponents [4]*big.Int) (bool, error) {
	return _TestVRF.Contract.FastVerify(&_TestVRF.CallOpts, _publicKey, _proof, _message, _uPoint, _vComponents)
}

// GammaToHash is a free data retrieval call binding the contract method 0x5f098341.
//
// Solidity: function gammaToHash(uint256 _gammaX, uint256 _gammaY) pure returns(bytes32)
func (_TestVRF *TestVRFCaller) GammaToHash(opts *bind.CallOpts, _gammaX *big.Int, _gammaY *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TestVRF.contract.Call(opts, &out, "gammaToHash", _gammaX, _gammaY)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GammaToHash is a free data retrieval call binding the contract method 0x5f098341.
//
// Solidity: function gammaToHash(uint256 _gammaX, uint256 _gammaY) pure returns(bytes32)
func (_TestVRF *TestVRFSession) GammaToHash(_gammaX *big.Int, _gammaY *big.Int) ([32]byte, error) {
	return _TestVRF.Contract.GammaToHash(&_TestVRF.CallOpts, _gammaX, _gammaY)
}

// GammaToHash is a free data retrieval call binding the contract method 0x5f098341.
//
// Solidity: function gammaToHash(uint256 _gammaX, uint256 _gammaY) pure returns(bytes32)
func (_TestVRF *TestVRFCallerSession) GammaToHash(_gammaX *big.Int, _gammaY *big.Int) ([32]byte, error) {
	return _TestVRF.Contract.GammaToHash(&_TestVRF.CallOpts, _gammaX, _gammaY)
}

// Verify is a free data retrieval call binding the contract method 0xbbe95b4c.
//
// Solidity: function verify(uint256[2] _publicKey, uint256[4] _proof, bytes _message) pure returns(bool)
func (_TestVRF *TestVRFCaller) Verify(opts *bind.CallOpts, _publicKey [2]*big.Int, _proof [4]*big.Int, _message []byte) (bool, error) {
	var out []interface{}
	err := _TestVRF.contract.Call(opts, &out, "verify", _publicKey, _proof, _message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xbbe95b4c.
//
// Solidity: function verify(uint256[2] _publicKey, uint256[4] _proof, bytes _message) pure returns(bool)
func (_TestVRF *TestVRFSession) Verify(_publicKey [2]*big.Int, _proof [4]*big.Int, _message []byte) (bool, error) {
	return _TestVRF.Contract.Verify(&_TestVRF.CallOpts, _publicKey, _proof, _message)
}

// Verify is a free data retrieval call binding the contract method 0xbbe95b4c.
//
// Solidity: function verify(uint256[2] _publicKey, uint256[4] _proof, bytes _message) pure returns(bool)
func (_TestVRF *TestVRFCallerSession) Verify(_publicKey [2]*big.Int, _proof [4]*big.Int, _message []byte) (bool, error) {
	return _TestVRF.Contract.Verify(&_TestVRF.CallOpts, _publicKey, _proof, _message)
}
