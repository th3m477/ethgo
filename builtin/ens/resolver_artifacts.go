package ens

import (
	"encoding/hex"
	"fmt"

	"github.com/th3m477/ethgo/abi"
)

var abiResolver *abi.ABI

// ResolverAbi returns the abi of the Resolver contract
func ResolverAbi() *abi.ABI {
	return abiResolver
}

var binResolver []byte

// ResolverBin returns the bin of the Resolver contract
func ResolverBin() []byte {
	return binResolver
}

func init() {
	var err error
	abiResolver, err = abi.NewABI(abiResolverStr)
	if err != nil {
		panic(fmt.Errorf("cannot parse Resolver abi: %v", err))
	}
	if len(binResolverStr) != 0 {
		binResolver, err = hex.DecodeString(binResolverStr[2:])
		if err != nil {
			panic(fmt.Errorf("cannot parse Resolver bin: %v", err))
		}
	}
}

var binResolverStr = "0x6060604052341561000f57600080fd5b6040516020806111b28339810160405280805160008054600160a060020a03909216600160a060020a0319909216919091179055505061115e806100546000396000f3006060604052600436106100ab5763ffffffff60e060020a60003504166301ffc9a781146100b057806310f13a8c146100e45780632203ab561461017e57806329cd62ea146102155780632dff6941146102315780633b3b57de1461025957806359d1d43c1461028b578063623195b014610358578063691f3431146103b457806377372213146103ca578063c3d014d614610420578063c869023314610439578063d5fa2b0014610467575b600080fd5b34156100bb57600080fd5b6100d0600160e060020a031960043516610489565b604051901515815260200160405180910390f35b34156100ef57600080fd5b61017c600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506105f695505050505050565b005b341561018957600080fd5b610197600435602435610807565b60405182815260406020820181815290820183818151815260200191508051906020019080838360005b838110156101d95780820151838201526020016101c1565b50505050905090810190601f1680156102065780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b341561022057600080fd5b61017c600435602435604435610931565b341561023c57600080fd5b610247600435610a30565b60405190815260200160405180910390f35b341561026457600080fd5b61026f600435610a46565b604051600160a060020a03909116815260200160405180910390f35b341561029657600080fd5b6102e1600480359060446024803590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610a6195505050505050565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561031d578082015183820152602001610305565b50505050905090810190601f16801561034a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561036357600080fd5b61017c600480359060248035919060649060443590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610b8095505050505050565b34156103bf57600080fd5b6102e1600435610c7c565b34156103d557600080fd5b61017c600480359060446024803590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610d4295505050505050565b341561042b57600080fd5b61017c600435602435610e8c565b341561044457600080fd5b61044f600435610f65565b60405191825260208201526040908101905180910390f35b341561047257600080fd5b61017c600435600160a060020a0360243516610f82565b6000600160e060020a031982167f3b3b57de0000000000000000000000000000000000000000000000000000000014806104ec5750600160e060020a031982167fd8389dc500000000000000000000000000000000000000000000000000000000145b806105205750600160e060020a031982167f691f343100000000000000000000000000000000000000000000000000000000145b806105545750600160e060020a031982167f2203ab5600000000000000000000000000000000000000000000000000000000145b806105885750600160e060020a031982167fc869023300000000000000000000000000000000000000000000000000000000145b806105bc5750600160e060020a031982167f59d1d43c00000000000000000000000000000000000000000000000000000000145b806105f05750600160e060020a031982167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b600080548491600160a060020a033381169216906302571be39084906040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561064f57600080fd5b6102c65a03f1151561066057600080fd5b50505060405180519050600160a060020a031614151561067f57600080fd5b6000848152600160205260409081902083916005909101908590518082805190602001908083835b602083106106c65780518252601f1990920191602091820191016106a7565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902090805161070a929160200190611085565b50826040518082805190602001908083835b6020831061073b5780518252601f19909201916020918201910161071c565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051908190039020847fd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a75508560405160208082528190810183818151815260200191508051906020019080838360005b838110156107c75780820151838201526020016107af565b50505050905090810190601f1680156107f45780820380516001836020036101000a031916815260200191505b509250505060405180910390a350505050565b6000610811611103565b60008481526001602081905260409091209092505b838311610924578284161580159061085f5750600083815260068201602052604081205460026000196101006001841615020190911604115b15610919578060060160008481526020019081526020016000208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561090d5780601f106108e25761010080835404028352916020019161090d565b820191906000526020600020905b8154815290600101906020018083116108f057829003601f168201915b50505050509150610929565b600290920291610826565b600092505b509250929050565b600080548491600160a060020a033381169216906302571be39084906040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561098a57600080fd5b6102c65a03f1151561099b57600080fd5b50505060405180519050600160a060020a03161415156109ba57600080fd5b6040805190810160409081528482526020808301859052600087815260019091522060030181518155602082015160019091015550837f1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46848460405191825260208201526040908101905180910390a250505050565b6000908152600160208190526040909120015490565b600090815260016020526040902054600160a060020a031690565b610a69611103565b60008381526001602052604090819020600501908390518082805190602001908083835b60208310610aac5780518252601f199092019160209182019101610a8d565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b735780601f10610b4857610100808354040283529160200191610b73565b820191906000526020600020905b815481529060010190602001808311610b5657829003601f168201915b5050505050905092915050565b600080548491600160a060020a033381169216906302571be39084906040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610bd957600080fd5b6102c65a03f11515610bea57600080fd5b50505060405180519050600160a060020a0316141515610c0957600080fd5b6000198301831615610c1a57600080fd5b60008481526001602090815260408083208684526006019091529020828051610c47929160200190611085565b5082847faa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe360405160405180910390a350505050565b610c84611103565b6001600083600019166000191681526020019081526020016000206002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d365780601f10610d0b57610100808354040283529160200191610d36565b820191906000526020600020905b815481529060010190602001808311610d1957829003601f168201915b50505050509050919050565b600080548391600160a060020a033381169216906302571be39084906040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610d9b57600080fd5b6102c65a03f11515610dac57600080fd5b50505060405180519050600160a060020a0316141515610dcb57600080fd5b6000838152600160205260409020600201828051610ded929160200190611085565b50827fb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f78360405160208082528190810183818151815260200191508051906020019080838360005b83811015610e4d578082015183820152602001610e35565b50505050905090810190601f168015610e7a5780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b600080548391600160a060020a033381169216906302571be39084906040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610ee557600080fd5b6102c65a03f11515610ef657600080fd5b50505060405180519050600160a060020a0316141515610f1557600080fd5b6000838152600160208190526040918290200183905583907f0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc9084905190815260200160405180910390a2505050565b600090815260016020526040902060038101546004909101549091565b600080548391600160a060020a033381169216906302571be39084906040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610fdb57600080fd5b6102c65a03f11515610fec57600080fd5b50505060405180519050600160a060020a031614151561100b57600080fd5b60008381526001602052604090819020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851617905583907f52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd290849051600160a060020a03909116815260200160405180910390a2505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110c657805160ff19168380011785556110f3565b828001600101855582156110f3579182015b828111156110f35782518255916020019190600101906110d8565b506110ff929150611115565b5090565b60206040519081016040526000815290565b61112f91905b808211156110ff576000815560010161111b565b905600a165627a7a723058201ecacbc445b9fbcd91b0ab164389f69d7283b856883bc7437eeed1008345a4920029"

