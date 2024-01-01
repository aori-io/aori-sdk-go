package types

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
