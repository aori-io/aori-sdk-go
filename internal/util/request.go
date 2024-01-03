package util

import (
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/internal/ethers"
	"github.com/aori-io/aori-sdk-go/pkg/types"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
	"strconv"
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

func CreateMakeOrderPayload(id, chainId int, walletAddress string, orderParams types.MakeOrderInput, isPublic bool) ([]byte, error) {
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
						Token:                orderParams.SellToken,
						IdentifierOrCriteria: "0",
						StartAmount:          orderParams.SellAmount,
						EndAmount:            orderParams.SellAmount,
					}},
					Consideration: []types.ConsiderationItem{{
						ItemType:             types.ERC20,
						Token:                orderParams.BuyToken,
						IdentifierOrCriteria: "0",
						StartAmount:          orderParams.BuyAmount,
						EndAmount:            orderParams.BuyAmount,
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

	sig, err := ethers.SignOrder(req.Params[0].Order.Parameters, chainId)
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

func CreateTakeOrderPayload(id, chainId, seatId int, walletAddress, orderId string, orderParams types.OrderParameters) ([]byte, error) {
	startTime := time.Now()
	endTime := startTime.AddDate(0, 0, 1)

	// swap offer and consideration
	for i := range orderParams.Consideration {
		orderParams.Offer[i].ItemType, orderParams.Consideration[i].ItemType = orderParams.Consideration[i].ItemType, orderParams.Offer[i].ItemType
		orderParams.Offer[i].StartAmount, orderParams.Consideration[i].StartAmount = orderParams.Consideration[i].StartAmount, orderParams.Offer[i].StartAmount
		orderParams.Offer[i].EndAmount, orderParams.Consideration[i].EndAmount = orderParams.Consideration[i].EndAmount, orderParams.Offer[i].EndAmount
		orderParams.Offer[i].Token, orderParams.Consideration[i].Token = orderParams.Consideration[i].Token, orderParams.Offer[i].Token
		orderParams.Offer[i].IdentifierOrCriteria, orderParams.Consideration[i].IdentifierOrCriteria = orderParams.Consideration[i].IdentifierOrCriteria, orderParams.Offer[i].IdentifierOrCriteria
		orderParams.Consideration[i].Recipient = walletAddress
	}

	// add fees
	for _, offer := range orderParams.Offer {
		startAmount, err := strconv.ParseFloat(offer.StartAmount, 64)
		if err != nil {
			return nil, err
		}
		endAmount, err := strconv.ParseFloat(offer.EndAmount, 64)
		if err != nil {
			return nil, err
		}

		offer.StartAmount = fmt.Sprintf("%f", startAmount*1.0003)
		offer.EndAmount = fmt.Sprintf("%f", endAmount*1.0003)
	}

	req := types.TakeOrderRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_takeOrder",
		Params: []types.TakeOrderParams{{
			Order: types.TakeOrderQuery{
				Signature: "",
				Parameters: types.OrderParameters{
					Offerer:                         walletAddress,
					Zone:                            types.DefaultOrderAddress,
					Offer:                           orderParams.Offer,
					Consideration:                   orderParams.Consideration,
					OrderType:                       orderParams.OrderType,
					StartTime:                       fmt.Sprintf("%v", startTime.Unix()),
					EndTime:                         fmt.Sprintf("%v", endTime.Unix()), // 24 hours later
					ZoneHash:                        types.DefaultZoneHash,
					Salt:                            "0",
					ConduitKey:                      types.DefaultConduitKey,
					TotalOriginalConsiderationItems: 1,
					Counter:                         "0",
				},
			},
			OrderId: orderId,
			ChainId: chainId,
			SeatId:  seatId,
		}},
	}

	sig, err := ethers.SignOrder(req.Params[0].Order.Parameters, chainId)
	if err != nil {
		return nil, fmt.Errorf("error signing order: %s", err)
	}

	req.Params[0].Order.Signature = sig

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling order: %s", err)
	}

	return b, nil
}

func CreateCancelOrderPayload(id int, orderId, apiKey string) ([]byte, error) {
	sig, err := ethers.SignCancelOrder(orderId)
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

func CreateCancelAllOrdersPayload(id int, wallet string) ([]byte, error) {
	key := os.Getenv("PRIVATE_KEY")
	if key == "" {
		return nil, fmt.Errorf("missing PRIVATE_KEY")
	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}

	signature, err := ethers.PersonalSign(wallet, privateKey)
	if err != nil {
		return nil, err
	}

	req := types.CancelAllOrdersRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_cancelAllOrders",
		Params: []types.CancelAllOrdersParams{{
			Offerer:   wallet,
			Signature: signature,
		}},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("order_status error marshalling order: %s", err)
	}

	return b, nil
}

func CreateSubscribeOrderbookPayload(id int) ([]byte, error) {
	req := types.SubscribeOrderbookRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_subscribeOrderbook",
		Params:  []string{},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("order_status error marshalling order: %s", err)
	}

	return b, nil
}
