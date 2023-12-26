package util

import (
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg/types"
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

func CreateViewOrderbookPayload(id int, query types.ViewOrderbookParams) ([]byte, error) {
	req := types.ViewOrderbookRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_viewOrderbook",
		Params:  []types.ViewOrderbookParams{query},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateMakeOrderPayload(id, chainId int, walletAddress, sellToken, sellAmount, buyToken, buyAmount string, isPublic bool) ([]byte, error) {
	startTime := time.Now()
	endTime := startTime.AddDate(0, 0, 1)
	req := types.MakeOrderRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_makeOrder",
		Params: []types.MakeOrderParams{{
			Order: types.MakeOrderQuery{
				Signature: "",
				Parameters: types.OrderParameters{
					Offerer: walletAddress,
					Zone:    types.DefaultOrderAddress,
					Offer: []types.OfferItem{{
						ItemType:             types.ERC20,
						Token:                sellToken,
						IdentifierOrCriteria: "0",
						StartAmount:          sellAmount,
						EndAmount:            sellAmount,
					}},
					Consideration: []types.ConsiderationItem{{
						ItemType:             types.ERC20,
						Token:                buyToken,
						IdentifierOrCriteria: "0",
						StartAmount:          buyAmount,
						EndAmount:            buyAmount,
						Recipient:            walletAddress,
					}},
					OrderType:                       types.PartialRestricted,
					StartTime:                       fmt.Sprintf("%v", startTime.Unix()),
					EndTime:                         fmt.Sprintf("%v", endTime.Unix()), // 24 hours later
					ZoneHash:                        types.DefaultZoneHash,
					Salt:                            "0",
					ConduitKey:                      types.DefaultConduitKey,
					TotalOriginalConsiderationItems: 1,
					Counter:                         "0",
				},
			},
			IsPublic: isPublic,
			ChainId:  chainId,
		}},
	}

	sig, err := SignOrder(req.Params[0].Order.Parameters, chainId)
	if err != nil {
		return nil, fmt.Errorf("make_order error signing order: %s", err)
	}

	req.Params[0].Order.Signature = sig

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("make_order error marshalling order: %s", err)
	}
	return b, nil
}

func CreateCancelOrderPayload(id int, orderId, apiKey string) ([]byte, error) {
	sig, err := SignCancelOrder(orderId)
	if err != nil {
		return nil, err
	}

	req := types.CancelOrderRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_cancelOrder",
		Params: []types.CancelOrderParams{{
			OrderId:   orderId,
			Signature: sig,
			ApiKey:    apiKey,
		}},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("cancel_order error marshalling order: %s", err)
	}

	fmt.Println(string(b))
	return b, nil
}

func CreateAccountOrdersPayload(id int, wallet, signature, apiKey string) ([]byte, error) {
	req := types.AccountOrdersRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_accountOrders",
		Params: []types.AccountOrdersParams{{
			Offerer:   wallet,
			Signature: signature,
			ApiKey:    apiKey,
		}},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("account_orders error marshalling order: %s", err)
	}

	return b, nil
}

func CreateOrderStatusPayload(id int, orderHash string) ([]byte, error) {
	req := types.OrderStatusRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_orderStatus",
		Params:  []types.OrderStatusParams{{OrderHash: orderHash}},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("order_status error marshalling order: %s", err)
	}

	return b, nil
}
