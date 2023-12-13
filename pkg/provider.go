package pkg

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/gorilla/websocket"
	"sync"
)

type AoriProvider interface {
	Send()
	Ping()
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
	feedConn    *websocket.Conn
	wallet      *accounts.Wallet
	chainID     uint64
	lastID      *sync.Mutex
	walletAddr  string
	walletSig   string
}

func NewAoriProvider() *provider {
	// TODO
	return &provider{}
}

func NewAoriProviderWithURL(requestURL, feedURL string) *provider {
	return &provider{}
}

func (p *provider) Send() {
	// TODO: impl
}

func (p *provider) Ping() {
	// TODO: impl
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
