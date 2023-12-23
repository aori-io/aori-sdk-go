package util

import (
	"crypto/ecdsa"
	"fmt"
	types2 "github.com/aori-io/aori-sdk-go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"os"
	"strings"
)

// SignOrder - Signs Order
func SignOrder(order types2.OrderParameters, chainId int) (string, error) {
	message := map[string]interface{}{
		"offerer": order.Offerer,
		"zone":    order.Zone,
		"offer": []map[string]interface{}{{
			"itemType":             fmt.Sprintf("%d", order.Offer[0].ItemType),
			"token":                order.Offer[0].Token,
			"identifierOrCriteria": order.Offer[0].IdentifierOrCriteria,
			"startAmount":          order.Offer[0].StartAmount,
			"endAmount":            order.Offer[0].EndAmount,
		}},
		"consideration": []map[string]interface{}{{
			"itemType":             fmt.Sprintf("%d", order.Consideration[0].ItemType),
			"token":                order.Consideration[0].Token,
			"identifierOrCriteria": order.Consideration[0].IdentifierOrCriteria,
			"startAmount":          order.Consideration[0].StartAmount,
			"endAmount":            order.Consideration[0].EndAmount,
			"recipient":            order.Consideration[0].Recipient,
		}},
		"orderType":                       fmt.Sprintf("%d", order.OrderType),
		"startTime":                       order.StartTime,
		"endTime":                         order.EndTime,
		"zoneHash":                        common.Hex2Bytes(strings.TrimPrefix(order.ZoneHash, "0x")),
		"salt":                            order.Salt,
		"conduitKey":                      common.Hex2Bytes(strings.TrimPrefix(order.ConduitKey, "0x")),
		"totalOriginalConsiderationItems": fmt.Sprintf("%d", order.TotalOriginalConsiderationItems),
		"counter":                         order.Counter,
	}

	domain := apitypes.TypedDataDomain{
		Name:              "Seaport",
		Version:           types2.CurrentSeaportVersion,
		ChainId:           math.NewHexOrDecimal256(int64(chainId)),
		VerifyingContract: types2.SeaportAddress,
	}
	typedData := apitypes.TypedData{
		Types:       types2.Eip712OrderType,
		PrimaryType: "OrderComponents",
		Domain:      domain,
		Message:     message,
	}

	key := os.Getenv("PRIVATE_KEY")
	if key == "" {
		return "", fmt.Errorf("missing PRIVATE_KEY")
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return "", err
	}

	sigBytes, err := SignTypedData(typedData, privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing typed data: %s", err)
	}
	signature := fmt.Sprintf("0x%s", common.Bytes2Hex(sigBytes))

	return signature, nil
}

// SignTypedData - Sign typed data
func SignTypedData(typedData apitypes.TypedData, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	hash, err := EncodeForSigning(typedData)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27

	return sig, nil
}

// EncodeForSigning - Encoding the typed data
func EncodeForSigning(typedData apitypes.TypedData) ([]byte, error) {
	hash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		fmt.Println("HHHH:", err)
		return nil, err
	}

	return hash, nil
}

// SignCancelOrder - Generates signature for cancel_order
func SignCancelOrder(orderId string) (string, error) {
	key := os.Getenv("PRIVATE_KEY")
	if key == "" {
		return "", fmt.Errorf("missing PRIVATE_KEY")
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return "", err
	}

	hash := crypto.Keccak256Hash([]byte(orderId))

	sigBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}

	sigBytes[64] += 27

	signature := fmt.Sprintf("0x%s", common.Bytes2Hex(sigBytes))

	return signature, nil
}
