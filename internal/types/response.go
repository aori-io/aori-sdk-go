package types

type AoriOrderStatusResponse struct {
	ID     *uint64             `json:"id"`
	Result AoriOrderStatusData `json:"result"`
}

type AoriOrderStatusData struct {
	Order OrderCreatedData `json:"order"`
}

type AoriViewOrderbookResponse struct {
	ID     *uint64           `json:"id"`
	Result AoriOrderbookData `json:"result"`
}

type AoriOrderbookData struct {
	Orders []OrderCreatedData `json:"orders"`
}
