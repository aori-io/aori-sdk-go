package types

type SubscriptionEvent string

const (
	OrderCreated   SubscriptionEvent = "OrderCreated"
	OrderCancelled SubscriptionEvent = "OrderCancelled"
	OrderTaken     SubscriptionEvent = "OrderTaken"
	OrderFulfilled SubscriptionEvent = "OrderFulfilled"
	OrderToExecute SubscriptionEvent = "OrderToExecute"
	QuoteRequested SubscriptionEvent = "QuoteRequested"
)

type MakeOrderInput struct {
	SellToken  string
	SellAmount string
	BuyToken   string
	BuyAmount  string
}

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
	ChainId      int                `json:"chainId,omitempty"`
	Query        ViewOrderbookQuery `json:"query,omitempty"`
	Side         string             `json:"side,omitempty"`
	SortBy       string             `json:"sortBy,omitempty"`
	InputAmount  string             `json:"inputAmount,omitempty"`
	OutputAmount string             `json:"outputAmount,omitempty"`
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
	Order     MakeOrderQuery `json:"order"`
	Signer    string         `json:"signer"`
	Signature string         `json:"signature"`
	IsPublic  bool           `json:"isPublic"`
	ChainId   int            `json:"chainId"`
	APIKey    string         `json:"apiKey"`
}

type MakeOrderQuery struct {
	Parameters OrderParameters `json:"parameters"`
}

type TakeOrderRequest struct {
	Id      int               `json:"id"`
	JsonRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []TakeOrderParams `json:"params"`
}

type TakeOrderParams struct {
	Order   TakeOrderQuery `json:"order"`
	OrderId string         `json:"orderId"`
	SeatId  int            `json:"seatId"`
	ChainId int            `json:"chainId"`
}

type TakeOrderQuery struct {
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

type AccountOrdersRequest struct {
	Id      int                   `json:"id"`
	JsonRPC string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  []AccountOrdersParams `json:"params"`
}

type AccountOrdersParams struct {
	Offerer   string `json:"offerer"`
	Signature string `json:"signature,omitempty"`
	ApiKey    string `json:"apiKey,omitempty"`
}

type OrderStatusRequest struct {
	Id      int                 `json:"id"`
	JsonRPC string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  []OrderStatusParams `json:"params"`
}

type OrderStatusParams struct {
	OrderHash string `json:"orderHash"`
}

type CancelAllOrdersRequest struct {
	Id      int                     `json:"id"`
	JsonRPC string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  []CancelAllOrdersParams `json:"params"`
}

type CancelAllOrdersParams struct {
	Offerer   string `json:"offerer"`
	Signature string `json:"signature,omitempty"`
}

type SubscribeOrderbookRequest struct {
	Id      int      `json:"id"`
	JsonRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type OrderParameters struct {
	Offerer       string `json:"offerer"`
	InputToken    string `json:"inputToken"`
	InputAmount   string `json:"inputAmount"`
	InputChainID  uint64 `json:"inputChainId"`
	InputZone     string `json:"inputZone"`
	OutputToken   string `json:"outputToken"`
	OutputAmount  string `json:"outputAmount"`
	OutputChainID uint64 `json:"outputChainId"`
	OutputZone    string `json:"outputZone"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	Salt          string `json:"salt"`
	Counter       uint64 `json:"counter"`
	ToWithdraw    bool   `json:"toWithdraw"`
}

type SubscribeQuoteRequestedEvent struct {
	InputToken   string `json:"inputToken"`
	OutputToken  string `json:"outputToken"`
	InputAmount  string `json:"inputAmount,omitempty"`
	OutputAmount string `json:"outputAmount,omitempty"`
	ChainID      int    `json:"chainId"`
}

type SubscribeOrderToExecuteEvent struct {
	// Relevant order details
	MakerOrderHash  string          `json:"makerOrderHash"`
	MakerParameters OrderParameters `json:"makerParameters"`
	TakerOrderHash  string          `json:"takerOrderHash"`
	TakerParameters OrderParameters `json:"takerParameters"`
	MatchingHash    string          `json:"matchingHash"`

	// Verification
	ChainID       int    `json:"chainId"`
	To            string `json:"to"`
	Value         int    `json:"value"`
	Data          string `json:"data"`
	BlockDeadline int    `json:"blockDeadline"`

	// Vanity
	Maker        string `json:"maker"`
	InputToken   string `json:"inputToken"`
	InputAmount  string `json:"inputAmount"`
	OutputToken  string `json:"outputToken"`
	OutputAmount string `json:"outputAmount"`
}
