package types

import "github.com/ethereum/go-ethereum/signer/core/apitypes"

const (
	DefaultOrderAddress   = "0xEF3137050f3a49ECAe2D2Bae0154B895310D9Dc4"
	DefaultConduitKey     = "0x0000000000000000000000000000000000000000000000000000000000000000"
	RequestURL            = "ws://localhost:8080/"
	DefaultZoneHash       = "0x0000000000000000000000000000000000000000000000000000000000000000"
	SeaportAddress        = "0x00000000000000adc04c56bf30ac9d3c0aaf14dc"
	CurrentSeaportVersion = "1.5"

	//DEFAULT_DURATION        = 86400000
	//MARKET_FEED_URL         = "wss://dev.beta.feed.aori.io/"
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
			Name: "totalOriginalConsiderationItems",
			Type: "uint256",
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
