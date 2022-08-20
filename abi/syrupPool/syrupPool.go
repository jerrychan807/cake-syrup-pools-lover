// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package syrupPool

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

// SyrupPoolMetaData contains all meta data concerning the SyrupPool contract.
var SyrupPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pancakeProfile\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_pancakeProfileIsRequested\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_pancakeProfileThresholdPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"NewPoolLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"NewRewardPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"NewStartAndEndBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RewardsStop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isProfileRequested\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"thresholdPoints\",\"type\":\"uint256\"}],\"name\":\"UpdateProfileAndThresholdPointsRequirement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SMART_CHEF_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accTokenPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyRewardWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasUserLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numberBlocksForUserLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberBlocksForUserLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeProfile\",\"outputs\":[{\"internalType\":\"contractIPancakeProfile\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeProfileIsRequested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeProfileThresholdPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLimitPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_userLimit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"updatePoolLimitPerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isRequested\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_thresholdPoints\",\"type\":\"uint256\"}],\"name\":\"updateProfileAndThresholdPointsRequirement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"updateRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"name\":\"updateStartAndEndBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SyrupPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SyrupPoolMetaData.ABI instead.
var SyrupPoolABI = SyrupPoolMetaData.ABI

// SyrupPool is an auto generated Go binding around an Ethereum contract.
type SyrupPool struct {
	SyrupPoolCaller     // Read-only binding to the contract
	SyrupPoolTransactor // Write-only binding to the contract
	SyrupPoolFilterer   // Log filterer for contract events
}

// SyrupPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SyrupPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyrupPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SyrupPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyrupPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SyrupPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyrupPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SyrupPoolSession struct {
	Contract     *SyrupPool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SyrupPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SyrupPoolCallerSession struct {
	Contract *SyrupPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SyrupPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SyrupPoolTransactorSession struct {
	Contract     *SyrupPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SyrupPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SyrupPoolRaw struct {
	Contract *SyrupPool // Generic contract binding to access the raw methods on
}

// SyrupPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SyrupPoolCallerRaw struct {
	Contract *SyrupPoolCaller // Generic read-only contract binding to access the raw methods on
}

// SyrupPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SyrupPoolTransactorRaw struct {
	Contract *SyrupPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSyrupPool creates a new instance of SyrupPool, bound to a specific deployed contract.
func NewSyrupPool(address common.Address, backend bind.ContractBackend) (*SyrupPool, error) {
	contract, err := bindSyrupPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SyrupPool{SyrupPoolCaller: SyrupPoolCaller{contract: contract}, SyrupPoolTransactor: SyrupPoolTransactor{contract: contract}, SyrupPoolFilterer: SyrupPoolFilterer{contract: contract}}, nil
}

// NewSyrupPoolCaller creates a new read-only instance of SyrupPool, bound to a specific deployed contract.
func NewSyrupPoolCaller(address common.Address, caller bind.ContractCaller) (*SyrupPoolCaller, error) {
	contract, err := bindSyrupPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SyrupPoolCaller{contract: contract}, nil
}

// NewSyrupPoolTransactor creates a new write-only instance of SyrupPool, bound to a specific deployed contract.
func NewSyrupPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SyrupPoolTransactor, error) {
	contract, err := bindSyrupPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SyrupPoolTransactor{contract: contract}, nil
}

// NewSyrupPoolFilterer creates a new log filterer instance of SyrupPool, bound to a specific deployed contract.
func NewSyrupPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SyrupPoolFilterer, error) {
	contract, err := bindSyrupPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SyrupPoolFilterer{contract: contract}, nil
}

