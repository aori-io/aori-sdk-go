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
	Offerer                         string              `json:"offerer,omitempty"`
	Zone                            string              `json:"zone,omitempty"`
	Offer                           []OfferItem         `json:"offer,omitempty"`
	Consideration                   []ConsiderationItem `json:"consideration,omitempty"`
	OrderType                       OrderType           `json:"orderType,omitempty"`
	StartTime                       string              `json:"startTime,omitempty"`
	EndTime                         string              `json:"endTime,omitempty"`
	ZoneHash                        string              `json:"zoneHash,omitempty"`
	Salt                            string              `json:"salt,omitempty"`
	ConduitKey                      string              `json:"conduitKey,omitempty"`
	TotalOriginalConsiderationItems int16               `json:"totalOriginalConsiderationItems,omitempty"`
	Counter                         string              `json:"counter,omitempty"`
}
