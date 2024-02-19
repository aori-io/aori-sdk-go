package util

import (
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/internal/ethers"
	"github.com/aori-io/aori-sdk-go/pkg/types"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
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
			Order: types.OrderParameters{
				Offerer:       walletAddress,
				InputToken:    orderParams.SellToken,
				InputAmount:   orderParams.SellAmount,
				InputChainID:  uint64(chainId),
				InputZone:     types.DefaultOrderAddress,
				OutputToken:   orderParams.BuyToken,
				OutputAmount:  orderParams.BuyAmount,
				OutputChainID: uint64(chainId),
				OutputZone:    types.DefaultOrderAddress,
				StartTime:     fmt.Sprintf("%v", startTime.Unix()),
				EndTime:       fmt.Sprintf("%v", endTime.Unix()),
				Salt:          "0",
				Counter:       1,
				ToWithdraw:    false,
			},
			IsPublic: isPublic,
			APIKey:   "",
		}},
	}

	hash, err := ethers.CalculateOrderHash(req.Params[0].Order)
	if err != nil {
		panic(err)
	}

	sig, err := ethers.SignOrderHash(hash)
	if err != nil {
		return nil, fmt.Errorf("make_order error signing order: %s", err)
	}
	req.Params[0].Signature = sig

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("make_order error marshalling order: %s", err)
	}
	return b, nil
}

func CreateTakeOrderPayload(id, chainId, seatId int, walletAddress, orderId string, orderParams types.OrderParameters) ([]byte, error) {
	startTime := time.Now()
	endTime := startTime.AddDate(0, 0, 1)

	req := types.TakeOrderRequest{
		Id:      id,
		JsonRPC: "2.0",
		Method:  "aori_takeOrder",
		Params: []types.TakeOrderParams{{
			Order: types.TakeOrderQuery{
				Signature: "",
				Parameters: types.OrderParameters{
					Offerer:       "",
					InputToken:    "",
					InputAmount:   "",
					InputChainID:  0,
					InputZone:     "",
					OutputToken:   "",
					OutputAmount:  "",
					OutputChainID: 0,
					OutputZone:    "",
					StartTime:     fmt.Sprintf("%v", startTime.Unix()),
					EndTime:       fmt.Sprintf("%v", endTime.Unix()),
					Salt:          "",
					Counter:       0,
					ToWithdraw:    false,
				},
			},
			OrderId: orderId,
			ChainId: chainId,
			SeatId:  seatId,
		}},
	}

	//sig, err := ethers.SignOrder(req.Params[0].Order.Parameters, chainId)
	//if err != nil {
	//	return nil, fmt.Errorf("error signing order: %s", err)
	//}

	//req.Params[0].Order.Signature = sig

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
