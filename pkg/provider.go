package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/internal/util"
	"github.com/aori-io/aori-sdk-go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type AoriProvider interface {
	Send(msg []byte) error
	Receive() ([]byte, error)
	SendFeed(msg []byte) error
	ReceiveFeed() ([]byte, error)
	AccountOrders(apiKey string) (*types.AccountOrdersResponse, error)
	AuthWallet() (*types.AuthWalletResponse, error)
	CancelAllOrders() (string, error)
	CancelOrder(orderId, apiKey string) (string, error)
	CheckAuth(jwt string) (string, error)
	MakeOrder(orderParams types.MakeOrderInput) (*types.MakeOrderResponse, error)
	MakePrivateOrder(orderParams types.MakeOrderInput) (*types.MakeOrderResponse, error)
	OnSubscribe(event types.SubscriptionEvent, handler func(payload []byte))
	OrderStatus(orderHash string) (*types.OrderStatusResponse, error)
	Ping() (string, error)
	SubscribeOrderbook() (string, error)
	TakeOrder(orderParams types.OrderParameters, orderHash string, seatId int) (string, error)
	ViewOrderbook(query types.ViewOrderbookParams) (*types.AoriViewOrderbookResponse, error)
}

type provider struct {
	requestConn *websocket.Conn
	feedConn    *websocket.Conn
	requestCh   chan []byte
	feedCh      chan []byte
	mu          sync.Mutex
	wallet      *bind.TransactOpts
	chainId     int
	lastId      int
	walletAddr  string
	walletSig   string
}

func NewAoriProvider() (*provider, error) {
	fmt.Println("Initializing Bot")

	return InitializeProvider(types.RequestURL, types.MarketFeedURL)
}

func NewAoriProviderWithURL(requestURL, feedURL string) (*provider, error) {
	fmt.Println("Initializing Bot")

	return InitializeProvider(requestURL, feedURL)
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
	case msg := <-p.requestCh:
		return msg, nil
	case <-time.After(5 * time.Second): // Adjust timeout as needed
		return nil, fmt.Errorf("timeout: no response received")
	}
}

func (p *provider) SendFeed(msg []byte) error {
	fmt.Println("Sending message...")
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.feedConn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return err
	}

	p.lastId++

	return nil
}

func (p *provider) ReceiveFeed() ([]byte, error) {
	msg := <-p.feedCh
	return msg, nil
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

func (p *provider) AuthWallet() (*types.AuthWalletResponse, error) {
	var authWalletResponse types.AuthWalletResponse

	req, err := util.CreateAuthWalletPayload(p.lastId, p.walletAddr, p.walletSig)
	if err != nil {
		return nil, fmt.Errorf("auth_wallet error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return nil, fmt.Errorf("auth_wallet error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return nil, fmt.Errorf("auth_wallet error getting response: %s", err)
	}

	err = json.Unmarshal(res, &authWalletResponse)
	if err != nil {
		return nil, fmt.Errorf("auth_wallet error getting unmarshalling: %s", err)
	}

	return &authWalletResponse, nil
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

func (p *provider) ViewOrderbook(query types.ViewOrderbookParams) (*types.AoriViewOrderbookResponse, error) {
	var viewOrderbookResponse types.AoriViewOrderbookResponse

	req, err := util.CreateViewOrderbookPayload(p.lastId, query)
	if err != nil {
		return nil, fmt.Errorf("view_orderbook error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return nil, fmt.Errorf("view_orderbook error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return nil, fmt.Errorf("view_orderbook error getting response: %s", err)
	}

	err = json.Unmarshal(res, &viewOrderbookResponse)
	if err != nil {
		return nil, fmt.Errorf("view_orderbook error getting unmarshalling: %s", err)
	}

	return &viewOrderbookResponse, nil
}

func (p *provider) MakeOrder(orderParams types.MakeOrderInput) (*types.MakeOrderResponse, error) {
	var makeOrderResponse types.MakeOrderResponse
	req, err := util.CreateMakeOrderPayload(p.lastId, p.chainId, p.walletAddr, orderParams, true)
	if err != nil {
		return nil, fmt.Errorf("make_order error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return nil, fmt.Errorf("make_order error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return nil, fmt.Errorf("make_order error getting response: %s", err)
	}

	err = json.Unmarshal(res, &makeOrderResponse)
	if err != nil {
		return nil, err
	}

	return &makeOrderResponse, nil
}

func (p *provider) MakePrivateOrder(orderParams types.MakeOrderInput) (*types.MakeOrderResponse, error) {
	var makeOrderResponse types.MakeOrderResponse

	req, err := util.CreateMakeOrderPayload(p.lastId, p.chainId, p.walletAddr, orderParams, false)
	if err != nil {
		return nil, fmt.Errorf("make_order error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return nil, fmt.Errorf("make_order error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return nil, fmt.Errorf("make_order error getting response: %s", err)
	}

	err = json.Unmarshal(res, &makeOrderResponse)
	if err != nil {
		return nil, err
	}

	return &makeOrderResponse, nil
}

func (p *provider) OnSubscribe(event types.SubscriptionEvent, handler func(payload []byte)) error {
	fmt.Println("Listening to the orderbook")
	//for {
	msg, err := p.ReceiveFeed()
	if err != nil {
		return err
	}
	handler(msg)
	//}
	return nil
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
		return "", fmt.Errorf("subscribe_orderbook error creating payload: %s", err)
	}
	err = p.SendFeed(req)
	if err != nil {
		return "", fmt.Errorf("subscribe_orderbook error sending request: %s", err)
	}

	res, err := p.ReceiveFeed()
	if err != nil {
		return "", fmt.Errorf("subscribe_orderbook error getting response: %s", err)
	}

	return string(res), nil
}

func (p *provider) AccountOrders(apiKey string) (*types.AccountOrdersResponse, error) {
	var accountOrderResponse types.AccountOrdersResponse

	req, err := util.CreateAccountOrdersPayload(p.lastId, p.walletAddr, p.walletSig, apiKey)
	if err != nil {
		return nil, fmt.Errorf("account_orders error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return nil, fmt.Errorf("account_orders error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return nil, fmt.Errorf("account_orders error getting response: %s", err)
	}

	err = json.Unmarshal(res, &accountOrderResponse)
	if err != nil {
		return nil, fmt.Errorf("account_orders error getting response: %s", err)
	}

	return &accountOrderResponse, nil
}

func (p *provider) OrderStatus(orderHash string) (*types.OrderStatusResponse, error) {
	var orderStatusResponse types.OrderStatusResponse

	req, err := util.CreateOrderStatusPayload(p.lastId, orderHash)
	if err != nil {
		return nil, fmt.Errorf("order_status error creating payload: %s", err)
	}
	err = p.Send(req)
	if err != nil {
		return nil, fmt.Errorf("order_status error sending request: %s", err)
	}

	res, err := p.Receive()
	if err != nil {
		return nil, fmt.Errorf("order_status error getting response: %s", err)
	}

	err = json.Unmarshal(res, &orderStatusResponse)
	if err != nil {
		return nil, fmt.Errorf("order_status error getting unmarshalling: %s", err)
	}

	return &orderStatusResponse, nil
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
