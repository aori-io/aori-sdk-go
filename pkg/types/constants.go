package types

import "github.com/ethereum/go-ethereum/signer/core/apitypes"

const (
	DefaultOrderAddress   = "0xCB93Ed64a1b4C61b809F556532Ff36EA25DaD473"
	DefaultConduitKey     = "0x0000000000000000000000000000000000000000000000000000000000000000"
	RequestURL            = "wss://dev.api.aori.io/"
	MarketFeedURL         = "wss://dev.feed.aori.io"
	DefaultZoneHash       = "0x0000000000000000000000000000000000000000000000000000000000000000"
	SeaportAddress        = "0x00000000000000adc04c56bf30ac9d3c0aaf14dc"
	CurrentSeaportVersion = "1.5"
	ZeroAddress           = "0x0000000000000000000000000000000000000000"

	//DEFAULT_DURATION        = 86400000
)

var Eip712OrderType = apitypes.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"OrderComponents": {
		{
			Name: "offerer",
			Type: "address",
		},
		{
			Name: "zone",
			Type: "address",
		},
		{
			Name: "offer",
			Type: "OfferItem[]",
		},
		{
			Name: "consideration",
			Type: "ConsiderationItem[]",
		},
		{
			Name: "orderType",
			Type: "uint8",
		},
		{
			Name: "startTime",
			Type: "uint256",
		},
		{
			Name: "endTime",
			Type: "uint256",
		},
		{
			Name: "zoneHash",
			Type: "bytes32",
		},
		{
			Name: "salt",
			Type: "uint256",
		},
		{
			Name: "conduitKey",
			Type: "bytes32",
		},
		{
			Name: "counter",
			Type: "uint256",
		},
	},
	"OfferItem": {
		{
			Name: "itemType",
			Type: "uint8",
		},
		{
			Name: "token",
			Type: "address",
		},
		{
			Name: "identifierOrCriteria",
			Type: "uint256",
		},
		{
			Name: "startAmount",
			Type: "uint256",
		},
		{
			Name: "endAmount",
			Type: "uint256",
		},
	},
	"ConsiderationItem": {
		{
			Name: "itemType",
			Type: "uint8",
		},
		{
			Name: "token",
			Type: "address",
		},
		{
			Name: "identifierOrCriteria",
			Type: "uint256",
		},
		{
			Name: "startAmount",
			Type: "uint256",
		},
		{
			Name: "endAmount",
			Type: "uint256",
		},
		{
			Name: "recipient",
			Type: "address",
		},
	},
}
