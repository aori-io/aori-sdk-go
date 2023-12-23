package types

type AoriResponse struct {
	ID     *uint64   `json:"id"`
	Result AoriEvent `json:"result"`
}

type OrderCreationData struct {
	Parameters OrderParameters `json:"parameters"`
	Signature  string          `json:"signature"`
}

type AoriEvent struct {
	Subscribed     string              `json:"Subscribed"`
	OrderCancelled *OrderCancelledData `json:"OrderCancelled"`
	OrderCreated   *OrderCreatedData   `json:"OrderCreated"`
	OrderTaken     *OrderTakenData     `json:"OrderTaken"`
}

type OrderCreatedData struct {
	Order         OrderCreationData `json:"order"`
	OrderHash     string            `json:"orderHash"`
	InputToken    string            `json:"inputToken"`
	OutputToken   string            `json:"outputToken"`
	InputAmount   string            `json:"inputAmount"`
	OutputAmount  string            `json:"outputAmount"`
	ChainID       int64             `json:"chainId"`
	Active        bool              `json:"active"`
	CreatedAt     int               `json:"createdAt"`
	LastUpdatedAt int               `json:"lastUpdatedAt"`
	IsPublic      bool              `json:"isPublic"`
	Rate          *float64          `json:"rate,omitempty"`
}

type OrderCancelledData struct {
	Order         OrderCreationData `json:"order"`
	OrderHash     string            `json:"orderHash"`
	InputToken    string            `json:"inputToken"`
	OutputToken   string            `json:"outputToken"`
	InputAmount   uint64            `json:"inputAmount"`
	OutputAmount  uint64            `json:"outputAmount"`
	ChainID       int64             `json:"chainId"`
	Active        bool              `json:"active"`
	CreatedAt     uint64            `json:"createdAt"`
	LastUpdatedAt uint64            `json:"lastUpdatedAt"`
	IsPublic      bool              `json:"isPublic"`
}

type OrderTakenData struct {
	Order         OrderCreationData `json:"order"`
	OrderHash     string            `json:"orderHash"`
	InputToken    string            `json:"inputToken"`
	OutputToken   string            `json:"outputToken"`
	InputAmount   uint64            `json:"inputAmount"`
	OutputAmount  uint64            `json:"outputAmount"`
	ChainID       int64             `json:"chainId"`
	Active        bool              `json:"active"`
	CreatedAt     uint64            `json:"createdAt"`
	LastUpdatedAt uint64            `json:"lastUpdatedAt"`
	IsPublic      bool              `json:"isPublic"`
	TakenAt       uint64            `json:"takenAt"`
}

type PingRequest struct {
	Id      int      `json:"id"`
	JsonRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type AuthWalletRequest struct {
	Id      int                `json:"id"`
	JsonRPC string             `json:"jsonrpc"`
	Method  string             `json:"method"`
	Params  []AuthWalletParams `json:"params"`
}

type AuthWalletParams struct {
	Address   string `json:"address"`
	Signature string `json:"signature"`
}

type CheckAuthRequest struct {
	Id      int               `json:"id"`
	JsonRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []CheckAuthParams `json:"params"`
}

type CheckAuthParams struct {
	Auth string `json:"auth"`
}

type ViewOrderbookRequest struct {
	Id      int                   `json:"id"`
	JsonRPC string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  []ViewOrderbookParams `json:"params"`
}

type ViewOrderbookParams struct {
	ChainId int                `json:"chainId"`
	Query   ViewOrderbookQuery `json:"query"`
	Side    string             `json:"side"`
}

type ViewOrderbookQuery struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}

type MakeOrderRequest struct {
	Id      int               `json:"id"`
	JsonRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []MakeOrderParams `json:"params"`
}

type MakeOrderParams struct {
	Order    MakeOrderQuery `json:"order"`
	IsPublic bool           `json:"isPublic"`
	ChainId  int            `json:"chainId"`
}

type MakeOrderQuery struct {
	Signature  string          `json:"signature"`
	Parameters OrderParameters `json:"parameters"`
}

type CancelOrderRequest struct {
	Id      int                 `json:"id"`
	JsonRPC string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  []CancelOrderParams `json:"params"`
}

type CancelOrderParams struct {
	OrderId   string `json:"orderId"`
	Signature string `json:"signature"`
	ApiKey    string `json:"apiKey"`
}

type OrderParameters struct {
	Offerer                         string              `json:"offerer"`
	Zone                            string              `json:"zone"`
	Offer                           []OfferItem         `json:"offer"`
	Consideration                   []ConsiderationItem `json:"consideration"`
	OrderType                       OrderType           `json:"orderType"`
	StartTime                       string              `json:"startTime"`
	EndTime                         string              `json:"endTime"`
	ZoneHash                        string              `json:"zoneHash"`
	Salt                            string              `json:"salt"`
	ConduitKey                      string              `json:"conduitKey"`
	TotalOriginalConsiderationItems int16               `json:"totalOriginalConsiderationItems"`
	Counter                         string              `json:"counter"`
}
