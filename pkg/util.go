package pkg

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MakeOrderInput struct {
	SellToken  string
	SellAmount string
	BuyToken   string
	BuyAmount  string
}

func InitializeWallet(key, address, nodeURL string) (*bind.TransactOpts, uint64, string, string, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, 0, "", "", err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, 0, "", "", err
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, 0, "", "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, 0, "", "", fmt.Errorf("error casting public key to ECDSA")
	}

	wallet, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, 0, "", "", err
	}
	wallet.Nonce = nil // Set the nonce to nil for auto calculation

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	signature, err := crypto.Sign(crypto.Keccak256([]byte(address)), privateKey)
	if err != nil {
		return nil, 0, "", "", err
	}

	return wallet, chainID.Uint64(), fromAddress, fmt.Sprintf("0x%x", signature), nil
}
