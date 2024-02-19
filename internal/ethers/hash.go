package ethers

import (
	"encoding/hex"
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// CalculateOrderHash - calculates the hash of an AoriOrder
func CalculateOrderHash(order types.OrderParameters) (string, error) {
	types := []string{"address", "address", "uint256", "uint256", "address", "address", "uint256", "uint256", "address", "uint256", "uint256", "uint256", "uint256", "bool"}
	values := []interface{}{
		order.Offerer,
		order.InputToken,
		order.InputAmount,
		fmt.Sprintf("%v", order.InputChainID),
		order.InputZone,
		order.OutputToken,
		order.OutputAmount,
		fmt.Sprintf("%v", order.OutputChainID),
		order.OutputZone,
		order.StartTime,
		order.EndTime,
		order.Salt,
		fmt.Sprintf("%v", order.Counter),
		order.ToWithdraw,
	}

	hash := solsha3.SoliditySHA3(types, values)
	return "0x" + hex.EncodeToString(hash), nil
}
