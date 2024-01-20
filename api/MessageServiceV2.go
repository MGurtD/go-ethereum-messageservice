// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// MessageServiceMessage is an auto generated low-level Go binding around an user-defined struct.
type MessageServiceMessage struct {
	Addr     common.Address
	Username string
	Text     string
	Read     bool
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getNumberOfMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"readByPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"read\",\"type\":\"bool\"}],\"internalType\":\"structMessageService.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveUnread\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"read\",\"type\":\"bool\"}],\"internalType\":\"structMessageService.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"write\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610bfe8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630c1be226146100645780631e7c01971461007957806341f654f71461008c57806345e63f221461009457806395607ced146100aa578063cf5b31f7146100bf575b5f80fd5b610077610072366004610852565b6100c7565b005b610077610087366004610906565b6101ae565b6100776102e4565b6001546040519081526020015b60405180910390f35b6100b261035f565b6040516100a191906109a9565b6100b2610590565b5f548111156101285760405162461bcd60e51b815260206004820152602260248201527f496e6469636174656420706f736974696f6e20646f6573206e6f742065786973604482015261747360f01b60648201526084015b60405180910390fd5b5f5b5f548110156101aa57818114801561016757505f818154811061014f5761014f610a52565b5f91825260209091206003600490920201015460ff16155b156101a25760015f828154811061018057610180610a52565b5f9182526020909120600490910201600301805460ff19169115159190911790555b60010161012a565b5050565b61012c815111156102015760405162461bcd60e51b815260206004820152601f60248201527f4d657373616765206c696d697420697320333030206368617261637465727300604482015260640161011f565b60408051608081018252338152602081018481529181018390525f60608201819052805460018101825590805281517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563600490920291820180546001600160a01b0319166001600160a01b0390921691909117815592519192917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e564909101906102aa9082610ae4565b50604082015160028201906102bf9082610ae4565b50606091909101516003909101805460ff191691151591909117905550505f54600155565b5f5b5f5481101561035c575f818154811061030157610301610a52565b5f918252602082206003600490920201015460ff16151590036103545760015f828154811061033257610332610a52565b5f9182526020909120600490910201600301805460ff19169115159190911790555b6001016102e6565b50565b5f80546060919067ffffffffffffffff81111561037e5761037e610869565b6040519080825280602002602001820160405280156103e357816020015b6103d060405180608001604052805f6001600160a01b0316815260200160608152602001606081526020015f151581525090565b81526020019060019003908161039c5790505b5090505f5b5f5481101561058a575f818154811061040357610403610a52565b5f918252602091829020604080516080810190915260049092020180546001600160a01b03168252600181018054929391929184019161044290610a66565b80601f016020809104026020016040519081016040528092919081815260200182805461046e90610a66565b80156104b95780601f10610490576101008083540402835291602001916104b9565b820191905f5260205f20905b81548152906001019060200180831161049c57829003601f168201915b505050505081526020016002820180546104d290610a66565b80601f01602080910402602001604051908101604052809291908181526020018280546104fe90610a66565b80156105495780601f1061052057610100808354040283529160200191610549565b820191905f5260205f20905b81548152906001019060200180831161052c57829003601f168201915b50505091835250506003919091015460ff161515602090910152825183908390811061057757610577610a52565b60209081029190910101526001016103e8565b50919050565b60605f805b5f548110156105e3575f81815481106105b0576105b0610a52565b5f918252602082206003600490920201015460ff16151590036105db57816105d781610ba4565b9250505b600101610595565b505f8167ffffffffffffffff8111156105fe576105fe610869565b60405190808252806020026020018201604052801561066357816020015b61065060405180608001604052805f6001600160a01b0316815260200160608152602001606081526020015f151581525090565b81526020019060019003908161061c5790505b5090505f805b5f54811015610849575f818154811061068457610684610a52565b5f918252602082206003600490920201015460ff1615159003610841575f81815481106106b3576106b3610a52565b5f918252602091829020604080516080810190915260049092020180546001600160a01b0316825260018101805492939192918401916106f290610a66565b80601f016020809104026020016040519081016040528092919081815260200182805461071e90610a66565b80156107695780601f1061074057610100808354040283529160200191610769565b820191905f5260205f20905b81548152906001019060200180831161074c57829003601f168201915b5050505050815260200160028201805461078290610a66565b80601f01602080910402602001604051908101604052809291908181526020018280546107ae90610a66565b80156107f95780601f106107d0576101008083540402835291602001916107f9565b820191905f5260205f20905b8154815290600101906020018083116107dc57829003601f168201915b50505091835250506003919091015460ff161515602090910152835184908490811061082757610827610a52565b6020026020010181905250818061083d90610ba4565b9250505b600101610669565b50909392505050565b5f60208284031215610862575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261088c575f80fd5b813567ffffffffffffffff808211156108a7576108a7610869565b604051601f8301601f19908116603f011681019082821181831017156108cf576108cf610869565b816040528381528660208588010111156108e7575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f8060408385031215610917575f80fd5b823567ffffffffffffffff8082111561092e575f80fd5b61093a8683870161087d565b9350602085013591508082111561094f575f80fd5b5061095c8582860161087d565b9150509250929050565b5f81518084525f5b8181101561098a5760208185018101518683018201520161096e565b505f602082860101526020601f19601f83011685010191505092915050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b83811015610a4457888303603f19018552815180516001600160a01b0316845287810151608089860181905290610a0982870182610966565b9150508782015185820389870152610a218282610966565b6060938401511515969093019590955250948701949250908601906001016109d0565b509098975050505050505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680610a7a57607f821691505b60208210810361058a57634e487b7160e01b5f52602260045260245ffd5b601f821115610adf57805f5260205f20601f840160051c81016020851015610abd5750805b601f840160051c820191505b81811015610adc575f8155600101610ac9565b50505b505050565b815167ffffffffffffffff811115610afe57610afe610869565b610b1281610b0c8454610a66565b84610a98565b602080601f831160018114610b45575f8415610b2e5750858301515b5f19600386901b1c1916600185901b178555610b9c565b5f85815260208120601f198616915b82811015610b7357888601518255948401946001909101908401610b54565b5085821015610b9057878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60018201610bc157634e487b7160e01b5f52601160045260245ffd5b506001019056fea26469706673582212207b4060ae8a08a4b3f7a45493773d8d1bcd8cf4966a019636d7ba78add1d846d064736f6c63430008170033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetNumberOfMessages is a free data retrieval call binding the contract method 0x45e63f22.
//
// Solidity: function getNumberOfMessages() view returns(uint256)
func (_Api *ApiCaller) GetNumberOfMessages(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getNumberOfMessages")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfMessages is a free data retrieval call binding the contract method 0x45e63f22.
//
// Solidity: function getNumberOfMessages() view returns(uint256)
func (_Api *ApiSession) GetNumberOfMessages() (*big.Int, error) {
	return _Api.Contract.GetNumberOfMessages(&_Api.CallOpts)
}

// GetNumberOfMessages is a free data retrieval call binding the contract method 0x45e63f22.
//
// Solidity: function getNumberOfMessages() view returns(uint256)
func (_Api *ApiCallerSession) GetNumberOfMessages() (*big.Int, error) {
	return _Api.Contract.GetNumberOfMessages(&_Api.CallOpts)
}

// RetrieveAll is a free data retrieval call binding the contract method 0x95607ced.
//
// Solidity: function retrieveAll() view returns((address,string,string,bool)[])
func (_Api *ApiCaller) RetrieveAll(opts *bind.CallOpts) ([]MessageServiceMessage, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "retrieveAll")

	if err != nil {
		return *new([]MessageServiceMessage), err
	}

	out0 := *abi.ConvertType(out[0], new([]MessageServiceMessage)).(*[]MessageServiceMessage)

	return out0, err

}

// RetrieveAll is a free data retrieval call binding the contract method 0x95607ced.
//
// Solidity: function retrieveAll() view returns((address,string,string,bool)[])
func (_Api *ApiSession) RetrieveAll() ([]MessageServiceMessage, error) {
	return _Api.Contract.RetrieveAll(&_Api.CallOpts)
}

// RetrieveAll is a free data retrieval call binding the contract method 0x95607ced.
//
// Solidity: function retrieveAll() view returns((address,string,string,bool)[])
func (_Api *ApiCallerSession) RetrieveAll() ([]MessageServiceMessage, error) {
	return _Api.Contract.RetrieveAll(&_Api.CallOpts)
}

// RetrieveUnread is a free data retrieval call binding the contract method 0xcf5b31f7.
//
// Solidity: function retrieveUnread() view returns((address,string,string,bool)[])
func (_Api *ApiCaller) RetrieveUnread(opts *bind.CallOpts) ([]MessageServiceMessage, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "retrieveUnread")

	if err != nil {
		return *new([]MessageServiceMessage), err
	}

	out0 := *abi.ConvertType(out[0], new([]MessageServiceMessage)).(*[]MessageServiceMessage)

	return out0, err

}

