package util

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"os"
)

// SignOrder - Signs Order
func SignOrder(order types.OrderComponents) (string, error) {
	message := make(map[string]interface{})

	orderComponentsBytes, _ := json.Marshal(order)
	err := json.Unmarshal(orderComponentsBytes, &message)
	if err != nil {
		return "", err
	}

	domain := apitypes.TypedDataDomain{
		Name:              "Seaport",
		Version:           types.CurrentSeaportVersion,
		ChainId:           math.NewHexOrDecimal256(5),
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
	sig, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27
	return sig, nil
}

// EncodeForSigning - Encoding the typed data
func EncodeForSigning(typedData apitypes.TypedData) (common.Hash, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		fmt.Println("TEST", typedData.Message["offerer"])
		return common.Hash{}, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := common.BytesToHash(crypto.Keccak256(rawData))
	return hash, nil
}
