package ethers

import (
	"crypto/ecdsa"
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
