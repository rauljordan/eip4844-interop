// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// DataHashesReaderMetaData contains all meta data concerning the DataHashesReader contract.
var DataHashesReaderMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"getDataHashes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"payable\":false,\"type\":\"function\"}]",
	Bin: "0x604d600d600039604d6000f3fe60003560e01c63e83a2d828103604b5760005b600115603757804980602357506037565b806040602084020152600182019150506012565b602060005280602052604060208202016000f35b50",
}

// DataHashesReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use DataHashesReaderMetaData.ABI instead.
var DataHashesReaderABI = DataHashesReaderMetaData.ABI

// DataHashesReaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DataHashesReaderMetaData.Bin instead.
var DataHashesReaderBin = DataHashesReaderMetaData.Bin

// DeployDataHashesReader deploys a new Ethereum contract, binding an instance of DataHashesReader to it.
func DeployDataHashesReader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataHashesReader, error) {
	parsed, err := DataHashesReaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DataHashesReaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataHashesReader{DataHashesReaderCaller: DataHashesReaderCaller{contract: contract}, DataHashesReaderTransactor: DataHashesReaderTransactor{contract: contract}, DataHashesReaderFilterer: DataHashesReaderFilterer{contract: contract}}, nil
}

// DataHashesReader is an auto generated Go binding around an Ethereum contract.
type DataHashesReader struct {
	DataHashesReaderCaller     // Read-only binding to the contract
	DataHashesReaderTransactor // Write-only binding to the contract
	DataHashesReaderFilterer   // Log filterer for contract events
}

// DataHashesReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataHashesReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataHashesReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataHashesReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataHashesReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataHashesReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataHashesReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataHashesReaderSession struct {
	Contract     *DataHashesReader // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataHashesReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataHashesReaderCallerSession struct {
	Contract *DataHashesReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DataHashesReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataHashesReaderTransactorSession struct {
	Contract     *DataHashesReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DataHashesReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataHashesReaderRaw struct {
	Contract *DataHashesReader // Generic contract binding to access the raw methods on
}

// DataHashesReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataHashesReaderCallerRaw struct {
	Contract *DataHashesReaderCaller // Generic read-only contract binding to access the raw methods on
}

// DataHashesReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataHashesReaderTransactorRaw struct {
	Contract *DataHashesReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataHashesReader creates a new instance of DataHashesReader, bound to a specific deployed contract.
func NewDataHashesReader(address common.Address, backend bind.ContractBackend) (*DataHashesReader, error) {
	contract, err := bindDataHashesReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataHashesReader{DataHashesReaderCaller: DataHashesReaderCaller{contract: contract}, DataHashesReaderTransactor: DataHashesReaderTransactor{contract: contract}, DataHashesReaderFilterer: DataHashesReaderFilterer{contract: contract}}, nil
}

// NewDataHashesReaderCaller creates a new read-only instance of DataHashesReader, bound to a specific deployed contract.
func NewDataHashesReaderCaller(address common.Address, caller bind.ContractCaller) (*DataHashesReaderCaller, error) {
	contract, err := bindDataHashesReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataHashesReaderCaller{contract: contract}, nil
}

// NewDataHashesReaderTransactor creates a new write-only instance of DataHashesReader, bound to a specific deployed contract.
func NewDataHashesReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*DataHashesReaderTransactor, error) {
	contract, err := bindDataHashesReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataHashesReaderTransactor{contract: contract}, nil
}

// NewDataHashesReaderFilterer creates a new log filterer instance of DataHashesReader, bound to a specific deployed contract.
func NewDataHashesReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*DataHashesReaderFilterer, error) {
	contract, err := bindDataHashesReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataHashesReaderFilterer{contract: contract}, nil
}

// bindDataHashesReader binds a generic wrapper to an already deployed contract.
func bindDataHashesReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataHashesReaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataHashesReader *DataHashesReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataHashesReader.Contract.DataHashesReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataHashesReader *DataHashesReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataHashesReader.Contract.DataHashesReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataHashesReader *DataHashesReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataHashesReader.Contract.DataHashesReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataHashesReader *DataHashesReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataHashesReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataHashesReader *DataHashesReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataHashesReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataHashesReader *DataHashesReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataHashesReader.Contract.contract.Transact(opts, method, params...)
}

// GetDataHashes is a free data retrieval call binding the contract method 0xe83a2d82.
//
// Solidity: function getDataHashes() view returns(bytes32[])
func (_DataHashesReader *DataHashesReaderCaller) GetDataHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _DataHashesReader.contract.Call(opts, &out, "getDataHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDataHashes is a free data retrieval call binding the contract method 0xe83a2d82.
//
// Solidity: function getDataHashes() view returns(bytes32[])
func (_DataHashesReader *DataHashesReaderSession) GetDataHashes() ([][32]byte, error) {
	return _DataHashesReader.Contract.GetDataHashes(&_DataHashesReader.CallOpts)
}

// GetDataHashes is a free data retrieval call binding the contract method 0xe83a2d82.
//
// Solidity: function getDataHashes() view returns(bytes32[])
func (_DataHashesReader *DataHashesReaderCallerSession) GetDataHashes() ([][32]byte, error) {
	return _DataHashesReader.Contract.GetDataHashes(&_DataHashesReader.CallOpts)
}
