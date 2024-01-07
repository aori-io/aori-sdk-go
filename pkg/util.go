package pkg

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/aori-io/aori-sdk-go/internal/ethers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

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
	signature, err := ethers.PersonalSign(address, privateKey)
	if err != nil {
		return nil, 0, "", "", err
	}

	return wallet, chainID.Uint64(), fromAddress, signature, nil
}

func InitializeProvider(requestUrl string) (*provider, error) {
	key := os.Getenv("PRIVATE_KEY")
	if key == "" {
		return nil, fmt.Errorf("missing PRIVATE_KEY")
	}
	address := os.Getenv("WALLET_ADDRESS")
	if address == "" {
		return nil, fmt.Errorf("missing WALLET_ADDRESS")
	}
	nodeURL := os.Getenv("NODE_URL")
	if nodeURL == "" {
		return nil, fmt.Errorf("missing NODE_URL")
	}

	wallet, chainID, walletAddr, walletSig, err := InitializeWallet(key, address, nodeURL)
	if err != nil {
		log.Fatal("Error initializing wallet:", err)
	}

	conn, _, err := websocket.DefaultDialer.Dial(requestUrl, nil)
	if err != nil {
		return nil, err
	}

	p := &provider{
		requestConn: conn,
		responseCh:  make(chan []byte),
		wallet:      wallet,
		chainId:     int(chainID),
		walletAddr:  walletAddr,
		walletSig:   walletSig,
		lastId:      1,
	}

	go func() {
		defer func(requestConn *websocket.Conn, requestChan chan []byte) {
			err := requestConn.Close()
			if err != nil {
				fmt.Println("Error closing connection: ", err)
			}

			close(requestChan)
		}(p.requestConn, p.responseCh)

		for {
			_, message, err := p.requestConn.ReadMessage()
			if err != nil {
				log.Println("Error receiving message:", err)
				return
			}
			p.responseCh <- message
		}
	}()

	return p, nil
}