// bindSyrupPool binds a generic wrapper to an already deployed contract.
func bindSyrupPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SyrupPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyrupPool *SyrupPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyrupPool.Contract.SyrupPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyrupPool *SyrupPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyrupPool.Contract.SyrupPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyrupPool *SyrupPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyrupPool.Contract.SyrupPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyrupPool *SyrupPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyrupPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyrupPool *SyrupPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyrupPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyrupPool *SyrupPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyrupPool.Contract.contract.Transact(opts, method, params...)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SyrupPool *SyrupPoolSession) PRECISIONFACTOR() (*big.Int, error) {
	return _SyrupPool.Contract.PRECISIONFACTOR(&_SyrupPool.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _SyrupPool.Contract.PRECISIONFACTOR(&_SyrupPool.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_SyrupPool *SyrupPoolCaller) SMARTCHEFFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "SMART_CHEF_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_SyrupPool *SyrupPoolSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _SyrupPool.Contract.SMARTCHEFFACTORY(&_SyrupPool.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_SyrupPool *SyrupPoolCallerSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _SyrupPool.Contract.SMARTCHEFFACTORY(&_SyrupPool.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) AccTokenPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "accTokenPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_SyrupPool *SyrupPoolSession) AccTokenPerShare() (*big.Int, error) {
	return _SyrupPool.Contract.AccTokenPerShare(&_SyrupPool.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) AccTokenPerShare() (*big.Int, error) {
	return _SyrupPool.Contract.AccTokenPerShare(&_SyrupPool.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolSession) BonusEndBlock() (*big.Int, error) {
	return _SyrupPool.Contract.BonusEndBlock(&_SyrupPool.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) BonusEndBlock() (*big.Int, error) {
	return _SyrupPool.Contract.BonusEndBlock(&_SyrupPool.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_SyrupPool *SyrupPoolCaller) HasUserLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "hasUserLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_SyrupPool *SyrupPoolSession) HasUserLimit() (bool, error) {
	return _SyrupPool.Contract.HasUserLimit(&_SyrupPool.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_SyrupPool *SyrupPoolCallerSession) HasUserLimit() (bool, error) {
	return _SyrupPool.Contract.HasUserLimit(&_SyrupPool.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SyrupPool *SyrupPoolCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SyrupPool *SyrupPoolSession) IsInitialized() (bool, error) {
	return _SyrupPool.Contract.IsInitialized(&_SyrupPool.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SyrupPool *SyrupPoolCallerSession) IsInitialized() (bool, error) {
	return _SyrupPool.Contract.IsInitialized(&_SyrupPool.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolSession) LastRewardBlock() (*big.Int, error) {
	return _SyrupPool.Contract.LastRewardBlock(&_SyrupPool.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) LastRewardBlock() (*big.Int, error) {
	return _SyrupPool.Contract.LastRewardBlock(&_SyrupPool.CallOpts)
}

// NumberBlocksForUserLimit is a free data retrieval call binding the contract method 0x8ad1071b.
//
// Solidity: function numberBlocksForUserLimit() view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) NumberBlocksForUserLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "numberBlocksForUserLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberBlocksForUserLimit is a free data retrieval call binding the contract method 0x8ad1071b.
//
// Solidity: function numberBlocksForUserLimit() view returns(uint256)
func (_SyrupPool *SyrupPoolSession) NumberBlocksForUserLimit() (*big.Int, error) {
	return _SyrupPool.Contract.NumberBlocksForUserLimit(&_SyrupPool.CallOpts)
}

// NumberBlocksForUserLimit is a free data retrieval call binding the contract method 0x8ad1071b.
//
// Solidity: function numberBlocksForUserLimit() view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) NumberBlocksForUserLimit() (*big.Int, error) {
	return _SyrupPool.Contract.NumberBlocksForUserLimit(&_SyrupPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyrupPool *SyrupPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyrupPool *SyrupPoolSession) Owner() (common.Address, error) {
	return _SyrupPool.Contract.Owner(&_SyrupPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyrupPool *SyrupPoolCallerSession) Owner() (common.Address, error) {
	return _SyrupPool.Contract.Owner(&_SyrupPool.CallOpts)
}

// PancakeProfile is a free data retrieval call binding the contract method 0xc7d936ec.
//
// Solidity: function pancakeProfile() view returns(address)
func (_SyrupPool *SyrupPoolCaller) PancakeProfile(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "pancakeProfile")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeProfile is a free data retrieval call binding the contract method 0xc7d936ec.
//
// Solidity: function pancakeProfile() view returns(address)
func (_SyrupPool *SyrupPoolSession) PancakeProfile() (common.Address, error) {
	return _SyrupPool.Contract.PancakeProfile(&_SyrupPool.CallOpts)
}

// PancakeProfile is a free data retrieval call binding the contract method 0xc7d936ec.
//
// Solidity: function pancakeProfile() view returns(address)
func (_SyrupPool *SyrupPoolCallerSession) PancakeProfile() (common.Address, error) {
	return _SyrupPool.Contract.PancakeProfile(&_SyrupPool.CallOpts)
}

// PancakeProfileIsRequested is a free data retrieval call binding the contract method 0x5a0b5f34.
//
// Solidity: function pancakeProfileIsRequested() view returns(bool)
func (_SyrupPool *SyrupPoolCaller) PancakeProfileIsRequested(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "pancakeProfileIsRequested")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PancakeProfileIsRequested is a free data retrieval call binding the contract method 0x5a0b5f34.
//
// Solidity: function pancakeProfileIsRequested() view returns(bool)
func (_SyrupPool *SyrupPoolSession) PancakeProfileIsRequested() (bool, error) {
	return _SyrupPool.Contract.PancakeProfileIsRequested(&_SyrupPool.CallOpts)
}

// PancakeProfileIsRequested is a free data retrieval call binding the contract method 0x5a0b5f34.
//
// Solidity: function pancakeProfileIsRequested() view returns(bool)
func (_SyrupPool *SyrupPoolCallerSession) PancakeProfileIsRequested() (bool, error) {
	return _SyrupPool.Contract.PancakeProfileIsRequested(&_SyrupPool.CallOpts)
}

// PancakeProfileThresholdPoints is a free data retrieval call binding the contract method 0x0ace6247.
//
// Solidity: function pancakeProfileThresholdPoints() view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) PancakeProfileThresholdPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "pancakeProfileThresholdPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PancakeProfileThresholdPoints is a free data retrieval call binding the contract method 0x0ace6247.
//
// Solidity: function pancakeProfileThresholdPoints() view returns(uint256)
func (_SyrupPool *SyrupPoolSession) PancakeProfileThresholdPoints() (*big.Int, error) {
	return _SyrupPool.Contract.PancakeProfileThresholdPoints(&_SyrupPool.CallOpts)
}

// PancakeProfileThresholdPoints is a free data retrieval call binding the contract method 0x0ace6247.
//
// Solidity: function pancakeProfileThresholdPoints() view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) PancakeProfileThresholdPoints() (*big.Int, error) {
	return _SyrupPool.Contract.PancakeProfileThresholdPoints(&_SyrupPool.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) PendingReward(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "pendingReward", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_SyrupPool *SyrupPoolSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _SyrupPool.Contract.PendingReward(&_SyrupPool.CallOpts, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _SyrupPool.Contract.PendingReward(&_SyrupPool.CallOpts, _user)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) PoolLimitPerUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "poolLimitPerUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_SyrupPool *SyrupPoolSession) PoolLimitPerUser() (*big.Int, error) {
	return _SyrupPool.Contract.PoolLimitPerUser(&_SyrupPool.CallOpts)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) PoolLimitPerUser() (*big.Int, error) {
	return _SyrupPool.Contract.PoolLimitPerUser(&_SyrupPool.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) RewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "rewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolSession) RewardPerBlock() (*big.Int, error) {
	return _SyrupPool.Contract.RewardPerBlock(&_SyrupPool.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) RewardPerBlock() (*big.Int, error) {
	return _SyrupPool.Contract.RewardPerBlock(&_SyrupPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SyrupPool *SyrupPoolCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SyrupPool *SyrupPoolSession) RewardToken() (common.Address, error) {
	return _SyrupPool.Contract.RewardToken(&_SyrupPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SyrupPool *SyrupPoolCallerSession) RewardToken() (common.Address, error) {
	return _SyrupPool.Contract.RewardToken(&_SyrupPool.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_SyrupPool *SyrupPoolCaller) StakedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "stakedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_SyrupPool *SyrupPoolSession) StakedToken() (common.Address, error) {
	return _SyrupPool.Contract.StakedToken(&_SyrupPool.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_SyrupPool *SyrupPoolCallerSession) StakedToken() (common.Address, error) {
	return _SyrupPool.Contract.StakedToken(&_SyrupPool.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolSession) StartBlock() (*big.Int, error) {
	return _SyrupPool.Contract.StartBlock(&_SyrupPool.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SyrupPool *SyrupPoolCallerSession) StartBlock() (*big.Int, error) {
	return _SyrupPool.Contract.StartBlock(&_SyrupPool.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_SyrupPool *SyrupPoolCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_SyrupPool *SyrupPoolSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _SyrupPool.Contract.UserInfo(&_SyrupPool.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_SyrupPool *SyrupPoolCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _SyrupPool.Contract.UserInfo(&_SyrupPool.CallOpts, arg0)
}

// UserLimit is a free data retrieval call binding the contract method 0x4a7c01ec.
//
// Solidity: function userLimit() view returns(bool)
func (_SyrupPool *SyrupPoolCaller) UserLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SyrupPool.contract.Call(opts, &out, "userLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserLimit is a free data retrieval call binding the contract method 0x4a7c01ec.
//
// Solidity: function userLimit() view returns(bool)
func (_SyrupPool *SyrupPoolSession) UserLimit() (bool, error) {
	return _SyrupPool.Contract.UserLimit(&_SyrupPool.CallOpts)
}

// UserLimit is a free data retrieval call binding the contract method 0x4a7c01ec.
//
// Solidity: function userLimit() view returns(bool)
func (_SyrupPool *SyrupPoolCallerSession) UserLimit() (bool, error) {
	return _SyrupPool.Contract.UserLimit(&_SyrupPool.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_SyrupPool *SyrupPoolTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_SyrupPool *SyrupPoolSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.Deposit(&_SyrupPool.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_SyrupPool *SyrupPoolTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.Deposit(&_SyrupPool.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_SyrupPool *SyrupPoolTransactor) EmergencyRewardWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "emergencyRewardWithdraw", _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_SyrupPool *SyrupPoolSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.EmergencyRewardWithdraw(&_SyrupPool.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_SyrupPool *SyrupPoolTransactorSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.EmergencyRewardWithdraw(&_SyrupPool.TransactOpts, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_SyrupPool *SyrupPoolTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_SyrupPool *SyrupPoolSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _SyrupPool.Contract.EmergencyWithdraw(&_SyrupPool.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_SyrupPool *SyrupPoolTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _SyrupPool.Contract.EmergencyWithdraw(&_SyrupPool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x2aa2c381.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, uint256 _numberBlocksForUserLimit, address _admin) returns()
func (_SyrupPool *SyrupPoolTransactor) Initialize(opts *bind.TransactOpts, _stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _numberBlocksForUserLimit *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "initialize", _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _numberBlocksForUserLimit, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x2aa2c381.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, uint256 _numberBlocksForUserLimit, address _admin) returns()
func (_SyrupPool *SyrupPoolSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _numberBlocksForUserLimit *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _SyrupPool.Contract.Initialize(&_SyrupPool.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _numberBlocksForUserLimit, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x2aa2c381.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, uint256 _numberBlocksForUserLimit, address _admin) returns()
func (_SyrupPool *SyrupPoolTransactorSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _numberBlocksForUserLimit *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _SyrupPool.Contract.Initialize(&_SyrupPool.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _numberBlocksForUserLimit, _admin)
}

// RecoverToken is a paid mutator transaction binding the contract method 0x9be65a60.
//
// Solidity: function recoverToken(address _token) returns()
func (_SyrupPool *SyrupPoolTransactor) RecoverToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "recoverToken", _token)
}

// RecoverToken is a paid mutator transaction binding the contract method 0x9be65a60.
//
// Solidity: function recoverToken(address _token) returns()
func (_SyrupPool *SyrupPoolSession) RecoverToken(_token common.Address) (*types.Transaction, error) {
	return _SyrupPool.Contract.RecoverToken(&_SyrupPool.TransactOpts, _token)
}

// RecoverToken is a paid mutator transaction binding the contract method 0x9be65a60.
//
// Solidity: function recoverToken(address _token) returns()
func (_SyrupPool *SyrupPoolTransactorSession) RecoverToken(_token common.Address) (*types.Transaction, error) {
	return _SyrupPool.Contract.RecoverToken(&_SyrupPool.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyrupPool *SyrupPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyrupPool *SyrupPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _SyrupPool.Contract.RenounceOwnership(&_SyrupPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyrupPool *SyrupPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SyrupPool.Contract.RenounceOwnership(&_SyrupPool.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_SyrupPool *SyrupPoolTransactor) StopReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "stopReward")
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_SyrupPool *SyrupPoolSession) StopReward() (*types.Transaction, error) {
	return _SyrupPool.Contract.StopReward(&_SyrupPool.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_SyrupPool *SyrupPoolTransactorSession) StopReward() (*types.Transaction, error) {
	return _SyrupPool.Contract.StopReward(&_SyrupPool.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyrupPool *SyrupPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyrupPool *SyrupPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SyrupPool.Contract.TransferOwnership(&_SyrupPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyrupPool *SyrupPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SyrupPool.Contract.TransferOwnership(&_SyrupPool.TransactOpts, newOwner)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _userLimit, uint256 _poolLimitPerUser) returns()
func (_SyrupPool *SyrupPoolTransactor) UpdatePoolLimitPerUser(opts *bind.TransactOpts, _userLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "updatePoolLimitPerUser", _userLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _userLimit, uint256 _poolLimitPerUser) returns()
func (_SyrupPool *SyrupPoolSession) UpdatePoolLimitPerUser(_userLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.UpdatePoolLimitPerUser(&_SyrupPool.TransactOpts, _userLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _userLimit, uint256 _poolLimitPerUser) returns()
func (_SyrupPool *SyrupPoolTransactorSession) UpdatePoolLimitPerUser(_userLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.UpdatePoolLimitPerUser(&_SyrupPool.TransactOpts, _userLimit, _poolLimitPerUser)
}

// UpdateProfileAndThresholdPointsRequirement is a paid mutator transaction binding the contract method 0x68109631.
//
// Solidity: function updateProfileAndThresholdPointsRequirement(bool _isRequested, uint256 _thresholdPoints) returns()
func (_SyrupPool *SyrupPoolTransactor) UpdateProfileAndThresholdPointsRequirement(opts *bind.TransactOpts, _isRequested bool, _thresholdPoints *big.Int) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "updateProfileAndThresholdPointsRequirement", _isRequested, _thresholdPoints)
}

// UpdateProfileAndThresholdPointsRequirement is a paid mutator transaction binding the contract method 0x68109631.
//
// Solidity: function updateProfileAndThresholdPointsRequirement(bool _isRequested, uint256 _thresholdPoints) returns()
func (_SyrupPool *SyrupPoolSession) UpdateProfileAndThresholdPointsRequirement(_isRequested bool, _thresholdPoints *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.UpdateProfileAndThresholdPointsRequirement(&_SyrupPool.TransactOpts, _isRequested, _thresholdPoints)
}

// UpdateProfileAndThresholdPointsRequirement is a paid mutator transaction binding the contract method 0x68109631.
//
// Solidity: function updateProfileAndThresholdPointsRequirement(bool _isRequested, uint256 _thresholdPoints) returns()
func (_SyrupPool *SyrupPoolTransactorSession) UpdateProfileAndThresholdPointsRequirement(_isRequested bool, _thresholdPoints *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.UpdateProfileAndThresholdPointsRequirement(&_SyrupPool.TransactOpts, _isRequested, _thresholdPoints)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_SyrupPool *SyrupPoolTransactor) UpdateRewardPerBlock(opts *bind.TransactOpts, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "updateRewardPerBlock", _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_SyrupPool *SyrupPoolSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.UpdateRewardPerBlock(&_SyrupPool.TransactOpts, _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_SyrupPool *SyrupPoolTransactorSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.UpdateRewardPerBlock(&_SyrupPool.TransactOpts, _rewardPerBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_SyrupPool *SyrupPoolTransactor) UpdateStartAndEndBlocks(opts *bind.TransactOpts, _startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "updateStartAndEndBlocks", _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_SyrupPool *SyrupPoolSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.UpdateStartAndEndBlocks(&_SyrupPool.TransactOpts, _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_SyrupPool *SyrupPoolTransactorSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.UpdateStartAndEndBlocks(&_SyrupPool.TransactOpts, _startBlock, _bonusEndBlock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_SyrupPool *SyrupPoolTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SyrupPool.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_SyrupPool *SyrupPoolSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.Withdraw(&_SyrupPool.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_SyrupPool *SyrupPoolTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SyrupPool.Contract.Withdraw(&_SyrupPool.TransactOpts, _amount)
}

// SyrupPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SyrupPool contract.
type SyrupPoolDepositIterator struct {
	Event *SyrupPoolDeposit // Event containing the contract specifics and raw log

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
func (it *SyrupPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolDeposit)
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
		it.Event = new(SyrupPoolDeposit)
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
func (it *SyrupPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolDeposit represents a Deposit event raised by the SyrupPool contract.
type SyrupPoolDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*SyrupPoolDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &SyrupPoolDepositIterator{contract: _SyrupPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SyrupPoolDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolDeposit)
				if err := _SyrupPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) ParseDeposit(log types.Log) (*SyrupPoolDeposit, error) {
	event := new(SyrupPoolDeposit)
	if err := _SyrupPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyrupPoolEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the SyrupPool contract.
type SyrupPoolEmergencyWithdrawIterator struct {
	Event *SyrupPoolEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *SyrupPoolEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolEmergencyWithdraw)
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
		it.Event = new(SyrupPoolEmergencyWithdraw)
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
func (it *SyrupPoolEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolEmergencyWithdraw represents a EmergencyWithdraw event raised by the SyrupPool contract.
type SyrupPoolEmergencyWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address) (*SyrupPoolEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &SyrupPoolEmergencyWithdrawIterator{contract: _SyrupPool.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *SyrupPoolEmergencyWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolEmergencyWithdraw)
				if err := _SyrupPool.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) ParseEmergencyWithdraw(log types.Log) (*SyrupPoolEmergencyWithdraw, error) {
	event := new(SyrupPoolEmergencyWithdraw)
	if err := _SyrupPool.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyrupPoolNewPoolLimitIterator is returned from FilterNewPoolLimit and is used to iterate over the raw logs and unpacked data for NewPoolLimit events raised by the SyrupPool contract.
type SyrupPoolNewPoolLimitIterator struct {
	Event *SyrupPoolNewPoolLimit // Event containing the contract specifics and raw log

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
func (it *SyrupPoolNewPoolLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolNewPoolLimit)
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
		it.Event = new(SyrupPoolNewPoolLimit)
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
func (it *SyrupPoolNewPoolLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolNewPoolLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolNewPoolLimit represents a NewPoolLimit event raised by the SyrupPool contract.
type SyrupPoolNewPoolLimit struct {
	PoolLimitPerUser *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPoolLimit is a free log retrieval operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_SyrupPool *SyrupPoolFilterer) FilterNewPoolLimit(opts *bind.FilterOpts) (*SyrupPoolNewPoolLimitIterator, error) {

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return &SyrupPoolNewPoolLimitIterator{contract: _SyrupPool.contract, event: "NewPoolLimit", logs: logs, sub: sub}, nil
}

// WatchNewPoolLimit is a free log subscription operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_SyrupPool *SyrupPoolFilterer) WatchNewPoolLimit(opts *bind.WatchOpts, sink chan<- *SyrupPoolNewPoolLimit) (event.Subscription, error) {

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolNewPoolLimit)
				if err := _SyrupPool.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
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

// ParseNewPoolLimit is a log parse operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_SyrupPool *SyrupPoolFilterer) ParseNewPoolLimit(log types.Log) (*SyrupPoolNewPoolLimit, error) {
	event := new(SyrupPoolNewPoolLimit)
	if err := _SyrupPool.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyrupPoolNewRewardPerBlockIterator is returned from FilterNewRewardPerBlock and is used to iterate over the raw logs and unpacked data for NewRewardPerBlock events raised by the SyrupPool contract.
type SyrupPoolNewRewardPerBlockIterator struct {
	Event *SyrupPoolNewRewardPerBlock // Event containing the contract specifics and raw log

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
func (it *SyrupPoolNewRewardPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolNewRewardPerBlock)
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
		it.Event = new(SyrupPoolNewRewardPerBlock)
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
func (it *SyrupPoolNewRewardPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolNewRewardPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolNewRewardPerBlock represents a NewRewardPerBlock event raised by the SyrupPool contract.
type SyrupPoolNewRewardPerBlock struct {
	RewardPerBlock *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewRewardPerBlock is a free log retrieval operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_SyrupPool *SyrupPoolFilterer) FilterNewRewardPerBlock(opts *bind.FilterOpts) (*SyrupPoolNewRewardPerBlockIterator, error) {

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return &SyrupPoolNewRewardPerBlockIterator{contract: _SyrupPool.contract, event: "NewRewardPerBlock", logs: logs, sub: sub}, nil
}

// WatchNewRewardPerBlock is a free log subscription operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_SyrupPool *SyrupPoolFilterer) WatchNewRewardPerBlock(opts *bind.WatchOpts, sink chan<- *SyrupPoolNewRewardPerBlock) (event.Subscription, error) {

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolNewRewardPerBlock)
				if err := _SyrupPool.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
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

// ParseNewRewardPerBlock is a log parse operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_SyrupPool *SyrupPoolFilterer) ParseNewRewardPerBlock(log types.Log) (*SyrupPoolNewRewardPerBlock, error) {
	event := new(SyrupPoolNewRewardPerBlock)
	if err := _SyrupPool.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyrupPoolNewStartAndEndBlocksIterator is returned from FilterNewStartAndEndBlocks and is used to iterate over the raw logs and unpacked data for NewStartAndEndBlocks events raised by the SyrupPool contract.
type SyrupPoolNewStartAndEndBlocksIterator struct {
	Event *SyrupPoolNewStartAndEndBlocks // Event containing the contract specifics and raw log

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
func (it *SyrupPoolNewStartAndEndBlocksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolNewStartAndEndBlocks)
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
		it.Event = new(SyrupPoolNewStartAndEndBlocks)
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
func (it *SyrupPoolNewStartAndEndBlocksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolNewStartAndEndBlocksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolNewStartAndEndBlocks represents a NewStartAndEndBlocks event raised by the SyrupPool contract.
type SyrupPoolNewStartAndEndBlocks struct {
	StartBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewStartAndEndBlocks is a free log retrieval operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_SyrupPool *SyrupPoolFilterer) FilterNewStartAndEndBlocks(opts *bind.FilterOpts) (*SyrupPoolNewStartAndEndBlocksIterator, error) {

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return &SyrupPoolNewStartAndEndBlocksIterator{contract: _SyrupPool.contract, event: "NewStartAndEndBlocks", logs: logs, sub: sub}, nil
}

// WatchNewStartAndEndBlocks is a free log subscription operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_SyrupPool *SyrupPoolFilterer) WatchNewStartAndEndBlocks(opts *bind.WatchOpts, sink chan<- *SyrupPoolNewStartAndEndBlocks) (event.Subscription, error) {

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolNewStartAndEndBlocks)
				if err := _SyrupPool.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
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

// ParseNewStartAndEndBlocks is a log parse operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_SyrupPool *SyrupPoolFilterer) ParseNewStartAndEndBlocks(log types.Log) (*SyrupPoolNewStartAndEndBlocks, error) {
	event := new(SyrupPoolNewStartAndEndBlocks)
	if err := _SyrupPool.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyrupPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SyrupPool contract.
type SyrupPoolOwnershipTransferredIterator struct {
	Event *SyrupPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SyrupPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolOwnershipTransferred)
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
		it.Event = new(SyrupPoolOwnershipTransferred)
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
func (it *SyrupPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolOwnershipTransferred represents a OwnershipTransferred event raised by the SyrupPool contract.
type SyrupPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SyrupPool *SyrupPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SyrupPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SyrupPoolOwnershipTransferredIterator{contract: _SyrupPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SyrupPool *SyrupPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SyrupPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolOwnershipTransferred)
				if err := _SyrupPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SyrupPool *SyrupPoolFilterer) ParseOwnershipTransferred(log types.Log) (*SyrupPoolOwnershipTransferred, error) {
	event := new(SyrupPoolOwnershipTransferred)
	if err := _SyrupPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyrupPoolRewardsStopIterator is returned from FilterRewardsStop and is used to iterate over the raw logs and unpacked data for RewardsStop events raised by the SyrupPool contract.
type SyrupPoolRewardsStopIterator struct {
	Event *SyrupPoolRewardsStop // Event containing the contract specifics and raw log

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
func (it *SyrupPoolRewardsStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolRewardsStop)
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
		it.Event = new(SyrupPoolRewardsStop)
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
func (it *SyrupPoolRewardsStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolRewardsStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolRewardsStop represents a RewardsStop event raised by the SyrupPool contract.
type SyrupPoolRewardsStop struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsStop is a free log retrieval operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_SyrupPool *SyrupPoolFilterer) FilterRewardsStop(opts *bind.FilterOpts) (*SyrupPoolRewardsStopIterator, error) {

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return &SyrupPoolRewardsStopIterator{contract: _SyrupPool.contract, event: "RewardsStop", logs: logs, sub: sub}, nil
}

// WatchRewardsStop is a free log subscription operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_SyrupPool *SyrupPoolFilterer) WatchRewardsStop(opts *bind.WatchOpts, sink chan<- *SyrupPoolRewardsStop) (event.Subscription, error) {

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolRewardsStop)
				if err := _SyrupPool.contract.UnpackLog(event, "RewardsStop", log); err != nil {
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

// ParseRewardsStop is a log parse operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_SyrupPool *SyrupPoolFilterer) ParseRewardsStop(log types.Log) (*SyrupPoolRewardsStop, error) {
	event := new(SyrupPoolRewardsStop)
	if err := _SyrupPool.contract.UnpackLog(event, "RewardsStop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyrupPoolTokenRecoveryIterator is returned from FilterTokenRecovery and is used to iterate over the raw logs and unpacked data for TokenRecovery events raised by the SyrupPool contract.
type SyrupPoolTokenRecoveryIterator struct {
	Event *SyrupPoolTokenRecovery // Event containing the contract specifics and raw log

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
func (it *SyrupPoolTokenRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolTokenRecovery)
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
		it.Event = new(SyrupPoolTokenRecovery)
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
func (it *SyrupPoolTokenRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolTokenRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolTokenRecovery represents a TokenRecovery event raised by the SyrupPool contract.
type SyrupPoolTokenRecovery struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRecovery is a free log retrieval operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) FilterTokenRecovery(opts *bind.FilterOpts, token []common.Address) (*SyrupPoolTokenRecoveryIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "TokenRecovery", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SyrupPoolTokenRecoveryIterator{contract: _SyrupPool.contract, event: "TokenRecovery", logs: logs, sub: sub}, nil
}

// WatchTokenRecovery is a free log subscription operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) WatchTokenRecovery(opts *bind.WatchOpts, sink chan<- *SyrupPoolTokenRecovery, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "TokenRecovery", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolTokenRecovery)
				if err := _SyrupPool.contract.UnpackLog(event, "TokenRecovery", log); err != nil {
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

// ParseTokenRecovery is a log parse operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) ParseTokenRecovery(log types.Log) (*SyrupPoolTokenRecovery, error) {
	event := new(SyrupPoolTokenRecovery)
	if err := _SyrupPool.contract.UnpackLog(event, "TokenRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyrupPoolUpdateProfileAndThresholdPointsRequirementIterator is returned from FilterUpdateProfileAndThresholdPointsRequirement and is used to iterate over the raw logs and unpacked data for UpdateProfileAndThresholdPointsRequirement events raised by the SyrupPool contract.
type SyrupPoolUpdateProfileAndThresholdPointsRequirementIterator struct {
	Event *SyrupPoolUpdateProfileAndThresholdPointsRequirement // Event containing the contract specifics and raw log

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
func (it *SyrupPoolUpdateProfileAndThresholdPointsRequirementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolUpdateProfileAndThresholdPointsRequirement)
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
		it.Event = new(SyrupPoolUpdateProfileAndThresholdPointsRequirement)
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
func (it *SyrupPoolUpdateProfileAndThresholdPointsRequirementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolUpdateProfileAndThresholdPointsRequirementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolUpdateProfileAndThresholdPointsRequirement represents a UpdateProfileAndThresholdPointsRequirement event raised by the SyrupPool contract.
type SyrupPoolUpdateProfileAndThresholdPointsRequirement struct {
	IsProfileRequested bool
	ThresholdPoints    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateProfileAndThresholdPointsRequirement is a free log retrieval operation binding the contract event 0x915d08e0e89c58e352d7c1d66c942cb15dac8a7294d2ca80ddf46f1998f0512b.
//
// Solidity: event UpdateProfileAndThresholdPointsRequirement(bool isProfileRequested, uint256 thresholdPoints)
func (_SyrupPool *SyrupPoolFilterer) FilterUpdateProfileAndThresholdPointsRequirement(opts *bind.FilterOpts) (*SyrupPoolUpdateProfileAndThresholdPointsRequirementIterator, error) {

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "UpdateProfileAndThresholdPointsRequirement")
	if err != nil {
		return nil, err
	}
	return &SyrupPoolUpdateProfileAndThresholdPointsRequirementIterator{contract: _SyrupPool.contract, event: "UpdateProfileAndThresholdPointsRequirement", logs: logs, sub: sub}, nil
}

// WatchUpdateProfileAndThresholdPointsRequirement is a free log subscription operation binding the contract event 0x915d08e0e89c58e352d7c1d66c942cb15dac8a7294d2ca80ddf46f1998f0512b.
//
// Solidity: event UpdateProfileAndThresholdPointsRequirement(bool isProfileRequested, uint256 thresholdPoints)
func (_SyrupPool *SyrupPoolFilterer) WatchUpdateProfileAndThresholdPointsRequirement(opts *bind.WatchOpts, sink chan<- *SyrupPoolUpdateProfileAndThresholdPointsRequirement) (event.Subscription, error) {

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "UpdateProfileAndThresholdPointsRequirement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolUpdateProfileAndThresholdPointsRequirement)
				if err := _SyrupPool.contract.UnpackLog(event, "UpdateProfileAndThresholdPointsRequirement", log); err != nil {
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

// ParseUpdateProfileAndThresholdPointsRequirement is a log parse operation binding the contract event 0x915d08e0e89c58e352d7c1d66c942cb15dac8a7294d2ca80ddf46f1998f0512b.
//
// Solidity: event UpdateProfileAndThresholdPointsRequirement(bool isProfileRequested, uint256 thresholdPoints)
func (_SyrupPool *SyrupPoolFilterer) ParseUpdateProfileAndThresholdPointsRequirement(log types.Log) (*SyrupPoolUpdateProfileAndThresholdPointsRequirement, error) {
	event := new(SyrupPoolUpdateProfileAndThresholdPointsRequirement)
	if err := _SyrupPool.contract.UnpackLog(event, "UpdateProfileAndThresholdPointsRequirement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyrupPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SyrupPool contract.
type SyrupPoolWithdrawIterator struct {
	Event *SyrupPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *SyrupPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyrupPoolWithdraw)
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
		it.Event = new(SyrupPoolWithdraw)
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
func (it *SyrupPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyrupPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyrupPoolWithdraw represents a Withdraw event raised by the SyrupPool contract.
type SyrupPoolWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*SyrupPoolWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SyrupPool.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &SyrupPoolWithdrawIterator{contract: _SyrupPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SyrupPoolWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SyrupPool.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyrupPoolWithdraw)
				if err := _SyrupPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_SyrupPool *SyrupPoolFilterer) ParseWithdraw(log types.Log) (*SyrupPoolWithdraw, error) {
	event := new(SyrupPoolWithdraw)
	if err := _SyrupPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
