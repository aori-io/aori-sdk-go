package ethers

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

// SignCancelOrder - Generates signature for cancel_order
func SignCancelOrder(orderHash string) (string, error) {
	key := os.Getenv("PRIVATE_KEY")
	if key == "" {
		return "", fmt.Errorf("missing PRIVATE_KEY")
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return "", err
	}

	return PersonalSign(orderHash, privateKey)
}

// PersonalSign - Returns a signature string
func PersonalSign(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}

// SignOrderHash - Returns a signature string for an AoriOrder order hash
func SignOrderHash(orderHash string) (string, error) {
	key := os.Getenv("PRIVATE_KEY")
	if key == "" {
		return "", fmt.Errorf("missing PRIVATE_KEY")
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}

	// remove 0x prefix
	orderHash = orderHash[2:]

	// Split the string into pairs of two characters
	var bytes []byte
	for i := 0; i < len(orderHash); i += 2 {
		// Convert each pair from hexadecimal to decimal
		byteVal, err := hex.DecodeString(orderHash[i : i+2])
		if err != nil {
			return "", err
		}
		bytes = append(bytes, byteVal[0])
	}

	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(string(bytes)), string(bytes))
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}
