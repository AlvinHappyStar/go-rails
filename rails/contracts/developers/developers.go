// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package developers

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
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"AdminTransferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminTransferCommenced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminTransferCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminTransferCancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"adminTransferCommence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminTransferConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"prohibit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611287806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a87d942c11610066578063a87d942c1461010a578063c4d66de814610128578063c4e1ccf214610144578063daea85c514610160578063f767742a1461017c57610093565b80630141bc1414610098578063328d8f72146100a2578063477e1694146100be578063673448dd146100da575b600080fd5b6100a0610186565b005b6100bc60048036038101906100b79190610cc3565b610316565b005b6100d860048036038101906100d39190610d4e565b610420565b005b6100f460048036038101906100ef9190610d4e565b610588565b6040516101019190610d8a565b60405180910390f35b6101126105de565b60405161011f9190610dbe565b60405180910390f35b610142600480360381019061013d9190610d4e565b6105fd565b005b61015e60048036038101906101599190610d4e565b6106a9565b005b61017a60048036038101906101759190610d4e565b6108be565b005b610184610ad4565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610216576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020d90610e5c565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe6aa55180442bd9b002c0b9ddf8ebf164046b2d06d4af4cfddd92b4ba770469c600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161030c9190610e8b565b60405180910390a1565b60008054906101000a900460ff16610363576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035a90610f18565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614806103c457506103c333610c31565b5b610403576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fa90610f84565b60405180910390fd5b80600060016101000a81548160ff02191690831515021790555050565b60008054906101000a900460ff1661046d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046490610f18565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614806104ce57506104cd33610c31565b5b61050d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050490610f84565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f547e8945dcdad71973153006f3a59166737ec7756716d5974fbe9ff4077131e88160405161057d9190610e8b565b60405180910390a150565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60008060169054906101000a900463ffffffff1663ffffffff16905090565b60008054906101000a900460ff161561064b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064290611016565b60405180910390fd5b80600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016000806101000a81548160ff02191690831515021790555050565b60008054906101000a900460ff166106f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ed90610f18565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480610757575061075633610c31565b5b610796576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078d90610f84565b60405180910390fd5b600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610822576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610819906110a8565b60405180910390fd5b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000601681819054906101000a900463ffffffff168092919061089c90611107565b91906101000a81548163ffffffff021916908363ffffffff1602179055505050565b60008054906101000a900460ff1661090b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090290610f18565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061096c575061096b33610c31565b5b6109ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a290610f84565b60405180910390fd5b600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610a38576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2f906111a2565b60405180910390fd5b6001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000601681819054906101000a900463ffffffff1680929190610ab2906111c2565b91906101000a81548163ffffffff021916908363ffffffff1602179055505050565b60008054906101000a900460ff16610b21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1890610f18565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480610b825750610b8133610c31565b5b610bc1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb890610f84565b60405180910390fd5b6000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2dac8fcf9c63063c63818ae1028d7f8bd44d417e65c849064a1d904748eae39b60405160405180910390a1565b60007f410ba2fdd1f5a02a6adaab564a7889fd9586c1d301911dfaebefee65f38c49e660001b82604051602001610c689190611236565b60405160208183030381529060405280519060200120149050919050565b600080fd5b60008115159050919050565b610ca081610c8b565b8114610cab57600080fd5b50565b600081359050610cbd81610c97565b92915050565b600060208284031215610cd957610cd8610c86565b5b6000610ce784828501610cae565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d1b82610cf0565b9050919050565b610d2b81610d10565b8114610d3657600080fd5b50565b600081359050610d4881610d22565b92915050565b600060208284031215610d6457610d63610c86565b5b6000610d7284828501610d39565b91505092915050565b610d8481610c8b565b82525050565b6000602082019050610d9f6000830184610d7b565b92915050565b6000819050919050565b610db881610da5565b82525050565b6000602082019050610dd36000830184610daf565b92915050565b600082825260208201905092915050565b7f446576656c6f706572732e61646d696e5472616e73666572436f6e6669726d3a60008201527f2050656e64696e672061646d696e206f6e6c7921000000000000000000000000602082015250565b6000610e46603483610dd9565b9150610e5182610dea565b604082019050919050565b60006020820190508181036000830152610e7581610e39565b9050919050565b610e8581610d10565b82525050565b6000602082019050610ea06000830184610e7c565b92915050565b7f446576656c6f706572732e6f6e6c7941646d696e3a204e6f7420696e6974696160008201527f6c697a6564210000000000000000000000000000000000000000000000000000602082015250565b6000610f02602683610dd9565b9150610f0d82610ea6565b604082019050919050565b60006020820190508181036000830152610f3181610ef5565b9050919050565b7f556e617574686f72697a65642100000000000000000000000000000000000000600082015250565b6000610f6e600d83610dd9565b9150610f7982610f38565b602082019050919050565b60006020820190508181036000830152610f9d81610f61565b9050919050565b7f446576656c6f706572732e696e697469616c697a653a20416c7265616479206960008201527f6e697469616c697a656421000000000000000000000000000000000000000000602082015250565b6000611000602b83610dd9565b915061100b82610fa4565b604082019050919050565b6000602082019050818103600083015261102f81610ff3565b9050919050565b7f446576656c6f706572732e70726f68696269743a20416c72656164792070726f60008201527f6869626974656421000000000000000000000000000000000000000000000000602082015250565b6000611092602883610dd9565b915061109d82611036565b604082019050919050565b600060208201905081810360008301526110c181611085565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600063ffffffff82169050919050565b6000611112826110f7565b915060008203611125576111246110c8565b5b600182039050919050565b7f446576656c6f706572732e617070726f76653a20416c7265616479206170707260008201527f6f76656421000000000000000000000000000000000000000000000000000000602082015250565b600061118c602583610dd9565b915061119782611130565b604082019050919050565b600060208201905081810360008301526111bb8161117f565b9050919050565b60006111cd826110f7565b915063ffffffff82036111e3576111e26110c8565b5b600182019050919050565b60008160601b9050919050565b6000611206826111ee565b9050919050565b6000611218826111fb565b9050919050565b61123061122b82610d10565b61120d565b82525050565b6000611242828461121f565b6014820191508190509291505056fea264697066735822122025fe75d812d92fffdf0c4d16d83c446787e2c7adb4b6f46a7a9ea58a1a7b018464736f6c63430008100033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Contract *ContractCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Contract *ContractSession) GetCount() (*big.Int, error) {
	return _Contract.Contract.GetCount(&_Contract.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Contract *ContractCallerSession) GetCount() (*big.Int, error) {
	return _Contract.Contract.GetCount(&_Contract.CallOpts)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address _address) view returns(bool)
