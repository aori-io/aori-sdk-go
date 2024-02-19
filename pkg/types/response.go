package types

// Aori Request Response

type AoriOrderStatusResponse struct {
	ID     int                 `json:"id"`
	Result AoriOrderStatusData `json:"result"`
}

type AoriOrderStatusData struct {
	Order OrderCreatedData `json:"order"`
}

type AoriViewOrderbookResponse struct {
	ID     int               `json:"id"`
	Result AoriOrderbookData `json:"result"`
}

type AoriOrderbookData struct {
	Orders []OrderCreatedData `json:"orders"`
}

type AuthWalletResponse struct {
	ID     int            `json:"id"`
	Result AuthWalletData `json:"result"`
}

type AuthWalletData struct {
	Auth string `json:"auth"`
}

type OrderStatusResponse struct {
	ID     int               `json:"id"`
	Result OrderStatusResult `json:"result"`
}

type OrderStatusResult struct {
	Order OrderCreatedData `json:"order"`
}

type AccountOrdersResponse struct {
	ID     int                 `json:"id"`
	Result AccountOrdersResult `json:"result"`
}

type AccountOrdersResult struct {
	Orders []OrderCreatedData `json:"orders"`
}

type MakeOrderResponse struct {
	ID     int             `json:"id"`
	Result MakeOrderResult `json:"result"`
}

type MakeOrderResult struct {
	CreatedAt     int64           `json:"createdAt"`
	InputAmount   string          `json:"inputAmount"`
	InputChainId  int             `json:"inputChainId"`
	InputToken    string          `json:"inputToken"`
	InputZone     string          `json:"inputZone"`
	IsActive      bool            `json:"isActive"`
	IsPublic      bool            `json:"isPublic"`
	LastUpdatedAt int64           `json:"lastUpdatedAt"`
	Order         OrderParameters `json:"order"`
	OrderHash     string          `json:"orderHash"`
	OutputAmount  string          `json:"outputAmount"`
	OutputChainId int             `json:"outputChainId"`
	OutputToken   string          `json:"outputToken"`
	OutputZone    string          `json:"outputZone"`
	Rate          int             `json:"rate"`
	Signature     string          `json:"signature"`
}

// Aori Feed Subscribe Orderbook Responses

// SubscribeOrderViewResponse For OrderCreated and OrderCancelled
type SubscribeOrderViewResponse struct {
	ID     int                      `json:"id"`
	Result SubscribeOrderViewResult `json:"result"`
}

type SubscribeOrderViewResult struct {
	Type string           `json:"type"`
	Data OrderCreatedData `json:"data"`
}

// SubscribeTakeOrderResponse For OrderTaken
type SubscribeTakeOrderResponse struct {
	ID     int                      `json:"id"`
	Result SubscribeTakeOrderResult `json:"result"`
}

type SubscribeTakeOrderResult struct {
	Type string           `json:"type"`
	Data OrderCreatedData `json:"data"`
}

// SubscribeFulfilledOrderResponse For OrderFulfilled
type SubscribeFulfilledOrderResponse struct {
	ID     int                           `json:"id"`
	Result SubscribeFulfilledOrderResult `json:"result"`
}

type SubscribeFulfilledOrderResult struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// SubscribeQuoteRequestedResponse For QuoteRequested
type SubscribeQuoteRequestedResponse struct {
	ID     int                           `json:"id"`
	Result SubscribeQuoteRequestedResult `json:"result"`
}

type SubscribeQuoteRequestedResult struct {
	Type string                       `json:"type"`
	Data SubscribeQuoteRequestedEvent `json:"data"`
}

// SubscribeOrderToExecuteResponse For OrderToExecute
type SubscribeOrderToExecuteResponse struct {
	ID     int                           `json:"id"`
	Result SubscribeOrderToExecuteResult `json:"result"`
}

type SubscribeOrderToExecuteResult struct {
	Type string                       `json:"type"`
	Data SubscribeOrderToExecuteEvent `json:"data"`
}
