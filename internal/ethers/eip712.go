package ethers

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
	var offerMessage []map[string]interface{}
	var considerationMessage []map[string]interface{}

	for _, offer := range order.Offer {
		m := map[string]interface{}{
			"itemType":             fmt.Sprintf("%d", offer.ItemType),
			"token":                offer.Token,
			"identifierOrCriteria": offer.IdentifierOrCriteria,
			"startAmount":          offer.StartAmount,
			"endAmount":            offer.EndAmount,
		}
		offerMessage = append(offerMessage, m)
	}

	for _, consideration := range order.Consideration {
		m := map[string]interface{}{
			"itemType":             fmt.Sprintf("%d", consideration.ItemType),
			"token":                consideration.Token,
			"identifierOrCriteria": consideration.IdentifierOrCriteria,
			"startAmount":          consideration.StartAmount,
			"endAmount":            consideration.EndAmount,
			"recipient":            consideration.Recipient,
		}
		considerationMessage = append(considerationMessage, m)
	}

	message := map[string]interface{}{
		"offerer":       order.Offerer,
		"zone":          order.Zone,
		"offer":         offerMessage,
		"consideration": considerationMessage,
		"orderType":     fmt.Sprintf("%d", order.OrderType),
		"startTime":     order.StartTime,
		"endTime":       order.EndTime,
		"zoneHash":      common.Hex2Bytes(strings.TrimPrefix(order.ZoneHash, "0x")),
		"salt":          order.Salt,
		"conduitKey":    common.Hex2Bytes(strings.TrimPrefix(order.ConduitKey, "0x")),
		"counter":       order.Counter,
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
