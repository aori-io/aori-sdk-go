package internal

import (
	"encoding/json"
	"github.com/aori-io/aori-sdk-go/internal/types"
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

func CreateMakeOrderPayload(id int) ([]byte, error) {
	req := types.MakeOrderRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_makeOrder",
		Params: []types.MakeOrderParams{{
			Order: types.MakeOrderQuery{
				Signature: "",
				Parameters: types.OrderParameters{
					Offerer:                         "",
					Zone:                            "",
					Offer:                           nil,
					Consideration:                   nil,
					OrderType:                       0,
					StartTime:                       "",
					EndTime:                         "",
					ZoneHash:                        "",
					Salt:                            "",
					ConduitKey:                      "",
					TotalOriginalConsiderationItems: 0,
				},
			},
			IsPublic: true,
			ChainId:  1,
		}},
	}
	b, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	return b, nil
}