var abiResolverStr = `[{"constant":true,"inputs":[{"name":"interfaceID","type":"bytes4"}],"name":"supportsInterface","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"node","type":"bytes32"},{"name":"contentTypes","type":"uint256"}],"name":"ABI","outputs":[{"name":"contentType","type":"uint256"},{"name":"data","type":"bytes"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"x","type":"bytes32"},{"name":"y","type":"bytes32"}],"name":"setPubkey","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"content","outputs":[{"name":"ret","type":"bytes32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"addr","outputs":[{"name":"ret","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"contentType","type":"uint256"},{"name":"data","type":"bytes"}],"name":"setABI","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"name","outputs":[{"name":"ret","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"name","type":"string"}],"name":"setName","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"hash","type":"bytes32"}],"name":"setContent","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"node","type":"bytes32"}],"name":"pubkey","outputs":[{"name":"x","type":"bytes32"},{"name":"y","type":"bytes32"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"node","type":"bytes32"},{"name":"addr","type":"address"}],"name":"setAddr","outputs":[],"payable":false,"type":"function"},{"inputs":[{"name":"ensAddr","type":"address"}],"payable":false,"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"a","type":"address"}],"name":"AddrChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"hash","type":"bytes32"}],"name":"ContentChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"name","type":"string"}],"name":"NameChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":true,"name":"contentType","type":"uint256"}],"name":"ABIChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"node","type":"bytes32"},{"indexed":false,"name":"x","type":"bytes32"},{"indexed":false,"name":"y","type":"bytes32"}],"name":"PubkeyChanged","type":"event"}]`
