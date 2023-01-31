// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package etherwatt

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

// EtherwattMetaData contains all meta data concerning the Etherwatt contract.
var EtherwattMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060400160405280600981526020017f4574686572576174740000000000000000000000000000000000000000000000815250600090816200004a919062000389565b506040518060400160405280600281526020017f45570000000000000000000000000000000000000000000000000000000000008152506001908162000091919062000389565b506001600260006101000a81548160ff021916908360ff160217905550620f4240600355348015620000c257600080fd5b50600354600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555062000470565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200019157607f821691505b602082108103620001a757620001a662000149565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620002117fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620001d2565b6200021d8683620001d2565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200026a620002646200025e8462000235565b6200023f565b62000235565b9050919050565b6000819050919050565b620002868362000249565b6200029e620002958262000271565b848454620001df565b825550505050565b600090565b620002b5620002a6565b620002c28184846200027b565b505050565b5b81811015620002ea57620002de600082620002ab565b600181019050620002c8565b5050565b601f82111562000339576200030381620001ad565b6200030e84620001c2565b810160208510156200031e578190505b620003366200032d85620001c2565b830182620002c7565b50505b505050565b600082821c905092915050565b60006200035e600019846008026200033e565b1980831691505092915050565b60006200037983836200034b565b9150826002028217905092915050565b62000394826200010f565b67ffffffffffffffff811115620003b057620003af6200011a565b5b620003bc825462000178565b620003c9828285620002ee565b600060209050601f831160018114620004015760008415620003ec578287015190505b620003f885826200036b565b86555062000468565b601f1984166200041186620001ad565b60005b828110156200043b5784890151825560018201915060208501945060208101905062000414565b868310156200045b578489015162000457601f8916826200034b565b8355505b6001600288020188555050505b505050505050565b61092380620004806000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063313ce5671161005b578063313ce567146100da57806370a08231146100f857806395d89b4114610128578063beabacc8146101465761007d565b806306fdde031461008257806318160ddd146100a057806321670f22146100be575b600080fd5b61008a610162565b604051610097919061061a565b60405180910390f35b6100a86101f0565b6040516100b59190610655565b60405180910390f35b6100d860048036038101906100d391906106ff565b6101f6565b005b6100e2610363565b6040516100ef919061075b565b60405180910390f35b610112600480360381019061010d9190610776565b610376565b60405161011f9190610655565b60405180910390f35b61013061038e565b60405161013d919061061a565b60405180910390f35b610160600480360381019061015b91906107a3565b61041c565b005b6000805461016f90610825565b80601f016020809104026020016040519081016040528092919081815260200182805461019b90610825565b80156101e85780601f106101bd576101008083540402835291602001916101e8565b820191906000526020600020905b8154815290600101906020018083116101cb57829003601f168201915b505050505081565b60035481565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101580156102455750600081115b61024e57600080fd5b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461029d9190610885565b9250508190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102f391906108b9565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516103579190610655565b60405180910390a35050565b600260009054906101000a900460ff1681565b60046020528060005260406000206000915090505481565b6001805461039b90610825565b80601f01602080910402602001604051908101604052809291908181526020018280546103c790610825565b80156104145780601f106103e957610100808354040283529160200191610414565b820191906000526020600020905b8154815290600101906020018083116103f757829003601f168201915b505050505081565b80600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015801561046b5750600081115b61047457600080fd5b80600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104c39190610885565b9250508190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461051991906108b9565b925050819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161057d9190610655565b60405180910390a3505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156105c45780820151818401526020810190506105a9565b60008484015250505050565b6000601f19601f8301169050919050565b60006105ec8261058a565b6105f68185610595565b93506106068185602086016105a6565b61060f816105d0565b840191505092915050565b6000602082019050818103600083015261063481846105e1565b905092915050565b6000819050919050565b61064f8161063c565b82525050565b600060208201905061066a6000830184610646565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106a082610675565b9050919050565b6106b081610695565b81146106bb57600080fd5b50565b6000813590506106cd816106a7565b92915050565b6106dc8161063c565b81146106e757600080fd5b50565b6000813590506106f9816106d3565b92915050565b6000806040838503121561071657610715610670565b5b6000610724858286016106be565b9250506020610735858286016106ea565b9150509250929050565b600060ff82169050919050565b6107558161073f565b82525050565b6000602082019050610770600083018461074c565b92915050565b60006020828403121561078c5761078b610670565b5b600061079a848285016106be565b91505092915050565b6000806000606084860312156107bc576107bb610670565b5b60006107ca868287016106be565b93505060206107db868287016106be565b92505060406107ec868287016106ea565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061083d57607f821691505b6020821081036108505761084f6107f6565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006108908261063c565b915061089b8361063c565b92508282039050818111156108b3576108b2610856565b5b92915050565b60006108c48261063c565b91506108cf8361063c565b92508282019050808211156108e7576108e6610856565b5b9291505056fea26469706673582212203dd747bdfdee990551bd920fa82c6822d74ee71386a5fac3238e6434d3bb7d6264736f6c63430008110033",
}

// EtherwattABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherwattMetaData.ABI instead.
var EtherwattABI = EtherwattMetaData.ABI

// EtherwattBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EtherwattMetaData.Bin instead.
var EtherwattBin = EtherwattMetaData.Bin

// DeployEtherwatt deploys a new Ethereum contract, binding an instance of Etherwatt to it.
func DeployEtherwatt(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Etherwatt, error) {
	parsed, err := EtherwattMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtherwattBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Etherwatt{EtherwattCaller: EtherwattCaller{contract: contract}, EtherwattTransactor: EtherwattTransactor{contract: contract}, EtherwattFilterer: EtherwattFilterer{contract: contract}}, nil
}

// Etherwatt is an auto generated Go binding around an Ethereum contract.
type Etherwatt struct {
	EtherwattCaller     // Read-only binding to the contract
	EtherwattTransactor // Write-only binding to the contract
	EtherwattFilterer   // Log filterer for contract events
}

// EtherwattCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherwattCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherwattTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherwattTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherwattFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherwattFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherwattSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherwattSession struct {
	Contract     *Etherwatt        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherwattCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherwattCallerSession struct {
	Contract *EtherwattCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EtherwattTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherwattTransactorSession struct {
	Contract     *EtherwattTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EtherwattRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherwattRaw struct {
	Contract *Etherwatt // Generic contract binding to access the raw methods on
}

// EtherwattCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherwattCallerRaw struct {
	Contract *EtherwattCaller // Generic read-only contract binding to access the raw methods on
}

// EtherwattTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherwattTransactorRaw struct {
	Contract *EtherwattTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherwatt creates a new instance of Etherwatt, bound to a specific deployed contract.
func NewEtherwatt(address common.Address, backend bind.ContractBackend) (*Etherwatt, error) {
	contract, err := bindEtherwatt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Etherwatt{EtherwattCaller: EtherwattCaller{contract: contract}, EtherwattTransactor: EtherwattTransactor{contract: contract}, EtherwattFilterer: EtherwattFilterer{contract: contract}}, nil
}

// NewEtherwattCaller creates a new read-only instance of Etherwatt, bound to a specific deployed contract.
func NewEtherwattCaller(address common.Address, caller bind.ContractCaller) (*EtherwattCaller, error) {
	contract, err := bindEtherwatt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherwattCaller{contract: contract}, nil
}

// NewEtherwattTransactor creates a new write-only instance of Etherwatt, bound to a specific deployed contract.
func NewEtherwattTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherwattTransactor, error) {
	contract, err := bindEtherwatt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherwattTransactor{contract: contract}, nil
}

// NewEtherwattFilterer creates a new log filterer instance of Etherwatt, bound to a specific deployed contract.
func NewEtherwattFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherwattFilterer, error) {
	contract, err := bindEtherwatt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherwattFilterer{contract: contract}, nil
}