// RetrieveUnread is a free data retrieval call binding the contract method 0xcf5b31f7.
//
// Solidity: function retrieveUnread() view returns((address,string,string,bool)[])
func (_Api *ApiSession) RetrieveUnread() ([]MessageServiceMessage, error) {
	return _Api.Contract.RetrieveUnread(&_Api.CallOpts)
}

// RetrieveUnread is a free data retrieval call binding the contract method 0xcf5b31f7.
//
// Solidity: function retrieveUnread() view returns((address,string,string,bool)[])
func (_Api *ApiCallerSession) RetrieveUnread() ([]MessageServiceMessage, error) {
	return _Api.Contract.RetrieveUnread(&_Api.CallOpts)
}

// ReadAll is a paid mutator transaction binding the contract method 0x41f654f7.
//
// Solidity: function readAll() returns()
func (_Api *ApiTransactor) ReadAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "readAll")
}

// ReadAll is a paid mutator transaction binding the contract method 0x41f654f7.
//
// Solidity: function readAll() returns()
func (_Api *ApiSession) ReadAll() (*types.Transaction, error) {
	return _Api.Contract.ReadAll(&_Api.TransactOpts)
}

// ReadAll is a paid mutator transaction binding the contract method 0x41f654f7.
//
// Solidity: function readAll() returns()
func (_Api *ApiTransactorSession) ReadAll() (*types.Transaction, error) {
	return _Api.Contract.ReadAll(&_Api.TransactOpts)
}

