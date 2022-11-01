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

// ToySequencerInboxMetaData contains all meta data concerning the ToySequencerInbox contract.
var ToySequencerInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dataHashReader\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numHashes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"name\":\"DataHashesParsed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"submitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610550380380610550833981810160405281019061003291906100db565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100a88261007d565b9050919050565b6100b88161009d565b81146100c357600080fd5b50565b6000815190506100d5816100af565b92915050565b6000602082840312156100f1576100f0610078565b5b60006100ff848285016100c6565b91505092915050565b610439806101176000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80634611538314610030575b600080fd5b61003861003a565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e83a2d826040518163ffffffff1660e01b8152600401600060405180830381865afa1580156100a8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906100d191906102b3565b90507fca7600c3e9508b8063c33ff24d5b042ead3f1d362a4f92868a3d5f8af14cd86b8151826040516101059291906103d3565b60405180910390a150565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61017282610129565b810181811067ffffffffffffffff821117156101915761019061013a565b5b80604052505050565b60006101a4610110565b90506101b08282610169565b919050565b600067ffffffffffffffff8211156101d0576101cf61013a565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b6101f9816101e6565b811461020457600080fd5b50565b600081519050610216816101f0565b92915050565b600061022f61022a846101b5565b61019a565b90508083825260208201905060208402830185811115610252576102516101e1565b5b835b8181101561027b57806102678882610207565b845260208401935050602081019050610254565b5050509392505050565b600082601f83011261029a57610299610124565b5b81516102aa84826020860161021c565b91505092915050565b6000602082840312156102c9576102c861011a565b5b600082015167ffffffffffffffff8111156102e7576102e661011f565b5b6102f384828501610285565b91505092915050565b6000819050919050565b61030f816102fc565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61034a816101e6565b82525050565b600061035c8383610341565b60208301905092915050565b6000602082019050919050565b600061038082610315565b61038a8185610320565b935061039583610331565b8060005b838110156103c65781516103ad8882610350565b97506103b883610368565b925050600181019050610399565b5085935050505092915050565b60006040820190506103e86000830185610306565b81810360208301526103fa8184610375565b9050939250505056fea2646970667358221220d0e0d21a4b9f811265e963e88fd0f57f913eba57ff66f5e2a51d546f4445345564736f6c63430008110033",
}

// ToySequencerInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use ToySequencerInboxMetaData.ABI instead.
var ToySequencerInboxABI = ToySequencerInboxMetaData.ABI

// ToySequencerInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ToySequencerInboxMetaData.Bin instead.
var ToySequencerInboxBin = ToySequencerInboxMetaData.Bin

// DeployToySequencerInbox deploys a new Ethereum contract, binding an instance of ToySequencerInbox to it.
func DeployToySequencerInbox(auth *bind.TransactOpts, backend bind.ContractBackend, _dataHashReader common.Address) (common.Address, *types.Transaction, *ToySequencerInbox, error) {
	parsed, err := ToySequencerInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ToySequencerInboxBin), backend, _dataHashReader)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ToySequencerInbox{ToySequencerInboxCaller: ToySequencerInboxCaller{contract: contract}, ToySequencerInboxTransactor: ToySequencerInboxTransactor{contract: contract}, ToySequencerInboxFilterer: ToySequencerInboxFilterer{contract: contract}}, nil
}

// ToySequencerInbox is an auto generated Go binding around an Ethereum contract.
type ToySequencerInbox struct {
	ToySequencerInboxCaller     // Read-only binding to the contract
	ToySequencerInboxTransactor // Write-only binding to the contract
	ToySequencerInboxFilterer   // Log filterer for contract events
}

// ToySequencerInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type ToySequencerInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToySequencerInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ToySequencerInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToySequencerInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ToySequencerInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToySequencerInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ToySequencerInboxSession struct {
	Contract     *ToySequencerInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ToySequencerInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ToySequencerInboxCallerSession struct {
	Contract *ToySequencerInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ToySequencerInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ToySequencerInboxTransactorSession struct {
	Contract     *ToySequencerInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ToySequencerInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type ToySequencerInboxRaw struct {
	Contract *ToySequencerInbox // Generic contract binding to access the raw methods on
}

// ToySequencerInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ToySequencerInboxCallerRaw struct {
	Contract *ToySequencerInboxCaller // Generic read-only contract binding to access the raw methods on
}

// ToySequencerInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ToySequencerInboxTransactorRaw struct {
	Contract *ToySequencerInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToySequencerInbox creates a new instance of ToySequencerInbox, bound to a specific deployed contract.
func NewToySequencerInbox(address common.Address, backend bind.ContractBackend) (*ToySequencerInbox, error) {
	contract, err := bindToySequencerInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ToySequencerInbox{ToySequencerInboxCaller: ToySequencerInboxCaller{contract: contract}, ToySequencerInboxTransactor: ToySequencerInboxTransactor{contract: contract}, ToySequencerInboxFilterer: ToySequencerInboxFilterer{contract: contract}}, nil
}

// NewToySequencerInboxCaller creates a new read-only instance of ToySequencerInbox, bound to a specific deployed contract.
func NewToySequencerInboxCaller(address common.Address, caller bind.ContractCaller) (*ToySequencerInboxCaller, error) {
	contract, err := bindToySequencerInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ToySequencerInboxCaller{contract: contract}, nil
}

// NewToySequencerInboxTransactor creates a new write-only instance of ToySequencerInbox, bound to a specific deployed contract.
func NewToySequencerInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*ToySequencerInboxTransactor, error) {
	contract, err := bindToySequencerInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ToySequencerInboxTransactor{contract: contract}, nil
}