func (_Contract *ContractCaller) IsApproved(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApproved", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address _address) view returns(bool)
func (_Contract *ContractSession) IsApproved(_address common.Address) (bool, error) {
	return _Contract.Contract.IsApproved(&_Contract.CallOpts, _address)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address _address) view returns(bool)
func (_Contract *ContractCallerSession) IsApproved(_address common.Address) (bool, error) {
	return _Contract.Contract.IsApproved(&_Contract.CallOpts, _address)
}

// AdminTransferCancel is a paid mutator transaction binding the contract method 0xf767742a.
//
// Solidity: function adminTransferCancel() returns()
func (_Contract *ContractTransactor) AdminTransferCancel(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "adminTransferCancel")
}

// AdminTransferCancel is a paid mutator transaction binding the contract method 0xf767742a.
//
// Solidity: function adminTransferCancel() returns()
func (_Contract *ContractSession) AdminTransferCancel() (*types.Transaction, error) {
	return _Contract.Contract.AdminTransferCancel(&_Contract.TransactOpts)
}

// AdminTransferCancel is a paid mutator transaction binding the contract method 0xf767742a.
//
// Solidity: function adminTransferCancel() returns()
func (_Contract *ContractTransactorSession) AdminTransferCancel() (*types.Transaction, error) {
	return _Contract.Contract.AdminTransferCancel(&_Contract.TransactOpts)
}

// AdminTransferCommence is a paid mutator transaction binding the contract method 0x477e1694.
//
// Solidity: function adminTransferCommence(address newAdmin) returns()
func (_Contract *ContractTransactor) AdminTransferCommence(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "adminTransferCommence", newAdmin)
}

// AdminTransferCommence is a paid mutator transaction binding the contract method 0x477e1694.
//
// Solidity: function adminTransferCommence(address newAdmin) returns()
func (_Contract *ContractSession) AdminTransferCommence(newAdmin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AdminTransferCommence(&_Contract.TransactOpts, newAdmin)
}

// AdminTransferCommence is a paid mutator transaction binding the contract method 0x477e1694.
//
// Solidity: function adminTransferCommence(address newAdmin) returns()
func (_Contract *ContractTransactorSession) AdminTransferCommence(newAdmin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AdminTransferCommence(&_Contract.TransactOpts, newAdmin)
}

// AdminTransferConfirm is a paid mutator transaction binding the contract method 0x0141bc14.
//
// Solidity: function adminTransferConfirm() returns()
func (_Contract *ContractTransactor) AdminTransferConfirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "adminTransferConfirm")
}

