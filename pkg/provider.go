package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/internal/util"
	"github.com/aori-io/aori-sdk-go/pkg/types"
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
	ViewOrderbook(query types.ViewOrderbookParams) (types.AoriViewOrderbookResponse, error)
	MakeOrder(orderParams types.MakeOrderInput) (string, error)
	MakePrivateOrder(orderParams types.MakeOrderInput) (string, error)
	TakeOrder(orderParams types.OrderParameters, orderHash string, seatId int) (string, error)
	CancelOrder(orderId, apiKey string) (string, error)
	SubscribeOrderbook() (string, error)
	AccountOrders(apiKey string) (string, error)
	OrderStatus(orderHash string) (types.OrderStatusResponse, error)
	CancelAllOrders() (string, error)
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

func (p *provider) ViewOrderbook(query types.ViewOrderbookParams) (types.AoriViewOrderbookResponse, error) {
	var viewOrderbookResponse types.AoriViewOrderbookResponse

	req, err := util.CreateViewOrderbookPayload(p.lastId, query)
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

func (p *provider) MakeOrder(orderParams types.MakeOrderInput) (string, error) {
	req, err := util.CreateMakeOrderPayload(p.lastId, p.chainId, p.walletAddr, orderParams, true)
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

func (p *provider) MakePrivateOrder(orderParams types.MakeOrderInput) (string, error) {
	req, err := util.CreateMakeOrderPayload(p.lastId, p.chainId, p.walletAddr, orderParams, false)
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

func (p *provider) TakeOrder(orderParams types.OrderParameters, orderHash string, seatId int) (string, error) {
	req, err := util.CreateTakeOrderPayload(p.lastId, p.chainId, seatId, p.walletAddr, orderHash, orderParams)
	if err != nil {
		return "", fmt.Errorf("take_order error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return "", fmt.Errorf("take_order error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return "", fmt.Errorf("take_order error getting response: %s", err)
	}

	return string(res), nil
}

func (p *provider) CancelOrder(orderId, apiKey string) (string, error) {
	req, err := util.CreateCancelOrderPayload(p.lastId, orderId, apiKey)
	if err != nil {
		return "", fmt.Errorf("cancel_order error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return "", fmt.Errorf("cancel_order error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return "", fmt.Errorf("cancel_order error getting response: %s", err)
	}

	return string(res), nil
}

func (p *provider) SubscribeOrderbook() (string, error) {
	req, err := util.CreateSubscribeOrderbookPayload(p.lastId)
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

func (p *provider) AccountOrders(apiKey string) (string, error) {
	fmt.Println("ddd", p.walletSig)
	req, err := util.CreateAccountOrdersPayload(p.lastId, p.walletAddr, p.walletSig, apiKey)
	if err != nil {
		return "", fmt.Errorf("account_orders error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return "", fmt.Errorf("account_orders error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return "", fmt.Errorf("account_orders error getting response: %s", err)
	}

	return string(res), nil
}

func (p *provider) OrderStatus(orderHash string) (types.OrderStatusResponse, error) {
	var orderStatusResponse types.OrderStatusResponse

	req, err := util.CreateOrderStatusPayload(p.lastId, orderHash)
	if err != nil {
		return orderStatusResponse, fmt.Errorf("order_status error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return orderStatusResponse, fmt.Errorf("order_status error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return orderStatusResponse, fmt.Errorf("order_status error getting response: %s", err)
	}

	fmt.Println("RES: ", string(res))

	err = json.Unmarshal(res, &orderStatusResponse)
	if err != nil {
		return orderStatusResponse, fmt.Errorf("order_status error getting unmarshalling: %s", err)
	}

	return orderStatusResponse, nil
}

func (p *provider) CancelAllOrders() (string, error) {
	req, err := util.CreateCancelAllOrdersPayload(p.lastId, p.walletAddr)
	if err != nil {
		return "", fmt.Errorf("cancel_all_orders error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return "", fmt.Errorf("cancel_all_orders error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return "", fmt.Errorf("cancel_all_orders error getting response: %s", err)
	}

	return string(res), nil
}
