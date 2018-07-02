// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AnyContractABI is the input ABI used to generate the binding from.
const AnyContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"texts\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"uint256\"},{\"name\":\"_text\",\"type\":\"string\"}],\"name\":\"batchWrite\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"numbers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_text\",\"type\":\"string\"}],\"name\":\"write\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AnyContractBin is the compiled bytecode used for deploying new contracts.
const AnyContractBin = `0x608060405234801561001057600080fd5b50610474806100206000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630e0f9cbb811461007c57806312065fe01461011f578063771602f714610146578063abc40f4114610163578063e349fe3a146101b6578063ebaac771146101e4575b600080fd5b34801561008857600080fd5b506100aa73ffffffffffffffffffffffffffffffffffffffff6004351661023d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e45781810151838201526020016100cc565b50505050905090810190601f1680156101115780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561012b57600080fd5b506101346102d7565b60408051918252519081900360200190f35b34801561015257600080fd5b506101616004356024356102f3565b005b604080516020600460443581810135601f810184900484028501840190955284845261016194823594602480359536959460649492019190819084018382808284375094975061031e9650505050505050565b3480156101c257600080fd5b5061013473ffffffffffffffffffffffffffffffffffffffff60043516610364565b3480156101f057600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101619436949293602493928401919081908401838280828437509497506103769650505050505050565b60016020818152600092835260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156102cf5780601f106102a4576101008083540402835291602001916102cf565b820191906000526020600020905b8154815290600101906020018083116102b257829003601f168201915b505050505081565b73ffffffffffffffffffffffffffffffffffffffff3016315b90565b73ffffffffffffffffffffffffffffffffffffffff3316600090815260208190526040902091019055565b73ffffffffffffffffffffffffffffffffffffffff3316600090815260208181526040808320868601905560018252909120825161035e928401906103b0565b50505050565b60006020819052908152604090205481565b73ffffffffffffffffffffffffffffffffffffffff3316600090815260016020908152604090912082516103ac928401906103b0565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106103f157805160ff191683800117855561041e565b8280016001018555821561041e579182015b8281111561041e578251825591602001919060010190610403565b5061042a92915061042e565b5090565b6102f091905b8082111561042a57600081556001016104345600a165627a7a7230582064fbac8530bd6a0bd60161dc3e9c0e06faf947beb7815100885441a88f0bb8ef0029`

// DeployAnyContract deploys a new Ethereum contract, binding an instance of AnyContract to it.
func DeployAnyContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AnyContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnyContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AnyContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AnyContract{AnyContractCaller: AnyContractCaller{contract: contract}, AnyContractTransactor: AnyContractTransactor{contract: contract}, AnyContractFilterer: AnyContractFilterer{contract: contract}}, nil
}

// AnyContract is an auto generated Go binding around an Ethereum contract.
type AnyContract struct {
	AnyContractCaller     // Read-only binding to the contract
	AnyContractTransactor // Write-only binding to the contract
	AnyContractFilterer   // Log filterer for contract events
}

// AnyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnyContractSession struct {
	Contract     *AnyContract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnyContractCallerSession struct {
	Contract *AnyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AnyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnyContractTransactorSession struct {
	Contract     *AnyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AnyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnyContractRaw struct {
	Contract *AnyContract // Generic contract binding to access the raw methods on
}

// AnyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnyContractCallerRaw struct {
	Contract *AnyContractCaller // Generic read-only contract binding to access the raw methods on
}

// AnyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnyContractTransactorRaw struct {
	Contract *AnyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnyContract creates a new instance of AnyContract, bound to a specific deployed contract.
func NewAnyContract(address common.Address, backend bind.ContractBackend) (*AnyContract, error) {
	contract, err := bindAnyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AnyContract{AnyContractCaller: AnyContractCaller{contract: contract}, AnyContractTransactor: AnyContractTransactor{contract: contract}, AnyContractFilterer: AnyContractFilterer{contract: contract}}, nil
}

// NewAnyContractCaller creates a new read-only instance of AnyContract, bound to a specific deployed contract.
func NewAnyContractCaller(address common.Address, caller bind.ContractCaller) (*AnyContractCaller, error) {
	contract, err := bindAnyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnyContractCaller{contract: contract}, nil
}

// NewAnyContractTransactor creates a new write-only instance of AnyContract, bound to a specific deployed contract.
func NewAnyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AnyContractTransactor, error) {
	contract, err := bindAnyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnyContractTransactor{contract: contract}, nil
}

