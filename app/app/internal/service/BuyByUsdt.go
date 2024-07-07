// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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

// BuyByUsdtMetaData contains all meta data concerning the BuyByUsdt contract.
var BuyByUsdtMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountType\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buyCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUserAmountByIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"user\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BuyByUsdtABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyByUsdtMetaData.ABI instead.
var BuyByUsdtABI = BuyByUsdtMetaData.ABI

// BuyByUsdt is an auto generated Go binding around an Ethereum contract.
type BuyByUsdt struct {
	BuyByUsdtCaller     // Read-only binding to the contract
	BuyByUsdtTransactor // Write-only binding to the contract
	BuyByUsdtFilterer   // Log filterer for contract events
}

// BuyByUsdtCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyByUsdtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyByUsdtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyByUsdtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyByUsdtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyByUsdtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyByUsdtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuyByUsdtSession struct {
	Contract     *BuyByUsdt        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyByUsdtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyByUsdtCallerSession struct {
	Contract *BuyByUsdtCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BuyByUsdtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyByUsdtTransactorSession struct {
	Contract     *BuyByUsdtTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BuyByUsdtRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyByUsdtRaw struct {
	Contract *BuyByUsdt // Generic contract binding to access the raw methods on
}

// BuyByUsdtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyByUsdtCallerRaw struct {
	Contract *BuyByUsdtCaller // Generic read-only contract binding to access the raw methods on
}

// BuyByUsdtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyByUsdtTransactorRaw struct {
	Contract *BuyByUsdtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuyByUsdt creates a new instance of BuyByUsdt, bound to a specific deployed contract.
func NewBuyByUsdt(address common.Address, backend bind.ContractBackend) (*BuyByUsdt, error) {
	contract, err := bindBuyByUsdt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BuyByUsdt{BuyByUsdtCaller: BuyByUsdtCaller{contract: contract}, BuyByUsdtTransactor: BuyByUsdtTransactor{contract: contract}, BuyByUsdtFilterer: BuyByUsdtFilterer{contract: contract}}, nil
}

// NewBuyByUsdtCaller creates a new read-only instance of BuyByUsdt, bound to a specific deployed contract.
func NewBuyByUsdtCaller(address common.Address, caller bind.ContractCaller) (*BuyByUsdtCaller, error) {
	contract, err := bindBuyByUsdt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyByUsdtCaller{contract: contract}, nil
}

// NewBuyByUsdtTransactor creates a new write-only instance of BuyByUsdt, bound to a specific deployed contract.
func NewBuyByUsdtTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyByUsdtTransactor, error) {
	contract, err := bindBuyByUsdt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyByUsdtTransactor{contract: contract}, nil
}

// NewBuyByUsdtFilterer creates a new log filterer instance of BuyByUsdt, bound to a specific deployed contract.
func NewBuyByUsdtFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyByUsdtFilterer, error) {
	contract, err := bindBuyByUsdt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyByUsdtFilterer{contract: contract}, nil
}

// bindBuyByUsdt binds a generic wrapper to an already deployed contract.
func bindBuyByUsdt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuyByUsdtABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyByUsdt *BuyByUsdtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyByUsdt.Contract.BuyByUsdtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyByUsdt *BuyByUsdtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyByUsdt.Contract.BuyByUsdtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyByUsdt *BuyByUsdtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyByUsdt.Contract.BuyByUsdtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyByUsdt *BuyByUsdtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuyByUsdt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyByUsdt *BuyByUsdtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyByUsdt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyByUsdt *BuyByUsdtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyByUsdt.Contract.contract.Transact(opts, method, params...)
}

// GetUserAmountByIndex is a free data retrieval call binding the contract method 0xa01ece07.
//
// Solidity: function getUserAmountByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_BuyByUsdt *BuyByUsdtCaller) GetUserAmountByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BuyByUsdt.contract.Call(opts, &out, "getUserAmountByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserAmountByIndex is a free data retrieval call binding the contract method 0xa01ece07.
//
// Solidity: function getUserAmountByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_BuyByUsdt *BuyByUsdtSession) GetUserAmountByIndex(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _BuyByUsdt.Contract.GetUserAmountByIndex(&_BuyByUsdt.CallOpts, startIndex, endIndex)
}