// bindEtherwatt binds a generic wrapper to an already deployed contract.
func bindEtherwatt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherwattMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etherwatt *EtherwattRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Etherwatt.Contract.EtherwattCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etherwatt *EtherwattRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etherwatt.Contract.EtherwattTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etherwatt *EtherwattRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etherwatt.Contract.EtherwattTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Etherwatt *EtherwattCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Etherwatt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Etherwatt *EtherwattTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Etherwatt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Etherwatt *EtherwattTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Etherwatt.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Etherwatt *EtherwattCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Etherwatt.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Etherwatt *EtherwattSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Etherwatt.Contract.BalanceOf(&_Etherwatt.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Etherwatt *EtherwattCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Etherwatt.Contract.BalanceOf(&_Etherwatt.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Etherwatt *EtherwattCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Etherwatt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Etherwatt *EtherwattSession) Decimals() (uint8, error) {
	return _Etherwatt.Contract.Decimals(&_Etherwatt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Etherwatt *EtherwattCallerSession) Decimals() (uint8, error) {
	return _Etherwatt.Contract.Decimals(&_Etherwatt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Etherwatt *EtherwattCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Etherwatt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Etherwatt *EtherwattSession) Name() (string, error) {
	return _Etherwatt.Contract.Name(&_Etherwatt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Etherwatt *EtherwattCallerSession) Name() (string, error) {
	return _Etherwatt.Contract.Name(&_Etherwatt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Etherwatt *EtherwattCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Etherwatt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Etherwatt *EtherwattSession) Symbol() (string, error) {
	return _Etherwatt.Contract.Symbol(&_Etherwatt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Etherwatt *EtherwattCallerSession) Symbol() (string, error) {
	return _Etherwatt.Contract.Symbol(&_Etherwatt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Etherwatt *EtherwattCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Etherwatt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Etherwatt *EtherwattSession) TotalSupply() (*big.Int, error) {
	return _Etherwatt.Contract.TotalSupply(&_Etherwatt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Etherwatt *EtherwattCallerSession) TotalSupply() (*big.Int, error) {
	return _Etherwatt.Contract.TotalSupply(&_Etherwatt.CallOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address _to, uint256 _value) returns()
func (_Etherwatt *EtherwattTransactor) Reward(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Etherwatt.contract.Transact(opts, "reward", _to, _value)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address _to, uint256 _value) returns()
func (_Etherwatt *EtherwattSession) Reward(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Etherwatt.Contract.Reward(&_Etherwatt.TransactOpts, _to, _value)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address _to, uint256 _value) returns()
func (_Etherwatt *EtherwattTransactorSession) Reward(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Etherwatt.Contract.Reward(&_Etherwatt.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _from, address _to, uint256 _value) returns()
func (_Etherwatt *EtherwattTransactor) Transfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Etherwatt.contract.Transact(opts, "transfer", _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _from, address _to, uint256 _value) returns()
func (_Etherwatt *EtherwattSession) Transfer(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Etherwatt.Contract.Transfer(&_Etherwatt.TransactOpts, _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _from, address _to, uint256 _value) returns()
func (_Etherwatt *EtherwattTransactorSession) Transfer(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Etherwatt.Contract.Transfer(&_Etherwatt.TransactOpts, _from, _to, _value)
}

// EtherwattTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Etherwatt contract.
type EtherwattTransferIterator struct {
	Event *EtherwattTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EtherwattTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherwattTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EtherwattTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EtherwattTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherwattTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherwattTransfer represents a Transfer event raised by the Etherwatt contract.
type EtherwattTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Etherwatt *EtherwattFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EtherwattTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Etherwatt.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EtherwattTransferIterator{contract: _Etherwatt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Etherwatt *EtherwattFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EtherwattTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Etherwatt.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherwattTransfer)
				if err := _Etherwatt.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Etherwatt *EtherwattFilterer) ParseTransfer(log types.Log) (*EtherwattTransfer, error) {
	event := new(EtherwattTransfer)
	if err := _Etherwatt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