// NewAnyContractFilterer creates a new log filterer instance of AnyContract, bound to a specific deployed contract.
func NewAnyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AnyContractFilterer, error) {
	contract, err := bindAnyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnyContractFilterer{contract: contract}, nil
}

// bindAnyContract binds a generic wrapper to an already deployed contract.
func bindAnyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnyContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnyContract *AnyContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AnyContract.Contract.AnyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnyContract *AnyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnyContract.Contract.AnyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnyContract *AnyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnyContract.Contract.AnyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnyContract *AnyContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AnyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnyContract *AnyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnyContract *AnyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnyContract.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_AnyContract *AnyContractCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AnyContract.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_AnyContract *AnyContractSession) GetBalance() (*big.Int, error) {
	return _AnyContract.Contract.GetBalance(&_AnyContract.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_AnyContract *AnyContractCallerSession) GetBalance() (*big.Int, error) {
	return _AnyContract.Contract.GetBalance(&_AnyContract.CallOpts)
}

// Numbers is a free data retrieval call binding the contract method 0xe349fe3a.
//
// Solidity: function numbers( address) constant returns(uint256)
func (_AnyContract *AnyContractCaller) Numbers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AnyContract.contract.Call(opts, out, "numbers", arg0)
	return *ret0, err
}

// Numbers is a free data retrieval call binding the contract method 0xe349fe3a.
//
// Solidity: function numbers( address) constant returns(uint256)
func (_AnyContract *AnyContractSession) Numbers(arg0 common.Address) (*big.Int, error) {
	return _AnyContract.Contract.Numbers(&_AnyContract.CallOpts, arg0)
}

// Numbers is a free data retrieval call binding the contract method 0xe349fe3a.
//
// Solidity: function numbers( address) constant returns(uint256)
func (_AnyContract *AnyContractCallerSession) Numbers(arg0 common.Address) (*big.Int, error) {
	return _AnyContract.Contract.Numbers(&_AnyContract.CallOpts, arg0)
}

// Texts is a free data retrieval call binding the contract method 0x0e0f9cbb.
//
// Solidity: function texts( address) constant returns(string)
func (_AnyContract *AnyContractCaller) Texts(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AnyContract.contract.Call(opts, out, "texts", arg0)
	return *ret0, err
}

// Texts is a free data retrieval call binding the contract method 0x0e0f9cbb.
//
// Solidity: function texts( address) constant returns(string)
func (_AnyContract *AnyContractSession) Texts(arg0 common.Address) (string, error) {
	return _AnyContract.Contract.Texts(&_AnyContract.CallOpts, arg0)
}

