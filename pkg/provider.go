package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/internal/types"
	"github.com/aori-io/aori-sdk-go/internal/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"sync"
	"time"
)

type AoriProvider interface {
	Send(msg []byte) error
	Receive() ([]byte, error)
	Ping() (string, error)
	AuthWallet() (types.AuthWalletResponse, error)
	CheckAuth(jwt string) (string, error)
	ViewOrderbook(chainId int, base, quote, side string) (types.AoriViewOrderbookResponse, error)
	MakeOrder()
	MakePrivateOrder()
	TakeOrder()
	CancelOrder()
	SubscribeOrderbook()
	AccountOrders()
	OrderStatus()
	CancelAllOrders()
}

type provider struct {
	requestConn *websocket.Conn
	responseCh  chan []byte
	mu          sync.Mutex
	wallet      *bind.TransactOpts
	chainId     int
	lastId      int
	walletAddr  string
	walletSig   string
}

func NewAoriProvider() (*provider, error) {
	fmt.Println("Initializing Bot")

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

	conn, _, err := websocket.DefaultDialer.Dial(types.RequestURL, nil)
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

func (p *provider) Send(msg []byte) error {
	fmt.Println("Sending message...")
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.requestConn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return err
	}

	p.lastId++

	return nil
}

func (p *provider) Receive() ([]byte, error) {
	select {
	case msg := <-p.responseCh:
		return msg, nil
	case <-time.After(5 * time.Second): // Adjust timeout as needed
		return nil, fmt.Errorf("timeout: no response received")
	}
}

func (p *provider) Ping() (string, error) {
	req, err := util.CreatePingPayload(p.lastId)
	if err != nil {
		return "", fmt.Errorf("error creating ping payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return "", fmt.Errorf("error sending ping request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return "", fmt.Errorf("error getting response: %s", err)
	}

	return string(res), nil
}

func (p *provider) AuthWallet() (types.AuthWalletResponse, error) {
	var authWalletResponse types.AuthWalletResponse

	req, err := util.CreateAuthWalletPayload(p.lastId, p.walletAddr, p.walletSig)
	if err != nil {
		return authWalletResponse, fmt.Errorf("auth_wallet error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return authWalletResponse, fmt.Errorf("auth_wallet error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return authWalletResponse, fmt.Errorf("auth_wallet error getting response: %s", err)
	}

	err = json.Unmarshal(res, &authWalletResponse)
	if err != nil {
		return authWalletResponse, fmt.Errorf("auth_wallet error getting unmarshalling: %s", err)
	}

	return authWalletResponse, nil
}

func (p *provider) CheckAuth(jwt string) (string, error) {
	req, err := util.CreateCheckAuthPayload(p.lastId, jwt)
	if err != nil {
		return "", fmt.Errorf("check_auth error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return "", fmt.Errorf("check_auth error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return "", fmt.Errorf("check_auth error getting response: %s", err)
	}

	return string(res), nil
}

func (p *provider) ViewOrderbook(chainId int, base, quote, side string) (types.AoriViewOrderbookResponse, error) {
	var viewOrderbookResponse types.AoriViewOrderbookResponse

	req, err := util.CreateViewOrderbookPayload(p.lastId, chainId, base, quote, side)
	if err != nil {
		return viewOrderbookResponse, fmt.Errorf("view_orderbook error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return viewOrderbookResponse, fmt.Errorf("view_orderbook error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return viewOrderbookResponse, fmt.Errorf("view_orderbook error getting response: %s", err)
	}

	err = json.Unmarshal(res, &viewOrderbookResponse)
	if err != nil {
		return viewOrderbookResponse, fmt.Errorf("view_orderbook error getting unmarshalling: %s", err)
	}

	return viewOrderbookResponse, nil
}

func (p *provider) MakeOrder(orderParams MakeOrderInput) (string, error) {
	fmt.Println("CHAIN: ", p.chainId)

	req, err := util.CreateMakeOrderPayload(p.lastId, p.chainId, p.walletAddr, orderParams.SellToken, orderParams.SellAmount, orderParams.BuyToken, orderParams.BuyAmount)
	if err != nil {
		return "", fmt.Errorf("make_order error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return "", fmt.Errorf("make_order error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return "", fmt.Errorf("make_order error getting response: %s", err)
	}

	return string(res), nil
}

func (p *provider) MakePrivateOrder() {
	// TODO: impl
}

func (p *provider) TakeOrder() {
	// TODO: impl
}

func (p *provider) CancelOrder() {
	// TODO: impl
}

func (p *provider) SubscribeOrderbook() {
	// TODO: impl
}

func (p *provider) AccountOrders() {
	// TODO: impl
}

func (p *provider) OrderStatus() {
	// TODO: impl
}

func (p *provider) CancelAllOrders() {
	// TODO: impl
}
