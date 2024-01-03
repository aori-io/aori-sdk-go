package types

type OrderType int

const (
	FullOpen OrderType = iota
	PartialOpen
	FullRestricted
	PartialRestricted
	Contract
)

type ItemType int

const (
	Native ItemType = iota
	ERC20
	ERC721
	ERC1155
	ERC721WithCriteria
	ERC1155WithCriteria
)

type OfferItem struct {
	ItemType             ItemType `json:"itemType"`
	Token                string   `json:"token"`
	IdentifierOrCriteria string   `json:"identifierOrCriteria"`
	StartAmount          string   `json:"startAmount"`
	EndAmount            string   `json:"endAmount"`
}

type ConsiderationItem struct {
	ItemType             ItemType `json:"itemType"`
	Token                string   `json:"token"`
	IdentifierOrCriteria string   `json:"identifierOrCriteria"`
	StartAmount          string   `json:"startAmount"`
	EndAmount            string   `json:"endAmount"`
	Recipient            string   `json:"recipient"`
}