// Texts is a free data retrieval call binding the contract method 0x0e0f9cbb.
//
// Solidity: function texts( address) constant returns(string)
func (_AnyContract *AnyContractCallerSession) Texts(arg0 common.Address) (string, error) {
	return _AnyContract.Contract.Texts(&_AnyContract.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(_a uint256, _b uint256) returns()
func (_AnyContract *AnyContractTransactor) Add(opts *bind.TransactOpts, _a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _AnyContract.contract.Transact(opts, "add", _a, _b)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(_a uint256, _b uint256) returns()
func (_AnyContract *AnyContractSession) Add(_a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _AnyContract.Contract.Add(&_AnyContract.TransactOpts, _a, _b)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(_a uint256, _b uint256) returns()
func (_AnyContract *AnyContractTransactorSession) Add(_a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _AnyContract.Contract.Add(&_AnyContract.TransactOpts, _a, _b)
}

// BatchWrite is a paid mutator transaction binding the contract method 0xabc40f41.
//
// Solidity: function batchWrite(_a uint256, _b uint256, _text string) returns()
func (_AnyContract *AnyContractTransactor) BatchWrite(opts *bind.TransactOpts, _a *big.Int, _b *big.Int, _text string) (*types.Transaction, error) {
	return _AnyContract.contract.Transact(opts, "batchWrite", _a, _b, _text)
}

// BatchWrite is a paid mutator transaction binding the contract method 0xabc40f41.
//
// Solidity: function batchWrite(_a uint256, _b uint256, _text string) returns()
func (_AnyContract *AnyContractSession) BatchWrite(_a *big.Int, _b *big.Int, _text string) (*types.Transaction, error) {
	return _AnyContract.Contract.BatchWrite(&_AnyContract.TransactOpts, _a, _b, _text)
}

// BatchWrite is a paid mutator transaction binding the contract method 0xabc40f41.
//
// Solidity: function batchWrite(_a uint256, _b uint256, _text string) returns()
func (_AnyContract *AnyContractTransactorSession) BatchWrite(_a *big.Int, _b *big.Int, _text string) (*types.Transaction, error) {
	return _AnyContract.Contract.BatchWrite(&_AnyContract.TransactOpts, _a, _b, _text)
}

// Write is a paid mutator transaction binding the contract method 0xebaac771.
//
// Solidity: function write(_text string) returns()
func (_AnyContract *AnyContractTransactor) Write(opts *bind.TransactOpts, _text string) (*types.Transaction, error) {
	return _AnyContract.contract.Transact(opts, "write", _text)
}

// Write is a paid mutator transaction binding the contract method 0xebaac771.
//
// Solidity: function write(_text string) returns()
func (_AnyContract *AnyContractSession) Write(_text string) (*types.Transaction, error) {
	return _AnyContract.Contract.Write(&_AnyContract.TransactOpts, _text)
}

// Write is a paid mutator transaction binding the contract method 0xebaac771.
//
// Solidity: function write(_text string) returns()
func (_AnyContract *AnyContractTransactorSession) Write(_text string) (*types.Transaction, error) {
	return _AnyContract.Contract.Write(&_AnyContract.TransactOpts, _text)
}

// DynamicCallerABI is the input ABI used to generate the binding from.
const DynamicCallerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_constract\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"dyn_call\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// DynamicCallerBin is the compiled bytecode used for deploying new contracts.
const DynamicCallerBin = `0x608060405234801561001057600080fd5b5061016e806100206000396000f3006080604052600436106100405763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637b77bbc88114610045575b600080fd5b60408051602060046024803582810135601f81018590048502860185019096528585526100ac95833573ffffffffffffffffffffffffffffffffffffffff169536956044949193909101919081908401838280828437509497506100ae9650505050505050565b005b8173ffffffffffffffffffffffffffffffffffffffff16348260405180828051906020019080838360005b838110156100f15781810151838201526020016100d9565b50505050905090810190601f16801561011e5780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af192505050151561013e57600080fd5b50505600a165627a7a7230582098d6d35d2a3d5dd35eb9c05ae2ce8464743f216a086d78ea00720d39f3de08420029`

// DeployDynamicCaller deploys a new Ethereum contract, binding an instance of DynamicCaller to it.
func DeployDynamicCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DynamicCaller, error) {
	parsed, err := abi.JSON(strings.NewReader(DynamicCallerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DynamicCallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DynamicCaller{DynamicCallerCaller: DynamicCallerCaller{contract: contract}, DynamicCallerTransactor: DynamicCallerTransactor{contract: contract}, DynamicCallerFilterer: DynamicCallerFilterer{contract: contract}}, nil
}

// DynamicCaller is an auto generated Go binding around an Ethereum contract.
type DynamicCaller struct {
	DynamicCallerCaller     // Read-only binding to the contract
	DynamicCallerTransactor // Write-only binding to the contract
	DynamicCallerFilterer   // Log filterer for contract events
}

// DynamicCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DynamicCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DynamicCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DynamicCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DynamicCallerSession struct {
	Contract     *DynamicCaller    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DynamicCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DynamicCallerCallerSession struct {
	Contract *DynamicCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DynamicCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DynamicCallerTransactorSession struct {
	Contract     *DynamicCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DynamicCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DynamicCallerRaw struct {
	Contract *DynamicCaller // Generic contract binding to access the raw methods on
}

// DynamicCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DynamicCallerCallerRaw struct {
	Contract *DynamicCallerCaller // Generic read-only contract binding to access the raw methods on
}

// DynamicCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DynamicCallerTransactorRaw struct {
	Contract *DynamicCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDynamicCaller creates a new instance of DynamicCaller, bound to a specific deployed contract.
func NewDynamicCaller(address common.Address, backend bind.ContractBackend) (*DynamicCaller, error) {
	contract, err := bindDynamicCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DynamicCaller{DynamicCallerCaller: DynamicCallerCaller{contract: contract}, DynamicCallerTransactor: DynamicCallerTransactor{contract: contract}, DynamicCallerFilterer: DynamicCallerFilterer{contract: contract}}, nil
}

// NewDynamicCallerCaller creates a new read-only instance of DynamicCaller, bound to a specific deployed contract.
func NewDynamicCallerCaller(address common.Address, caller bind.ContractCaller) (*DynamicCallerCaller, error) {
	contract, err := bindDynamicCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DynamicCallerCaller{contract: contract}, nil
}

// NewDynamicCallerTransactor creates a new write-only instance of DynamicCaller, bound to a specific deployed contract.
func NewDynamicCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*DynamicCallerTransactor, error) {
	contract, err := bindDynamicCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DynamicCallerTransactor{contract: contract}, nil
}

// NewDynamicCallerFilterer creates a new log filterer instance of DynamicCaller, bound to a specific deployed contract.
func NewDynamicCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*DynamicCallerFilterer, error) {
	contract, err := bindDynamicCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DynamicCallerFilterer{contract: contract}, nil
}

// bindDynamicCaller binds a generic wrapper to an already deployed contract.
func bindDynamicCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DynamicCallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DynamicCaller *DynamicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DynamicCaller.Contract.DynamicCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DynamicCaller *DynamicCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicCaller.Contract.DynamicCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DynamicCaller *DynamicCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DynamicCaller.Contract.DynamicCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DynamicCaller *DynamicCallerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DynamicCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DynamicCaller *DynamicCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DynamicCaller *DynamicCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DynamicCaller.Contract.contract.Transact(opts, method, params...)
}

// DynCall is a paid mutator transaction binding the contract method 0x7b77bbc8.
//
// Solidity: function dyn_call(_constract address, _data bytes) returns()
func (_DynamicCaller *DynamicCallerTransactor) DynCall(opts *bind.TransactOpts, _constract common.Address, _data []byte) (*types.Transaction, error) {
	return _DynamicCaller.contract.Transact(opts, "dyn_call", _constract, _data)
}

// DynCall is a paid mutator transaction binding the contract method 0x7b77bbc8.
//
// Solidity: function dyn_call(_constract address, _data bytes) returns()
func (_DynamicCaller *DynamicCallerSession) DynCall(_constract common.Address, _data []byte) (*types.Transaction, error) {
	return _DynamicCaller.Contract.DynCall(&_DynamicCaller.TransactOpts, _constract, _data)
}

// DynCall is a paid mutator transaction binding the contract method 0x7b77bbc8.
//
// Solidity: function dyn_call(_constract address, _data bytes) returns()
func (_DynamicCaller *DynamicCallerTransactorSession) DynCall(_constract common.Address, _data []byte) (*types.Transaction, error) {
	return _DynamicCaller.Contract.DynCall(&_DynamicCaller.TransactOpts, _constract, _data)
}