// GetUserAmountByIndex is a free data retrieval call binding the contract method 0xa01ece07.
//
// Solidity: function getUserAmountByIndex(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_BuyByUsdt *BuyByUsdtCallerSession) GetUserAmountByIndex(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _BuyByUsdt.Contract.GetUserAmountByIndex(&_BuyByUsdt.CallOpts, startIndex, endIndex)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_BuyByUsdt *BuyByUsdtCaller) GetUserLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BuyByUsdt.contract.Call(opts, &out, "getUserLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_BuyByUsdt *BuyByUsdtSession) GetUserLength() (*big.Int, error) {
	return _BuyByUsdt.Contract.GetUserLength(&_BuyByUsdt.CallOpts)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_BuyByUsdt *BuyByUsdtCallerSession) GetUserLength() (*big.Int, error) {
	return _BuyByUsdt.Contract.GetUserLength(&_BuyByUsdt.CallOpts)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_BuyByUsdt *BuyByUsdtCaller) GetUsersByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _BuyByUsdt.contract.Call(opts, &out, "getUsersByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_BuyByUsdt *BuyByUsdtSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _BuyByUsdt.Contract.GetUsersByIndex(&_BuyByUsdt.CallOpts, startIndex, endIndex)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_BuyByUsdt *BuyByUsdtCallerSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _BuyByUsdt.Contract.GetUsersByIndex(&_BuyByUsdt.CallOpts, startIndex, endIndex)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BuyByUsdt *BuyByUsdtCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyByUsdt.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BuyByUsdt *BuyByUsdtSession) Owner() (common.Address, error) {
	return _BuyByUsdt.Contract.Owner(&_BuyByUsdt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BuyByUsdt *BuyByUsdtCallerSession) Owner() (common.Address, error) {
	return _BuyByUsdt.Contract.Owner(&_BuyByUsdt.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_BuyByUsdt *BuyByUsdtCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyByUsdt.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_BuyByUsdt *BuyByUsdtSession) Usdt() (common.Address, error) {
	return _BuyByUsdt.Contract.Usdt(&_BuyByUsdt.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_BuyByUsdt *BuyByUsdtCallerSession) Usdt() (common.Address, error) {
	return _BuyByUsdt.Contract.Usdt(&_BuyByUsdt.CallOpts)
}

// User is a free data retrieval call binding the contract method 0x4f8632ba.
//
// Solidity: function user() view returns(address)
func (_BuyByUsdt *BuyByUsdtCaller) User(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuyByUsdt.contract.Call(opts, &out, "user")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// User is a free data retrieval call binding the contract method 0x4f8632ba.
//
// Solidity: function user() view returns(address)
func (_BuyByUsdt *BuyByUsdtSession) User() (common.Address, error) {
	return _BuyByUsdt.Contract.User(&_BuyByUsdt.CallOpts)
}

// User is a free data retrieval call binding the contract method 0x4f8632ba.
//
// Solidity: function user() view returns(address)
func (_BuyByUsdt *BuyByUsdtCallerSession) User() (common.Address, error) {
	return _BuyByUsdt.Contract.User(&_BuyByUsdt.CallOpts)
}

// UserAmount is a free data retrieval call binding the contract method 0xd2b23b72.
//
// Solidity: function userAmount(uint256 ) view returns(uint256)
func (_BuyByUsdt *BuyByUsdtCaller) UserAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BuyByUsdt.contract.Call(opts, &out, "userAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserAmount is a free data retrieval call binding the contract method 0xd2b23b72.
//
// Solidity: function userAmount(uint256 ) view returns(uint256)
func (_BuyByUsdt *BuyByUsdtSession) UserAmount(arg0 *big.Int) (*big.Int, error) {
	return _BuyByUsdt.Contract.UserAmount(&_BuyByUsdt.CallOpts, arg0)
}

// UserAmount is a free data retrieval call binding the contract method 0xd2b23b72.
//
// Solidity: function userAmount(uint256 ) view returns(uint256)
func (_BuyByUsdt *BuyByUsdtCallerSession) UserAmount(arg0 *big.Int) (*big.Int, error) {
	return _BuyByUsdt.Contract.UserAmount(&_BuyByUsdt.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_BuyByUsdt *BuyByUsdtCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BuyByUsdt.contract.Call(opts, &out, "users", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_BuyByUsdt *BuyByUsdtSession) Users(arg0 *big.Int) (common.Address, error) {
	return _BuyByUsdt.Contract.Users(&_BuyByUsdt.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_BuyByUsdt *BuyByUsdtCallerSession) Users(arg0 *big.Int) (common.Address, error) {
	return _BuyByUsdt.Contract.Users(&_BuyByUsdt.CallOpts, arg0)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 amountType) returns()
func (_BuyByUsdt *BuyByUsdtTransactor) Buy(opts *bind.TransactOpts, amountType *big.Int) (*types.Transaction, error) {
	return _BuyByUsdt.contract.Transact(opts, "buy", amountType)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 amountType) returns()
func (_BuyByUsdt *BuyByUsdtSession) Buy(amountType *big.Int) (*types.Transaction, error) {
	return _BuyByUsdt.Contract.Buy(&_BuyByUsdt.TransactOpts, amountType)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 amountType) returns()
func (_BuyByUsdt *BuyByUsdtTransactorSession) Buy(amountType *big.Int) (*types.Transaction, error) {
	return _BuyByUsdt.Contract.Buy(&_BuyByUsdt.TransactOpts, amountType)
}

// SetUser is a paid mutator transaction binding the contract method 0x858ced35.
//
// Solidity: function setUser(address _user) returns()
func (_BuyByUsdt *BuyByUsdtTransactor) SetUser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _BuyByUsdt.contract.Transact(opts, "setUser", _user)
}

// SetUser is a paid mutator transaction binding the contract method 0x858ced35.
//
// Solidity: function setUser(address _user) returns()
func (_BuyByUsdt *BuyByUsdtSession) SetUser(_user common.Address) (*types.Transaction, error) {
	return _BuyByUsdt.Contract.SetUser(&_BuyByUsdt.TransactOpts, _user)
}

// SetUser is a paid mutator transaction binding the contract method 0x858ced35.
//
// Solidity: function setUser(address _user) returns()
func (_BuyByUsdt *BuyByUsdtTransactorSession) SetUser(_user common.Address) (*types.Transaction, error) {
	return _BuyByUsdt.Contract.SetUser(&_BuyByUsdt.TransactOpts, _user)
}

// BuyByUsdtBuyCompletedIterator is returned from FilterBuyCompleted and is used to iterate over the raw logs and unpacked data for BuyCompleted events raised by the BuyByUsdt contract.
type BuyByUsdtBuyCompletedIterator struct {
	Event *BuyByUsdtBuyCompleted // Event containing the contract specifics and raw log

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
func (it *BuyByUsdtBuyCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyByUsdtBuyCompleted)
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
		it.Event = new(BuyByUsdtBuyCompleted)
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
func (it *BuyByUsdtBuyCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyByUsdtBuyCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyByUsdtBuyCompleted represents a BuyCompleted event raised by the BuyByUsdt contract.
type BuyByUsdtBuyCompleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyCompleted is a free log retrieval operation binding the contract event 0x05ab47d0dcbc24cb46b0c25c580ee836acbbba15a6e78f4b409b6f7428dc7477.
//
// Solidity: event buyCompleted(address indexed account, uint256 amount)
func (_BuyByUsdt *BuyByUsdtFilterer) FilterBuyCompleted(opts *bind.FilterOpts, account []common.Address) (*BuyByUsdtBuyCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BuyByUsdt.contract.FilterLogs(opts, "buyCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &BuyByUsdtBuyCompletedIterator{contract: _BuyByUsdt.contract, event: "buyCompleted", logs: logs, sub: sub}, nil
}

// WatchBuyCompleted is a free log subscription operation binding the contract event 0x05ab47d0dcbc24cb46b0c25c580ee836acbbba15a6e78f4b409b6f7428dc7477.
//
// Solidity: event buyCompleted(address indexed account, uint256 amount)
func (_BuyByUsdt *BuyByUsdtFilterer) WatchBuyCompleted(opts *bind.WatchOpts, sink chan<- *BuyByUsdtBuyCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BuyByUsdt.contract.WatchLogs(opts, "buyCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyByUsdtBuyCompleted)
				if err := _BuyByUsdt.contract.UnpackLog(event, "buyCompleted", log); err != nil {
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

// ParseBuyCompleted is a log parse operation binding the contract event 0x05ab47d0dcbc24cb46b0c25c580ee836acbbba15a6e78f4b409b6f7428dc7477.
//
// Solidity: event buyCompleted(address indexed account, uint256 amount)
func (_BuyByUsdt *BuyByUsdtFilterer) ParseBuyCompleted(log types.Log) (*BuyByUsdtBuyCompleted, error) {
	event := new(BuyByUsdtBuyCompleted)
	if err := _BuyByUsdt.contract.UnpackLog(event, "buyCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
