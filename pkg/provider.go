package pkg

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/internal"
	"github.com/aori-io/aori-sdk-go/internal/types"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type AoriProvider interface {
	Send(msg []byte) error
	Receive() ([]byte, error)
	Ping() (string, error)
	CheckAuth()
	ViewOrderbook()
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
}

func NewAoriProvider() (*provider, error) {
	fmt.Println("Initializing Bot")
	conn, _, err := websocket.DefaultDialer.Dial(types.RequestURL, nil)
	if err != nil {
		return nil, err
	}

	p := &provider{
		requestConn: conn,
		responseCh:  make(chan []byte),
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
	req, err := internal.CreatePingPayload()
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

func (p *provider) CheckAuth() {
	// TODO: impl
}

func (p *provider) ViewOrderbook() {
	// TODO: impl
}

func (p *provider) MakeOrder() {
	// TODO: impl
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
