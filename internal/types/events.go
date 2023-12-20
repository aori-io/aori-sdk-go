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

func NewOrderParameters(wallet string) OrderParameters {
	return OrderParameters{
		Offerer: wallet,
		Zone:    DefaultOrderAddress,
		Offer: []OfferItem{
			NewOfferItem(1, "", "0", "", ""),
		},
		Consideration: []ConsiderationItem{
			NewConsiderationItem(1, "", "0", "", "", wallet),
		},
		OrderType:                       3,
		StartTime:                       "",
		EndTime:                         "",
		ZoneHash:                        "0x0000000000000000000000000000000000000000000000000000000000000000",
		Salt:                            "0",
		ConduitKey:                      "0x0000000000000000000000000000000000000000000000000000000000000000",
		TotalOriginalConsiderationItems: 1,
	}
}

// DEPRECATE:

func NewOfferItem(itemType uint8, token, identifierOrCriteria, startAmount, endAmount string) OfferItem {
	return OfferItem{
		ItemType:             itemType,
		Token:                token,
		IdentifierOrCriteria: identifierOrCriteria,
		StartAmount:          startAmount,
		EndAmount:            endAmount,
	}
}

func NewConsiderationItem(itemType uint8, token, identifierOrCriteria, startAmount, endAmount, recipient string) ConsiderationItem {
	return ConsiderationItem{
		ItemType:             itemType,
		Token:                token,
		IdentifierOrCriteria: identifierOrCriteria,
		StartAmount:          startAmount,
		EndAmount:            endAmount,
		Recipient:            recipient,
	}
}
