package ethers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg/types"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

func SignOrderHash(order types.OrderParameters) (string, error) {
	key := os.Getenv("PRIVATE_KEY")
	if key == "" {
		return "", fmt.Errorf("missing PRIVATE_KEY")
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return "", err
	}

	// Convert order to JSON string
	orderJSON, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	// Calculate hash of JSON string
	hash := sha256.Sum256(orderJSON)

	// Convert hash to hexadecimal string
	hashStr := hex.EncodeToString(hash[:])
	return PersonalSign(hashStr, privateKey)
}
