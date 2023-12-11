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

type OrderParameters struct {
	Offerer                         string              `json:"offerer"`
	Zone                            string              `json:"zone"`
	Offer                           []OfferItem         `json:"offer"`
	Consideration                   []ConsiderationItem `json:"consideration"`
	OrderType                       uint8               `json:"orderType"`
	StartTime                       string              `json:"startTime"`
	EndTime                         string              `json:"endTime"`
	ZoneHash                        string              `json:"zoneHash"`
	Salt                            string              `json:"salt"`
	ConduitKey                      string              `json:"conduitKey"`
	TotalOriginalConsiderationItems int16               `json:"totalOriginalConsiderationItems"`
}

type OfferItem struct {
	ItemType             uint8  `json:"itemType"`
	Token                string `json:"token"`
	IdentifierOrCriteria string `json:"identifierOrCriteria"`
	StartAmount          string `json:"startAmount"`
	EndAmount            string `json:"endAmount"`
}

type ConsiderationItem struct {
	ItemType             uint8  `json:"itemType"`
	Token                string `json:"token"`
	IdentifierOrCriteria string `json:"identifierOrCriteria"`
	StartAmount          string `json:"startAmount"`
	EndAmount            string `json:"endAmount"`
	Recipient            string `json:"recipient"`
}

type OrderComponents struct {
	Offerer       string
	Zone          string
	Offer         []OfferItem
	Consideration []ConsiderationItem
	OrderType     OrderType
	StartTime     string
	EndTime       string
	ZoneHash      string
	Salt          string
	ConduitKey    string
	Counter       string
}
