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
	ABI: "[{\"inputs\":[],\"name\":\"readAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"readByPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"read\",\"type\":\"bool\"}],\"internalType\":\"structMessageService.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveUnread\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"read\",\"type\":\"bool\"}],\"internalType\":\"structMessageService.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"write\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610be18061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80630c1be226146100595780631e7c01971461006e57806341f654f71461008157806395607ced14610089578063cf5b31f7146100a7575b5f80fd5b61006c610067366004610835565b6100af565b005b61006c61007c3660046108e9565b610196565b61006c6102c7565b610091610342565b60405161009e919061098c565b60405180910390f35b610091610573565b5f548111156101105760405162461bcd60e51b815260206004820152602260248201527f496e6469636174656420706f736974696f6e20646f6573206e6f742065786973604482015261747360f01b60648201526084015b60405180910390fd5b5f5b5f5481101561019257818114801561014f57505f818154811061013757610137610a35565b5f91825260209091206003600490920201015460ff16155b1561018a5760015f828154811061016857610168610a35565b5f9182526020909120600490910201600301805460ff19169115159190911790555b600101610112565b5050565b61012c815111156101e95760405162461bcd60e51b815260206004820152601f60248201527f4d657373616765206c696d6974206973203330302063686172616374657273006044820152606401610107565b60408051608081018252338152602081018481529181018390525f60608201819052805460018101825590805281517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563600490920291820180546001600160a01b0319166001600160a01b0390921691909117815592519192917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e564909101906102929082610ac7565b50604082015160028201906102a79082610ac7565b50606091909101516003909101805460ff19169115159190911790555050565b5f5b5f5481101561033f575f81815481106102e4576102e4610a35565b5f918252602082206003600490920201015460ff16151590036103375760015f828154811061031557610315610a35565b5f9182526020909120600490910201600301805460ff19169115159190911790555b6001016102c9565b50565b5f80546060919067ffffffffffffffff8111156103615761036161084c565b6040519080825280602002602001820160405280156103c657816020015b6103b360405180608001604052805f6001600160a01b0316815260200160608152602001606081526020015f151581525090565b81526020019060019003908161037f5790505b5090505f5b5f5481101561056d575f81815481106103e6576103e6610a35565b5f918252602091829020604080516080810190915260049092020180546001600160a01b03168252600181018054929391929184019161042590610a49565b80601f016020809104026020016040519081016040528092919081815260200182805461045190610a49565b801561049c5780601f106104735761010080835404028352916020019161049c565b820191905f5260205f20905b81548152906001019060200180831161047f57829003601f168201915b505050505081526020016002820180546104b590610a49565b80601f01602080910402602001604051908101604052809291908181526020018280546104e190610a49565b801561052c5780601f106105035761010080835404028352916020019161052c565b820191905f5260205f20905b81548152906001019060200180831161050f57829003601f168201915b50505091835250506003919091015460ff161515602090910152825183908390811061055a5761055a610a35565b60209081029190910101526001016103cb565b50919050565b60605f805b5f548110156105c6575f818154811061059357610593610a35565b5f918252602082206003600490920201015460ff16151590036105be57816105ba81610b87565b9250505b600101610578565b505f8167ffffffffffffffff8111156105e1576105e161084c565b60405190808252806020026020018201604052801561064657816020015b61063360405180608001604052805f6001600160a01b0316815260200160608152602001606081526020015f151581525090565b8152602001906001900390816105ff5790505b5090505f805b5f5481101561082c575f818154811061066757610667610a35565b5f918252602082206003600490920201015460ff1615159003610824575f818154811061069657610696610a35565b5f918252602091829020604080516080810190915260049092020180546001600160a01b0316825260018101805492939192918401916106d590610a49565b80601f016020809104026020016040519081016040528092919081815260200182805461070190610a49565b801561074c5780601f106107235761010080835404028352916020019161074c565b820191905f5260205f20905b81548152906001019060200180831161072f57829003601f168201915b5050505050815260200160028201805461076590610a49565b80601f016020809104026020016040519081016040528092919081815260200182805461079190610a49565b80156107dc5780601f106107b3576101008083540402835291602001916107dc565b820191905f5260205f20905b8154815290600101906020018083116107bf57829003601f168201915b50505091835250506003919091015460ff161515602090910152835184908490811061080a5761080a610a35565b6020026020010181905250818061082090610b87565b9250505b60010161064c565b50909392505050565b5f60208284031215610845575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261086f575f80fd5b813567ffffffffffffffff8082111561088a5761088a61084c565b604051601f8301601f19908116603f011681019082821181831017156108b2576108b261084c565b816040528381528660208588010111156108ca575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f80604083850312156108fa575f80fd5b823567ffffffffffffffff80821115610911575f80fd5b61091d86838701610860565b93506020850135915080821115610932575f80fd5b5061093f85828601610860565b9150509250929050565b5f81518084525f5b8181101561096d57602081850181015186830182015201610951565b505f602082860101526020601f19601f83011685010191505092915050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b83811015610a2757888303603f19018552815180516001600160a01b03168452878101516080898601819052906109ec82870182610949565b9150508782015185820389870152610a048282610949565b6060938401511515969093019590955250948701949250908601906001016109b3565b509098975050505050505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680610a5d57607f821691505b60208210810361056d57634e487b7160e01b5f52602260045260245ffd5b601f821115610ac257805f5260205f20601f840160051c81016020851015610aa05750805b601f840160051c820191505b81811015610abf575f8155600101610aac565b50505b505050565b815167ffffffffffffffff811115610ae157610ae161084c565b610af581610aef8454610a49565b84610a7b565b602080601f831160018114610b28575f8415610b115750858301515b5f19600386901b1c1916600185901b178555610b7f565b5f85815260208120601f198616915b82811015610b5657888601518255948401946001909101908401610b37565b5085821015610b7357878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60018201610ba457634e487b7160e01b5f52601160045260245ffd5b506001019056fea264697066735822122046fa8e0a4d09608f338b17d5453ffb94db3da0e81ae4ab3619606e9c87861a9864736f6c63430008170033",
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