// ReadByPosition is a paid mutator transaction binding the contract method 0x0c1be226.
//
// Solidity: function readByPosition(uint256 position) returns()
func (_Api *ApiTransactor) ReadByPosition(opts *bind.TransactOpts, position *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "readByPosition", position)
}

// ReadByPosition is a paid mutator transaction binding the contract method 0x0c1be226.
//
// Solidity: function readByPosition(uint256 position) returns()
func (_Api *ApiSession) ReadByPosition(position *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ReadByPosition(&_Api.TransactOpts, position)
}

// ReadByPosition is a paid mutator transaction binding the contract method 0x0c1be226.
//
// Solidity: function readByPosition(uint256 position) returns()
func (_Api *ApiTransactorSession) ReadByPosition(position *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ReadByPosition(&_Api.TransactOpts, position)
}

// Write is a paid mutator transaction binding the contract method 0x1e7c0197.
//
// Solidity: function write(string username, string message) returns()
func (_Api *ApiTransactor) Write(opts *bind.TransactOpts, username string, message string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "write", username, message)
}

// Write is a paid mutator transaction binding the contract method 0x1e7c0197.
//
// Solidity: function write(string username, string message) returns()
func (_Api *ApiSession) Write(username string, message string) (*types.Transaction, error) {
	return _Api.Contract.Write(&_Api.TransactOpts, username, message)
}

// Write is a paid mutator transaction binding the contract method 0x1e7c0197.
//
// Solidity: function write(string username, string message) returns()
func (_Api *ApiTransactorSession) Write(username string, message string) (*types.Transaction, error) {
	return _Api.Contract.Write(&_Api.TransactOpts, username, message)
}