// AdminTransferConfirm is a paid mutator transaction binding the contract method 0x0141bc14.
//
// Solidity: function adminTransferConfirm() returns()
func (_Contract *ContractSession) AdminTransferConfirm() (*types.Transaction, error) {
	return _Contract.Contract.AdminTransferConfirm(&_Contract.TransactOpts)
}

// AdminTransferConfirm is a paid mutator transaction binding the contract method 0x0141bc14.
//
// Solidity: function adminTransferConfirm() returns()
func (_Contract *ContractTransactorSession) AdminTransferConfirm() (*types.Transaction, error) {
	return _Contract.Contract.AdminTransferConfirm(&_Contract.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _address) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", _address)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _address) returns()
func (_Contract *ContractSession) Approve(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _address)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _address) returns()
func (_Contract *ContractTransactorSession) Approve(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _address)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_Contract *ContractSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_Contract *ContractTransactorSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _admin)
}

// Prohibit is a paid mutator transaction binding the contract method 0xc4e1ccf2.
//
// Solidity: function prohibit(address _address) returns()
func (_Contract *ContractTransactor) Prohibit(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prohibit", _address)
}

// Prohibit is a paid mutator transaction binding the contract method 0xc4e1ccf2.
//
// Solidity: function prohibit(address _address) returns()
func (_Contract *ContractSession) Prohibit(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Prohibit(&_Contract.TransactOpts, _address)
}

