// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deployer

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

// DeployerMetaData contains all meta data concerning the Deployer contract.
var DeployerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_collectionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_artistAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"deploy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_collectionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_artistAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"forceDeploy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_artistAddrs\",\"type\":\"address[]\"}],\"name\":\"getAllCollection\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_artistAddr\",\"type\":\"address\"}],\"name\":\"getArtistCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use DeployerMetaData.ABI instead.
var DeployerABI = DeployerMetaData.ABI

// Deployer is an auto generated Go binding around an Ethereum contract.
type Deployer struct {
	DeployerCaller     // Read-only binding to the contract
	DeployerTransactor // Write-only binding to the contract
	DeployerFilterer   // Log filterer for contract events
}

// DeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeployerSession struct {
	Contract     *Deployer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeployerCallerSession struct {
	Contract *DeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeployerTransactorSession struct {
	Contract     *DeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeployerRaw struct {
	Contract *Deployer // Generic contract binding to access the raw methods on
}

// DeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeployerCallerRaw struct {
	Contract *DeployerCaller // Generic read-only contract binding to access the raw methods on
}

// DeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeployerTransactorRaw struct {
	Contract *DeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployer creates a new instance of Deployer, bound to a specific deployed contract.
func NewDeployer(address common.Address, backend bind.ContractBackend) (*Deployer, error) {
	contract, err := bindDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deployer{DeployerCaller: DeployerCaller{contract: contract}, DeployerTransactor: DeployerTransactor{contract: contract}, DeployerFilterer: DeployerFilterer{contract: contract}}, nil
}

// NewDeployerCaller creates a new read-only instance of Deployer, bound to a specific deployed contract.
func NewDeployerCaller(address common.Address, caller bind.ContractCaller) (*DeployerCaller, error) {
	contract, err := bindDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeployerCaller{contract: contract}, nil
}

// NewDeployerTransactor creates a new write-only instance of Deployer, bound to a specific deployed contract.
func NewDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*DeployerTransactor, error) {
	contract, err := bindDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeployerTransactor{contract: contract}, nil
}

// NewDeployerFilterer creates a new log filterer instance of Deployer, bound to a specific deployed contract.
func NewDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*DeployerFilterer, error) {
	contract, err := bindDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeployerFilterer{contract: contract}, nil
}

// bindDeployer binds a generic wrapper to an already deployed contract.
func bindDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployer *DeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployer.Contract.DeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployer *DeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.Contract.DeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployer *DeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployer.Contract.DeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployer *DeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployer *DeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployer *DeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployer.Contract.contract.Transact(opts, method, params...)
}

// GetAllCollection is a free data retrieval call binding the contract method 0x1784e964.
//
// Solidity: function getAllCollection(address[] _artistAddrs) view returns(address[])
func (_Deployer *DeployerCaller) GetAllCollection(opts *bind.CallOpts, _artistAddrs []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "getAllCollection", _artistAddrs)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllCollection is a free data retrieval call binding the contract method 0x1784e964.
//
// Solidity: function getAllCollection(address[] _artistAddrs) view returns(address[])
func (_Deployer *DeployerSession) GetAllCollection(_artistAddrs []common.Address) ([]common.Address, error) {
	return _Deployer.Contract.GetAllCollection(&_Deployer.CallOpts, _artistAddrs)
}

// GetAllCollection is a free data retrieval call binding the contract method 0x1784e964.
//
// Solidity: function getAllCollection(address[] _artistAddrs) view returns(address[])
func (_Deployer *DeployerCallerSession) GetAllCollection(_artistAddrs []common.Address) ([]common.Address, error) {
	return _Deployer.Contract.GetAllCollection(&_Deployer.CallOpts, _artistAddrs)
}

