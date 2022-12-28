// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package artist_contract

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

// ArtistContractMetaData contains all meta data concerning the ArtistContract contract.
var ArtistContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_artistAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalRoyaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBasisPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOperator\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_saleCost\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newuri\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newName\",\"type\":\"string\"}],\"name\":\"setContractName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBasisPoints\",\"type\":\"uint256\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"setCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currenctyBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_royaltyBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setGlobalRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ArtistContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ArtistContractMetaData.ABI instead.
var ArtistContractABI = ArtistContractMetaData.ABI

// ArtistContract is an auto generated Go binding around an Ethereum contract.
type ArtistContract struct {
	ArtistContractCaller     // Read-only binding to the contract
	ArtistContractTransactor // Write-only binding to the contract
	ArtistContractFilterer   // Log filterer for contract events
}

// ArtistContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArtistContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArtistContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArtistContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArtistContractSession struct {
	Contract     *ArtistContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtistContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArtistContractCallerSession struct {
	Contract *ArtistContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ArtistContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArtistContractTransactorSession struct {
	Contract     *ArtistContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ArtistContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArtistContractRaw struct {
	Contract *ArtistContract // Generic contract binding to access the raw methods on
}

// ArtistContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArtistContractCallerRaw struct {
	Contract *ArtistContractCaller // Generic read-only contract binding to access the raw methods on
}

// ArtistContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArtistContractTransactorRaw struct {
	Contract *ArtistContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArtistContract creates a new instance of ArtistContract, bound to a specific deployed contract.
func NewArtistContract(address common.Address, backend bind.ContractBackend) (*ArtistContract, error) {
	contract, err := bindArtistContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArtistContract{ArtistContractCaller: ArtistContractCaller{contract: contract}, ArtistContractTransactor: ArtistContractTransactor{contract: contract}, ArtistContractFilterer: ArtistContractFilterer{contract: contract}}, nil
}

// NewArtistContractCaller creates a new read-only instance of ArtistContract, bound to a specific deployed contract.
func NewArtistContractCaller(address common.Address, caller bind.ContractCaller) (*ArtistContractCaller, error) {
	contract, err := bindArtistContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArtistContractCaller{contract: contract}, nil
}

// NewArtistContractTransactor creates a new write-only instance of ArtistContract, bound to a specific deployed contract.
func NewArtistContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ArtistContractTransactor, error) {
	contract, err := bindArtistContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArtistContractTransactor{contract: contract}, nil
}

// NewArtistContractFilterer creates a new log filterer instance of ArtistContract, bound to a specific deployed contract.
func NewArtistContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ArtistContractFilterer, error) {
	contract, err := bindArtistContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArtistContractFilterer{contract: contract}, nil
}