// Prohibit is a paid mutator transaction binding the contract method 0xc4e1ccf2.
//
// Solidity: function prohibit(address _address) returns()
func (_Contract *ContractTransactorSession) Prohibit(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Prohibit(&_Contract.TransactOpts, _address)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool value) returns()
func (_Contract *ContractTransactor) SetEnabled(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setEnabled", value)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool value) returns()
func (_Contract *ContractSession) SetEnabled(value bool) (*types.Transaction, error) {
	return _Contract.Contract.SetEnabled(&_Contract.TransactOpts, value)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool value) returns()
func (_Contract *ContractTransactorSession) SetEnabled(value bool) (*types.Transaction, error) {
	return _Contract.Contract.SetEnabled(&_Contract.TransactOpts, value)
}

// ContractAdminTransferCancelledIterator is returned from FilterAdminTransferCancelled and is used to iterate over the raw logs and unpacked data for AdminTransferCancelled events raised by the Contract contract.
type ContractAdminTransferCancelledIterator struct {
	Event *ContractAdminTransferCancelled // Event containing the contract specifics and raw log

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
func (it *ContractAdminTransferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAdminTransferCancelled)
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
		it.Event = new(ContractAdminTransferCancelled)
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
func (it *ContractAdminTransferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAdminTransferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAdminTransferCancelled represents a AdminTransferCancelled event raised by the Contract contract.
type ContractAdminTransferCancelled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAdminTransferCancelled is a free log retrieval operation binding the contract event 0x2dac8fcf9c63063c63818ae1028d7f8bd44d417e65c849064a1d904748eae39b.
//
// Solidity: event AdminTransferCancelled()
func (_Contract *ContractFilterer) FilterAdminTransferCancelled(opts *bind.FilterOpts) (*ContractAdminTransferCancelledIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AdminTransferCancelled")
	if err != nil {
		return nil, err
	}
	return &ContractAdminTransferCancelledIterator{contract: _Contract.contract, event: "AdminTransferCancelled", logs: logs, sub: sub}, nil
}

// WatchAdminTransferCancelled is a free log subscription operation binding the contract event 0x2dac8fcf9c63063c63818ae1028d7f8bd44d417e65c849064a1d904748eae39b.
//
// Solidity: event AdminTransferCancelled()
func (_Contract *ContractFilterer) WatchAdminTransferCancelled(opts *bind.WatchOpts, sink chan<- *ContractAdminTransferCancelled) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AdminTransferCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAdminTransferCancelled)
				if err := _Contract.contract.UnpackLog(event, "AdminTransferCancelled", log); err != nil {
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

// ParseAdminTransferCancelled is a log parse operation binding the contract event 0x2dac8fcf9c63063c63818ae1028d7f8bd44d417e65c849064a1d904748eae39b.
//
// Solidity: event AdminTransferCancelled()
func (_Contract *ContractFilterer) ParseAdminTransferCancelled(log types.Log) (*ContractAdminTransferCancelled, error) {
	event := new(ContractAdminTransferCancelled)
	if err := _Contract.contract.UnpackLog(event, "AdminTransferCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAdminTransferCommencedIterator is returned from FilterAdminTransferCommenced and is used to iterate over the raw logs and unpacked data for AdminTransferCommenced events raised by the Contract contract.
type ContractAdminTransferCommencedIterator struct {
	Event *ContractAdminTransferCommenced // Event containing the contract specifics and raw log

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
func (it *ContractAdminTransferCommencedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAdminTransferCommenced)
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
		it.Event = new(ContractAdminTransferCommenced)
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
func (it *ContractAdminTransferCommencedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAdminTransferCommencedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAdminTransferCommenced represents a AdminTransferCommenced event raised by the Contract contract.
type ContractAdminTransferCommenced struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminTransferCommenced is a free log retrieval operation binding the contract event 0x547e8945dcdad71973153006f3a59166737ec7756716d5974fbe9ff4077131e8.
//
// Solidity: event AdminTransferCommenced(address newAdmin)
func (_Contract *ContractFilterer) FilterAdminTransferCommenced(opts *bind.FilterOpts) (*ContractAdminTransferCommencedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AdminTransferCommenced")
	if err != nil {
		return nil, err
	}
	return &ContractAdminTransferCommencedIterator{contract: _Contract.contract, event: "AdminTransferCommenced", logs: logs, sub: sub}, nil
}

// WatchAdminTransferCommenced is a free log subscription operation binding the contract event 0x547e8945dcdad71973153006f3a59166737ec7756716d5974fbe9ff4077131e8.
//
// Solidity: event AdminTransferCommenced(address newAdmin)
func (_Contract *ContractFilterer) WatchAdminTransferCommenced(opts *bind.WatchOpts, sink chan<- *ContractAdminTransferCommenced) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AdminTransferCommenced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAdminTransferCommenced)
				if err := _Contract.contract.UnpackLog(event, "AdminTransferCommenced", log); err != nil {
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

// ParseAdminTransferCommenced is a log parse operation binding the contract event 0x547e8945dcdad71973153006f3a59166737ec7756716d5974fbe9ff4077131e8.
//
// Solidity: event AdminTransferCommenced(address newAdmin)
func (_Contract *ContractFilterer) ParseAdminTransferCommenced(log types.Log) (*ContractAdminTransferCommenced, error) {
	event := new(ContractAdminTransferCommenced)
	if err := _Contract.contract.UnpackLog(event, "AdminTransferCommenced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAdminTransferCompletedIterator is returned from FilterAdminTransferCompleted and is used to iterate over the raw logs and unpacked data for AdminTransferCompleted events raised by the Contract contract.
type ContractAdminTransferCompletedIterator struct {
	Event *ContractAdminTransferCompleted // Event containing the contract specifics and raw log

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
func (it *ContractAdminTransferCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAdminTransferCompleted)
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
		it.Event = new(ContractAdminTransferCompleted)
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
func (it *ContractAdminTransferCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAdminTransferCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAdminTransferCompleted represents a AdminTransferCompleted event raised by the Contract contract.
type ContractAdminTransferCompleted struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminTransferCompleted is a free log retrieval operation binding the contract event 0xe6aa55180442bd9b002c0b9ddf8ebf164046b2d06d4af4cfddd92b4ba770469c.
//
// Solidity: event AdminTransferCompleted(address newAdmin)
func (_Contract *ContractFilterer) FilterAdminTransferCompleted(opts *bind.FilterOpts) (*ContractAdminTransferCompletedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AdminTransferCompleted")
	if err != nil {
		return nil, err
	}
	return &ContractAdminTransferCompletedIterator{contract: _Contract.contract, event: "AdminTransferCompleted", logs: logs, sub: sub}, nil
}

// WatchAdminTransferCompleted is a free log subscription operation binding the contract event 0xe6aa55180442bd9b002c0b9ddf8ebf164046b2d06d4af4cfddd92b4ba770469c.
//
// Solidity: event AdminTransferCompleted(address newAdmin)
func (_Contract *ContractFilterer) WatchAdminTransferCompleted(opts *bind.WatchOpts, sink chan<- *ContractAdminTransferCompleted) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AdminTransferCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAdminTransferCompleted)
				if err := _Contract.contract.UnpackLog(event, "AdminTransferCompleted", log); err != nil {
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

// ParseAdminTransferCompleted is a log parse operation binding the contract event 0xe6aa55180442bd9b002c0b9ddf8ebf164046b2d06d4af4cfddd92b4ba770469c.
//
// Solidity: event AdminTransferCompleted(address newAdmin)
func (_Contract *ContractFilterer) ParseAdminTransferCompleted(log types.Log) (*ContractAdminTransferCompleted, error) {
	event := new(ContractAdminTransferCompleted)
	if err := _Contract.contract.UnpackLog(event, "AdminTransferCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