// GetArtistCollection is a free data retrieval call binding the contract method 0xc11c065c.
//
// Solidity: function getArtistCollection(address _artistAddr) view returns(address)
func (_Deployer *DeployerCaller) GetArtistCollection(opts *bind.CallOpts, _artistAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "getArtistCollection", _artistAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetArtistCollection is a free data retrieval call binding the contract method 0xc11c065c.
//
// Solidity: function getArtistCollection(address _artistAddr) view returns(address)
func (_Deployer *DeployerSession) GetArtistCollection(_artistAddr common.Address) (common.Address, error) {
	return _Deployer.Contract.GetArtistCollection(&_Deployer.CallOpts, _artistAddr)
}

// GetArtistCollection is a free data retrieval call binding the contract method 0xc11c065c.
//
// Solidity: function getArtistCollection(address _artistAddr) view returns(address)
func (_Deployer *DeployerCallerSession) GetArtistCollection(_artistAddr common.Address) (common.Address, error) {
	return _Deployer.Contract.GetArtistCollection(&_Deployer.CallOpts, _artistAddr)
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() view returns(address)
func (_Deployer *DeployerCaller) GetCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "getCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() view returns(address)
func (_Deployer *DeployerSession) GetCreator() (common.Address, error) {
	return _Deployer.Contract.GetCreator(&_Deployer.CallOpts)
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() view returns(address)
func (_Deployer *DeployerCallerSession) GetCreator() (common.Address, error) {
	return _Deployer.Contract.GetCreator(&_Deployer.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Deployer *DeployerCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Deployer *DeployerSession) GetOwner() (common.Address, error) {
	return _Deployer.Contract.GetOwner(&_Deployer.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Deployer *DeployerCallerSession) GetOwner() (common.Address, error) {
	return _Deployer.Contract.GetOwner(&_Deployer.CallOpts)
}

// Deploy is a paid mutator transaction binding the contract method 0x97799c6b.
//
// Solidity: function deploy(string _collectionName, string _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) returns()
func (_Deployer *DeployerTransactor) Deploy(opts *bind.TransactOpts, _collectionName string, _baseURI string, _artistAddr common.Address, _royaltyBasisPoints *big.Int, _currency common.Address) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "deploy", _collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency)
}

// Deploy is a paid mutator transaction binding the contract method 0x97799c6b.
//
// Solidity: function deploy(string _collectionName, string _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) returns()
func (_Deployer *DeployerSession) Deploy(_collectionName string, _baseURI string, _artistAddr common.Address, _royaltyBasisPoints *big.Int, _currency common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.Deploy(&_Deployer.TransactOpts, _collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency)
}

// Deploy is a paid mutator transaction binding the contract method 0x97799c6b.
//
// Solidity: function deploy(string _collectionName, string _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) returns()
func (_Deployer *DeployerTransactorSession) Deploy(_collectionName string, _baseURI string, _artistAddr common.Address, _royaltyBasisPoints *big.Int, _currency common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.Deploy(&_Deployer.TransactOpts, _collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency)
}

// ForceDeploy is a paid mutator transaction binding the contract method 0x50342a05.
//
// Solidity: function forceDeploy(string _collectionName, string _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) returns()
func (_Deployer *DeployerTransactor) ForceDeploy(opts *bind.TransactOpts, _collectionName string, _baseURI string, _artistAddr common.Address, _royaltyBasisPoints *big.Int, _currency common.Address) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "forceDeploy", _collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency)
}

// ForceDeploy is a paid mutator transaction binding the contract method 0x50342a05.
//
// Solidity: function forceDeploy(string _collectionName, string _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) returns()
func (_Deployer *DeployerSession) ForceDeploy(_collectionName string, _baseURI string, _artistAddr common.Address, _royaltyBasisPoints *big.Int, _currency common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.ForceDeploy(&_Deployer.TransactOpts, _collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency)
}

// ForceDeploy is a paid mutator transaction binding the contract method 0x50342a05.
//
// Solidity: function forceDeploy(string _collectionName, string _baseURI, address _artistAddr, uint256 _royaltyBasisPoints, address _currency) returns()
func (_Deployer *DeployerTransactorSession) ForceDeploy(_collectionName string, _baseURI string, _artistAddr common.Address, _royaltyBasisPoints *big.Int, _currency common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.ForceDeploy(&_Deployer.TransactOpts, _collectionName, _baseURI, _artistAddr, _royaltyBasisPoints, _currency)
}

// DeployerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Deployer contract.
type DeployerOwnershipTransferredIterator struct {
	Event *DeployerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DeployerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployerOwnershipTransferred)
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
		it.Event = new(DeployerOwnershipTransferred)
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
func (it *DeployerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployerOwnershipTransferred represents a OwnershipTransferred event raised by the Deployer contract.
type DeployerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Deployer *DeployerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DeployerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Deployer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DeployerOwnershipTransferredIterator{contract: _Deployer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Deployer *DeployerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DeployerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Deployer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployerOwnershipTransferred)
				if err := _Deployer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Deployer *DeployerFilterer) ParseOwnershipTransferred(log types.Log) (*DeployerOwnershipTransferred, error) {
	event := new(DeployerOwnershipTransferred)
	if err := _Deployer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