// bindArtistContract binds a generic wrapper to an already deployed contract.
func bindArtistContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtistContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtistContract *ArtistContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArtistContract.Contract.ArtistContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtistContract *ArtistContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtistContract.Contract.ArtistContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtistContract *ArtistContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtistContract.Contract.ArtistContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtistContract *ArtistContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArtistContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtistContract *ArtistContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtistContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtistContract *ArtistContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtistContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_ArtistContract *ArtistContractCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "balanceOf", _owner, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_ArtistContract *ArtistContractSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _ArtistContract.Contract.BalanceOf(&_ArtistContract.CallOpts, _owner, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_ArtistContract *ArtistContractCallerSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _ArtistContract.Contract.BalanceOf(&_ArtistContract.CallOpts, _owner, _id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_ArtistContract *ArtistContractCaller) BalanceOfBatch(opts *bind.CallOpts, _owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "balanceOfBatch", _owners, _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_ArtistContract *ArtistContractSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _ArtistContract.Contract.BalanceOfBatch(&_ArtistContract.CallOpts, _owners, _ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_ArtistContract *ArtistContractCallerSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _ArtistContract.Contract.BalanceOfBatch(&_ArtistContract.CallOpts, _owners, _ids)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_ArtistContract *ArtistContractCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_ArtistContract *ArtistContractSession) BaseURI() (string, error) {
	return _ArtistContract.Contract.BaseURI(&_ArtistContract.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_ArtistContract *ArtistContractCallerSession) BaseURI() (string, error) {
	return _ArtistContract.Contract.BaseURI(&_ArtistContract.CallOpts)
}

// CurrenctyBalance is a free data retrieval call binding the contract method 0x4c6af17c.
//
// Solidity: function currenctyBalance() view returns(uint256)
func (_ArtistContract *ArtistContractCaller) CurrenctyBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "currenctyBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrenctyBalance is a free data retrieval call binding the contract method 0x4c6af17c.
//
// Solidity: function currenctyBalance() view returns(uint256)
func (_ArtistContract *ArtistContractSession) CurrenctyBalance() (*big.Int, error) {
	return _ArtistContract.Contract.CurrenctyBalance(&_ArtistContract.CallOpts)
}

// CurrenctyBalance is a free data retrieval call binding the contract method 0x4c6af17c.
//
// Solidity: function currenctyBalance() view returns(uint256)
func (_ArtistContract *ArtistContractCallerSession) CurrenctyBalance() (*big.Int, error) {
	return _ArtistContract.Contract.CurrenctyBalance(&_ArtistContract.CallOpts)
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() view returns(address)
func (_ArtistContract *ArtistContractCaller) GetCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "getCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() view returns(address)
func (_ArtistContract *ArtistContractSession) GetCreator() (common.Address, error) {
	return _ArtistContract.Contract.GetCreator(&_ArtistContract.CallOpts)
}

// GetCreator is a free data retrieval call binding the contract method 0x0ee2cb10.
//
// Solidity: function getCreator() view returns(address)
func (_ArtistContract *ArtistContractCallerSession) GetCreator() (common.Address, error) {
	return _ArtistContract.Contract.GetCreator(&_ArtistContract.CallOpts)
}

// GetCurrency is a free data retrieval call binding the contract method 0x6945c1fd.
//
// Solidity: function getCurrency() view returns(address)
func (_ArtistContract *ArtistContractCaller) GetCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "getCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrency is a free data retrieval call binding the contract method 0x6945c1fd.
//
// Solidity: function getCurrency() view returns(address)
func (_ArtistContract *ArtistContractSession) GetCurrency() (common.Address, error) {
	return _ArtistContract.Contract.GetCurrency(&_ArtistContract.CallOpts)
}

// GetCurrency is a free data retrieval call binding the contract method 0x6945c1fd.
//
// Solidity: function getCurrency() view returns(address)
func (_ArtistContract *ArtistContractCallerSession) GetCurrency() (common.Address, error) {
	return _ArtistContract.Contract.GetCurrency(&_ArtistContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ArtistContract *ArtistContractCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ArtistContract *ArtistContractSession) GetOwner() (common.Address, error) {
	return _ArtistContract.Contract.GetOwner(&_ArtistContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ArtistContract *ArtistContractCallerSession) GetOwner() (common.Address, error) {
	return _ArtistContract.Contract.GetOwner(&_ArtistContract.CallOpts)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xc457fb37.
//
// Solidity: function getTokenPrice(uint256 _id) view returns(uint256)
func (_ArtistContract *ArtistContractCaller) GetTokenPrice(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "getTokenPrice", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPrice is a free data retrieval call binding the contract method 0xc457fb37.
//
// Solidity: function getTokenPrice(uint256 _id) view returns(uint256)
func (_ArtistContract *ArtistContractSession) GetTokenPrice(_id *big.Int) (*big.Int, error) {
	return _ArtistContract.Contract.GetTokenPrice(&_ArtistContract.CallOpts, _id)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xc457fb37.
//
// Solidity: function getTokenPrice(uint256 _id) view returns(uint256)
func (_ArtistContract *ArtistContractCallerSession) GetTokenPrice(_id *big.Int) (*big.Int, error) {
	return _ArtistContract.Contract.GetTokenPrice(&_ArtistContract.CallOpts, _id)
}

// GlobalRoyaltyInfo is a free data retrieval call binding the contract method 0xc9823cc6.
//
// Solidity: function globalRoyaltyInfo() view returns(address receiver, uint256 feeBasisPoints)
func (_ArtistContract *ArtistContractCaller) GlobalRoyaltyInfo(opts *bind.CallOpts) (struct {
	Receiver       common.Address
	FeeBasisPoints *big.Int
}, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "globalRoyaltyInfo")

	outstruct := new(struct {
		Receiver       common.Address
		FeeBasisPoints *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FeeBasisPoints = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GlobalRoyaltyInfo is a free data retrieval call binding the contract method 0xc9823cc6.
//
// Solidity: function globalRoyaltyInfo() view returns(address receiver, uint256 feeBasisPoints)
func (_ArtistContract *ArtistContractSession) GlobalRoyaltyInfo() (struct {
	Receiver       common.Address
	FeeBasisPoints *big.Int
}, error) {
	return _ArtistContract.Contract.GlobalRoyaltyInfo(&_ArtistContract.CallOpts)
}

// GlobalRoyaltyInfo is a free data retrieval call binding the contract method 0xc9823cc6.
//
// Solidity: function globalRoyaltyInfo() view returns(address receiver, uint256 feeBasisPoints)
func (_ArtistContract *ArtistContractCallerSession) GlobalRoyaltyInfo() (struct {
	Receiver       common.Address
	FeeBasisPoints *big.Int
}, error) {
	return _ArtistContract.Contract.GlobalRoyaltyInfo(&_ArtistContract.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_ArtistContract *ArtistContractCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_ArtistContract *ArtistContractSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ArtistContract.Contract.IsApprovedForAll(&_ArtistContract.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_ArtistContract *ArtistContractCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ArtistContract.Contract.IsApprovedForAll(&_ArtistContract.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtistContract *ArtistContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtistContract *ArtistContractSession) Name() (string, error) {
	return _ArtistContract.Contract.Name(&_ArtistContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtistContract *ArtistContractCallerSession) Name() (string, error) {
	return _ArtistContract.Contract.Name(&_ArtistContract.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _saleCost) view returns(address receiver, uint256 royaltyAmount)
func (_ArtistContract *ArtistContractCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, _saleCost *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "royaltyInfo", arg0, _saleCost)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _saleCost) view returns(address receiver, uint256 royaltyAmount)
func (_ArtistContract *ArtistContractSession) RoyaltyInfo(arg0 *big.Int, _saleCost *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _ArtistContract.Contract.RoyaltyInfo(&_ArtistContract.CallOpts, arg0, _saleCost)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _saleCost) view returns(address receiver, uint256 royaltyAmount)
func (_ArtistContract *ArtistContractCallerSession) RoyaltyInfo(arg0 *big.Int, _saleCost *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _ArtistContract.Contract.RoyaltyInfo(&_ArtistContract.CallOpts, arg0, _saleCost)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_ArtistContract *ArtistContractCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_ArtistContract *ArtistContractSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ArtistContract.Contract.SupportsInterface(&_ArtistContract.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_ArtistContract *ArtistContractCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ArtistContract.Contract.SupportsInterface(&_ArtistContract.CallOpts, _interfaceID)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_ArtistContract *ArtistContractCaller) Uri(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _ArtistContract.contract.Call(opts, &out, "uri", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_ArtistContract *ArtistContractSession) Uri(_id *big.Int) (string, error) {
	return _ArtistContract.Contract.Uri(&_ArtistContract.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_ArtistContract *ArtistContractCallerSession) Uri(_id *big.Int) (string, error) {
	return _ArtistContract.Contract.Uri(&_ArtistContract.CallOpts, _id)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x20ec271b.
//
// Solidity: function batchBurn(uint256[] _ids, uint256[] _amounts) returns()
func (_ArtistContract *ArtistContractTransactor) BatchBurn(opts *bind.TransactOpts, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "batchBurn", _ids, _amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x20ec271b.
//
// Solidity: function batchBurn(uint256[] _ids, uint256[] _amounts) returns()
func (_ArtistContract *ArtistContractSession) BatchBurn(_ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.BatchBurn(&_ArtistContract.TransactOpts, _ids, _amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x20ec271b.
//
// Solidity: function batchBurn(uint256[] _ids, uint256[] _amounts) returns()
func (_ArtistContract *ArtistContractTransactorSession) BatchBurn(_ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.BatchBurn(&_ArtistContract.TransactOpts, _ids, _amounts)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 _id, uint256 _amount) returns()
func (_ArtistContract *ArtistContractTransactor) Burn(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "burn", _id, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 _id, uint256 _amount) returns()
func (_ArtistContract *ArtistContractSession) Burn(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.Burn(&_ArtistContract.TransactOpts, _id, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 _id, uint256 _amount) returns()
func (_ArtistContract *ArtistContractTransactorSession) Burn(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.Burn(&_ArtistContract.TransactOpts, _id, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0x1efd73a5.
//
// Solidity: function buy(uint256 _id, uint256 _amount, bytes _data) returns()
func (_ArtistContract *ArtistContractTransactor) Buy(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "buy", _id, _amount, _data)
}

// Buy is a paid mutator transaction binding the contract method 0x1efd73a5.
//
// Solidity: function buy(uint256 _id, uint256 _amount, bytes _data) returns()
func (_ArtistContract *ArtistContractSession) Buy(_id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.Buy(&_ArtistContract.TransactOpts, _id, _amount, _data)
}

// Buy is a paid mutator transaction binding the contract method 0x1efd73a5.
//
// Solidity: function buy(uint256 _id, uint256 _amount, bytes _data) returns()
func (_ArtistContract *ArtistContractTransactorSession) Buy(_id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.Buy(&_ArtistContract.TransactOpts, _id, _amount, _data)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 id, uint256 amount, uint256 price, bytes data) returns()
func (_ArtistContract *ArtistContractTransactor) Mint(opts *bind.TransactOpts, id *big.Int, amount *big.Int, price *big.Int, data []byte) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "mint", id, amount, price, data)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 id, uint256 amount, uint256 price, bytes data) returns()
func (_ArtistContract *ArtistContractSession) Mint(id *big.Int, amount *big.Int, price *big.Int, data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.Mint(&_ArtistContract.TransactOpts, id, amount, price, data)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 id, uint256 amount, uint256 price, bytes data) returns()
func (_ArtistContract *ArtistContractTransactorSession) Mint(id *big.Int, amount *big.Int, price *big.Int, data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.Mint(&_ArtistContract.TransactOpts, id, amount, price, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x87fb8e48.
//
// Solidity: function mintBatch(uint256[] ids, uint256[] amounts, uint256[] prices, bytes data) returns()
func (_ArtistContract *ArtistContractTransactor) MintBatch(opts *bind.TransactOpts, ids []*big.Int, amounts []*big.Int, prices []*big.Int, data []byte) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "mintBatch", ids, amounts, prices, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x87fb8e48.
//
// Solidity: function mintBatch(uint256[] ids, uint256[] amounts, uint256[] prices, bytes data) returns()
func (_ArtistContract *ArtistContractSession) MintBatch(ids []*big.Int, amounts []*big.Int, prices []*big.Int, data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.MintBatch(&_ArtistContract.TransactOpts, ids, amounts, prices, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x87fb8e48.
//
// Solidity: function mintBatch(uint256[] ids, uint256[] amounts, uint256[] prices, bytes data) returns()
func (_ArtistContract *ArtistContractTransactorSession) MintBatch(ids []*big.Int, amounts []*big.Int, prices []*big.Int, data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.MintBatch(&_ArtistContract.TransactOpts, ids, amounts, prices, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts, bytes _data) returns()
func (_ArtistContract *ArtistContractTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "safeBatchTransferFrom", _from, _to, _ids, _amounts, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts, bytes _data) returns()
func (_ArtistContract *ArtistContractSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.SafeBatchTransferFrom(&_ArtistContract.TransactOpts, _from, _to, _ids, _amounts, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts, bytes _data) returns()
func (_ArtistContract *ArtistContractTransactorSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.SafeBatchTransferFrom(&_ArtistContract.TransactOpts, _from, _to, _ids, _amounts, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount, bytes _data) returns()
func (_ArtistContract *ArtistContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "safeTransferFrom", _from, _to, _id, _amount, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount, bytes _data) returns()
func (_ArtistContract *ArtistContractSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.SafeTransferFrom(&_ArtistContract.TransactOpts, _from, _to, _id, _amount, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount, bytes _data) returns()
func (_ArtistContract *ArtistContractTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.SafeTransferFrom(&_ArtistContract.TransactOpts, _from, _to, _id, _amount, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ArtistContract *ArtistContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ArtistContract *ArtistContractSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetApprovalForAll(&_ArtistContract.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_ArtistContract *ArtistContractTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetApprovalForAll(&_ArtistContract.TransactOpts, _operator, _approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newuri) returns()
func (_ArtistContract *ArtistContractTransactor) SetBaseURI(opts *bind.TransactOpts, _newuri string) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "setBaseURI", _newuri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newuri) returns()
func (_ArtistContract *ArtistContractSession) SetBaseURI(_newuri string) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetBaseURI(&_ArtistContract.TransactOpts, _newuri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newuri) returns()
func (_ArtistContract *ArtistContractTransactorSession) SetBaseURI(_newuri string) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetBaseURI(&_ArtistContract.TransactOpts, _newuri)
}

// SetContractName is a paid mutator transaction binding the contract method 0x0b5ee006.
//
// Solidity: function setContractName(string _newName) returns()
func (_ArtistContract *ArtistContractTransactor) SetContractName(opts *bind.TransactOpts, _newName string) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "setContractName", _newName)
}

// SetContractName is a paid mutator transaction binding the contract method 0x0b5ee006.
//
// Solidity: function setContractName(string _newName) returns()
func (_ArtistContract *ArtistContractSession) SetContractName(_newName string) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetContractName(&_ArtistContract.TransactOpts, _newName)
}

// SetContractName is a paid mutator transaction binding the contract method 0x0b5ee006.
//
// Solidity: function setContractName(string _newName) returns()
func (_ArtistContract *ArtistContractTransactorSession) SetContractName(_newName string) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetContractName(&_ArtistContract.TransactOpts, _newName)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x2f84c391.
//
// Solidity: function setCurrency(address _currency) returns()
func (_ArtistContract *ArtistContractTransactor) SetCurrency(opts *bind.TransactOpts, _currency common.Address) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "setCurrency", _currency)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x2f84c391.
//
// Solidity: function setCurrency(address _currency) returns()
func (_ArtistContract *ArtistContractSession) SetCurrency(_currency common.Address) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetCurrency(&_ArtistContract.TransactOpts, _currency)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x2f84c391.
//
// Solidity: function setCurrency(address _currency) returns()
func (_ArtistContract *ArtistContractTransactorSession) SetCurrency(_currency common.Address) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetCurrency(&_ArtistContract.TransactOpts, _currency)
}

// SetGlobalRoyaltyInfo is a paid mutator transaction binding the contract method 0x045a1e99.
//
// Solidity: function setGlobalRoyaltyInfo(uint256 _royaltyBasisPoints) returns()
func (_ArtistContract *ArtistContractTransactor) SetGlobalRoyaltyInfo(opts *bind.TransactOpts, _royaltyBasisPoints *big.Int) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "setGlobalRoyaltyInfo", _royaltyBasisPoints)
}

// SetGlobalRoyaltyInfo is a paid mutator transaction binding the contract method 0x045a1e99.
//
// Solidity: function setGlobalRoyaltyInfo(uint256 _royaltyBasisPoints) returns()
func (_ArtistContract *ArtistContractSession) SetGlobalRoyaltyInfo(_royaltyBasisPoints *big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetGlobalRoyaltyInfo(&_ArtistContract.TransactOpts, _royaltyBasisPoints)
}

// SetGlobalRoyaltyInfo is a paid mutator transaction binding the contract method 0x045a1e99.
//
// Solidity: function setGlobalRoyaltyInfo(uint256 _royaltyBasisPoints) returns()
func (_ArtistContract *ArtistContractTransactorSession) SetGlobalRoyaltyInfo(_royaltyBasisPoints *big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetGlobalRoyaltyInfo(&_ArtistContract.TransactOpts, _royaltyBasisPoints)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 _id, uint256 _price) returns()
func (_ArtistContract *ArtistContractTransactor) SetPrice(opts *bind.TransactOpts, _id *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "setPrice", _id, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 _id, uint256 _price) returns()
func (_ArtistContract *ArtistContractSession) SetPrice(_id *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetPrice(&_ArtistContract.TransactOpts, _id, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 _id, uint256 _price) returns()
func (_ArtistContract *ArtistContractTransactorSession) SetPrice(_id *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.SetPrice(&_ArtistContract.TransactOpts, _id, _price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x0a276680.
//
// Solidity: function transferOwnership(address _newOwner, uint256 _royaltyBasisPoints) returns()
func (_ArtistContract *ArtistContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address, _royaltyBasisPoints *big.Int) (*types.Transaction, error) {
	return _ArtistContract.contract.Transact(opts, "transferOwnership", _newOwner, _royaltyBasisPoints)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x0a276680.
//
// Solidity: function transferOwnership(address _newOwner, uint256 _royaltyBasisPoints) returns()
func (_ArtistContract *ArtistContractSession) TransferOwnership(_newOwner common.Address, _royaltyBasisPoints *big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.TransferOwnership(&_ArtistContract.TransactOpts, _newOwner, _royaltyBasisPoints)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x0a276680.
//
// Solidity: function transferOwnership(address _newOwner, uint256 _royaltyBasisPoints) returns()
func (_ArtistContract *ArtistContractTransactorSession) TransferOwnership(_newOwner common.Address, _royaltyBasisPoints *big.Int) (*types.Transaction, error) {
	return _ArtistContract.Contract.TransferOwnership(&_ArtistContract.TransactOpts, _newOwner, _royaltyBasisPoints)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ArtistContract *ArtistContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ArtistContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ArtistContract *ArtistContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.Fallback(&_ArtistContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ArtistContract *ArtistContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ArtistContract.Contract.Fallback(&_ArtistContract.TransactOpts, calldata)
}

// ArtistContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ArtistContract contract.
type ArtistContractApprovalForAllIterator struct {
	Event *ArtistContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ArtistContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistContractApprovalForAll)
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
		it.Event = new(ArtistContractApprovalForAll)
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
func (it *ArtistContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistContractApprovalForAll represents a ApprovalForAll event raised by the ArtistContract contract.
type ArtistContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ArtistContract *ArtistContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ArtistContractApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ArtistContract.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ArtistContractApprovalForAllIterator{contract: _ArtistContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ArtistContract *ArtistContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ArtistContractApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ArtistContract.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistContractApprovalForAll)
				if err := _ArtistContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_ArtistContract *ArtistContractFilterer) ParseApprovalForAll(log types.Log) (*ArtistContractApprovalForAll, error) {
	event := new(ArtistContractApprovalForAll)
	if err := _ArtistContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ArtistContract contract.
type ArtistContractOwnershipTransferredIterator struct {
	Event *ArtistContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArtistContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistContractOwnershipTransferred)
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
		it.Event = new(ArtistContractOwnershipTransferred)
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
func (it *ArtistContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistContractOwnershipTransferred represents a OwnershipTransferred event raised by the ArtistContract contract.
type ArtistContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArtistContract *ArtistContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArtistContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArtistContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArtistContractOwnershipTransferredIterator{contract: _ArtistContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArtistContract *ArtistContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArtistContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArtistContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistContractOwnershipTransferred)
				if err := _ArtistContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ArtistContract *ArtistContractFilterer) ParseOwnershipTransferred(log types.Log) (*ArtistContractOwnershipTransferred, error) {
	event := new(ArtistContractOwnershipTransferred)
	if err := _ArtistContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistContractTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ArtistContract contract.
type ArtistContractTransferBatchIterator struct {
	Event *ArtistContractTransferBatch // Event containing the contract specifics and raw log

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
func (it *ArtistContractTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistContractTransferBatch)
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
		it.Event = new(ArtistContractTransferBatch)
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
func (it *ArtistContractTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistContractTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistContractTransferBatch represents a TransferBatch event raised by the ArtistContract contract.
type ArtistContractTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts)
func (_ArtistContract *ArtistContractFilterer) FilterTransferBatch(opts *bind.FilterOpts, _operator []common.Address, _from []common.Address, _to []common.Address) (*ArtistContractTransferBatchIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ArtistContract.contract.FilterLogs(opts, "TransferBatch", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ArtistContractTransferBatchIterator{contract: _ArtistContract.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts)
func (_ArtistContract *ArtistContractFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ArtistContractTransferBatch, _operator []common.Address, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ArtistContract.contract.WatchLogs(opts, "TransferBatch", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistContractTransferBatch)
				if err := _ArtistContract.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts)
func (_ArtistContract *ArtistContractFilterer) ParseTransferBatch(log types.Log) (*ArtistContractTransferBatch, error) {
	event := new(ArtistContractTransferBatch)
	if err := _ArtistContract.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistContractTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ArtistContract contract.
type ArtistContractTransferSingleIterator struct {
	Event *ArtistContractTransferSingle // Event containing the contract specifics and raw log

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
func (it *ArtistContractTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistContractTransferSingle)
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
		it.Event = new(ArtistContractTransferSingle)
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
func (it *ArtistContractTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistContractTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistContractTransferSingle represents a TransferSingle event raised by the ArtistContract contract.
type ArtistContractTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func (_ArtistContract *ArtistContractFilterer) FilterTransferSingle(opts *bind.FilterOpts, _operator []common.Address, _from []common.Address, _to []common.Address) (*ArtistContractTransferSingleIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ArtistContract.contract.FilterLogs(opts, "TransferSingle", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ArtistContractTransferSingleIterator{contract: _ArtistContract.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func (_ArtistContract *ArtistContractFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ArtistContractTransferSingle, _operator []common.Address, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ArtistContract.contract.WatchLogs(opts, "TransferSingle", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistContractTransferSingle)
				if err := _ArtistContract.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func (_ArtistContract *ArtistContractFilterer) ParseTransferSingle(log types.Log) (*ArtistContractTransferSingle, error) {
	event := new(ArtistContractTransferSingle)
	if err := _ArtistContract.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistContractURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ArtistContract contract.
type ArtistContractURIIterator struct {
	Event *ArtistContractURI // Event containing the contract specifics and raw log

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
func (it *ArtistContractURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistContractURI)
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
		it.Event = new(ArtistContractURI)
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
func (it *ArtistContractURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistContractURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistContractURI represents a URI event raised by the ArtistContract contract.
type ArtistContractURI struct {
	Uri string
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string _uri, uint256 indexed _id)
func (_ArtistContract *ArtistContractFilterer) FilterURI(opts *bind.FilterOpts, _id []*big.Int) (*ArtistContractURIIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ArtistContract.contract.FilterLogs(opts, "URI", _idRule)
	if err != nil {
		return nil, err
	}
	return &ArtistContractURIIterator{contract: _ArtistContract.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string _uri, uint256 indexed _id)
func (_ArtistContract *ArtistContractFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ArtistContractURI, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ArtistContract.contract.WatchLogs(opts, "URI", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistContractURI)
				if err := _ArtistContract.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string _uri, uint256 indexed _id)
func (_ArtistContract *ArtistContractFilterer) ParseURI(log types.Log) (*ArtistContractURI, error) {
	event := new(ArtistContractURI)
	if err := _ArtistContract.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
