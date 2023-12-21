package util

import (
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/internal/types"
	"time"
)

func CreatePingPayload(id int) ([]byte, error) {
	req := types.PingRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_ping",
		Params:  []string{},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateAuthWalletPayload(id int, address, signature string) ([]byte, error) {
	req := types.AuthWalletRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_authWallet",
		Params:  []types.AuthWalletParams{{Address: address, Signature: signature}},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateCheckAuthPayload(id int, jwt string) ([]byte, error) {
	req := types.CheckAuthRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_checkAuth",
		Params:  []types.CheckAuthParams{{Auth: jwt}},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateViewOrderbookPayload(id int, chainId int, base, quote, side string) ([]byte, error) {
	req := types.ViewOrderbookRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_viewOrderbook",
		Params: []types.ViewOrderbookParams{{ChainId: chainId,
			Query: types.ViewOrderbookQuery{
				Base:  base,
				Quote: quote,
			}, Side: side,
		}},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateMakeOrderPayload(id, chainId int, walletAddress, sellToken, sellAmount, buyToken, buyAmount string) ([]byte, error) {
	startTime := time.Now()
	endTime := startTime.AddDate(0, 0, 1)

	req := types.MakeOrderRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_makeOrder",
		Params: []types.MakeOrderParams{{
			Order: types.MakeOrderQuery{
				Signature: "",
				Parameters: types.OrderComponents{
					Offerer: walletAddress,
					Zone:    types.DefaultOrderAddress,
					Offer: []types.OfferItem{{
						ItemType:             1,
						Token:                sellToken,
						IdentifierOrCriteria: "0",
						StartAmount:          sellAmount,
						EndAmount:            sellAmount,
					}},
					Consideration: []types.ConsiderationItem{{
						ItemType:             1,
						Token:                buyToken,
						IdentifierOrCriteria: "0",
						StartAmount:          buyAmount,
						EndAmount:            buyAmount,
						Recipient:            walletAddress,
					}},
					OrderType:                       3,
					StartTime:                       fmt.Sprintf("%v", startTime.Unix()),
					EndTime:                         fmt.Sprintf("%v", endTime.Unix()), // 24 hours later
					ZoneHash:                        types.DefaultZoneHash,
					Salt:                            "0",
					ConduitKey:                      types.DefaultConduitKey,
					TotalOriginalConsiderationItems: 1,
					Counter:                         "0",
				},
			},
			IsPublic: true,
			ChainId:  chainId,
		}},
	}

	sig, err := SignOrder(req.Params[0].Order.Parameters)
	if err != nil {
		return nil, fmt.Errorf("make_order error signing order: %s", err)
	}

	fmt.Println("SIGNATURE: ", sig)

	req.Params[0].Order.Signature = sig

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("make_order error marshalling order: %s", err)
	}
	return b, nil
}