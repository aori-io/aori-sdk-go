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
	InputAmount   uint64            `json:"inputAmount"`
	OutputAmount  uint64            `json:"outputAmount"`
	ChainID       int64             `json:"chainId"`
	Active        bool              `json:"active"`
	CreatedAt     uint64            `json:"createdAt"`
	LastUpdatedAt uint64            `json:"lastUpdatedAt"`
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