// NewToySequencerInboxFilterer creates a new log filterer instance of ToySequencerInbox, bound to a specific deployed contract.
func NewToySequencerInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*ToySequencerInboxFilterer, error) {
	contract, err := bindToySequencerInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ToySequencerInboxFilterer{contract: contract}, nil
}

// bindToySequencerInbox binds a generic wrapper to an already deployed contract.
func bindToySequencerInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ToySequencerInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ToySequencerInbox *ToySequencerInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ToySequencerInbox.Contract.ToySequencerInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ToySequencerInbox *ToySequencerInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ToySequencerInbox.Contract.ToySequencerInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ToySequencerInbox *ToySequencerInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ToySequencerInbox.Contract.ToySequencerInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ToySequencerInbox *ToySequencerInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ToySequencerInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ToySequencerInbox *ToySequencerInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ToySequencerInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ToySequencerInbox *ToySequencerInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ToySequencerInbox.Contract.contract.Transact(opts, method, params...)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x46115383.
//
// Solidity: function submitBatch() returns()
func (_ToySequencerInbox *ToySequencerInboxTransactor) SubmitBatch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ToySequencerInbox.contract.Transact(opts, "submitBatch")
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x46115383.
//
// Solidity: function submitBatch() returns()
func (_ToySequencerInbox *ToySequencerInboxSession) SubmitBatch() (*types.Transaction, error) {
	return _ToySequencerInbox.Contract.SubmitBatch(&_ToySequencerInbox.TransactOpts)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x46115383.
//
// Solidity: function submitBatch() returns()
func (_ToySequencerInbox *ToySequencerInboxTransactorSession) SubmitBatch() (*types.Transaction, error) {
	return _ToySequencerInbox.Contract.SubmitBatch(&_ToySequencerInbox.TransactOpts)
}

// ToySequencerInboxDataHashesParsedIterator is returned from FilterDataHashesParsed and is used to iterate over the raw logs and unpacked data for DataHashesParsed events raised by the ToySequencerInbox contract.
type ToySequencerInboxDataHashesParsedIterator struct {
	Event *ToySequencerInboxDataHashesParsed // Event containing the contract specifics and raw log

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
func (it *ToySequencerInboxDataHashesParsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ToySequencerInboxDataHashesParsed)
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
		it.Event = new(ToySequencerInboxDataHashesParsed)
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
func (it *ToySequencerInboxDataHashesParsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ToySequencerInboxDataHashesParsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ToySequencerInboxDataHashesParsed represents a DataHashesParsed event raised by the ToySequencerInbox contract.
type ToySequencerInboxDataHashesParsed struct {
	NumHashes *big.Int
	Hashes    [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDataHashesParsed is a free log retrieval operation binding the contract event 0xca7600c3e9508b8063c33ff24d5b042ead3f1d362a4f92868a3d5f8af14cd86b.
//
// Solidity: event DataHashesParsed(uint256 numHashes, bytes32[] hashes)
func (_ToySequencerInbox *ToySequencerInboxFilterer) FilterDataHashesParsed(opts *bind.FilterOpts) (*ToySequencerInboxDataHashesParsedIterator, error) {

	logs, sub, err := _ToySequencerInbox.contract.FilterLogs(opts, "DataHashesParsed")
	if err != nil {
		return nil, err
	}
	return &ToySequencerInboxDataHashesParsedIterator{contract: _ToySequencerInbox.contract, event: "DataHashesParsed", logs: logs, sub: sub}, nil
}

// WatchDataHashesParsed is a free log subscription operation binding the contract event 0xca7600c3e9508b8063c33ff24d5b042ead3f1d362a4f92868a3d5f8af14cd86b.
//
// Solidity: event DataHashesParsed(uint256 numHashes, bytes32[] hashes)
func (_ToySequencerInbox *ToySequencerInboxFilterer) WatchDataHashesParsed(opts *bind.WatchOpts, sink chan<- *ToySequencerInboxDataHashesParsed) (event.Subscription, error) {

	logs, sub, err := _ToySequencerInbox.contract.WatchLogs(opts, "DataHashesParsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ToySequencerInboxDataHashesParsed)
				if err := _ToySequencerInbox.contract.UnpackLog(event, "DataHashesParsed", log); err != nil {
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

// ParseDataHashesParsed is a log parse operation binding the contract event 0xca7600c3e9508b8063c33ff24d5b042ead3f1d362a4f92868a3d5f8af14cd86b.
//
// Solidity: event DataHashesParsed(uint256 numHashes, bytes32[] hashes)
func (_ToySequencerInbox *ToySequencerInboxFilterer) ParseDataHashesParsed(log types.Log) (*ToySequencerInboxDataHashesParsed, error) {
	event := new(ToySequencerInboxDataHashesParsed)
	if err := _ToySequencerInbox.contract.UnpackLog(event, "DataHashesParsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
