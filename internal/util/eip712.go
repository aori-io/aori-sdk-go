package util

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"os"
	"strings"
)

// SignOrder - Signs Order
func SignOrder(order types.OrderParameters, chainId int) (string, error) {
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
		"orderType":  fmt.Sprintf("%d", order.OrderType),
		"startTime":  order.StartTime,
		"endTime":    order.EndTime,
		"zoneHash":   common.Hex2Bytes(strings.TrimPrefix(order.ZoneHash, "0x")),
		"salt":       order.Salt,
		"conduitKey": common.Hex2Bytes(strings.TrimPrefix(order.ConduitKey, "0x")),
		"counter":    order.Counter,
	}

	domain := apitypes.TypedDataDomain{
		Name:              "Seaport",
		Version:           types.CurrentSeaportVersion,
		ChainId:           math.NewHexOrDecimal256(int64(chainId)),
		VerifyingContract: types.SeaportAddress,
	}
	typedData := apitypes.TypedData{
		Types:       types.Eip712OrderType,
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

	return SignTypedData(typedData, privateKey)
}

// SignTypedData - Sign typed data
func SignTypedData(typedData apitypes.TypedData, privateKey *ecdsa.PrivateKey) (string, error) {
	hash, err := EncodeForSigning(typedData)
	if err != nil {
		return "", err
	}

	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27

	return hexutil.Encode(signatureBytes), nil
}

// EncodeForSigning - Encoding the typed data
func EncodeForSigning(typedData apitypes.TypedData) (*common.Hash, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := crypto.Keccak256Hash(rawData)
	return &hash, nil
}
