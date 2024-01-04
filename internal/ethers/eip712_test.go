package ethers

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg/types"
	"os"
	"testing"
)

func TestSignOrder(t *testing.T) {
	// set random private key
	err := os.Setenv("PRIVATE_KEY", "7e5ba2c97e9bb3cefe15e1bbbecf677a0ca99765a2ac47a03a08ca722cc32baa")
	if err != nil {
		t.Fatalf("error setting env var: %v", err)
	}

	type args struct {
		order   types.OrderParameters
		chainId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "No Offer or Consideration",
			args: args{
				order: types.OrderParameters{
					Offerer:                         "0xb794f5ea0ba39494ce839613fffba74279579268", // random address
					Zone:                            types.DefaultOrderAddress,
					Offer:                           nil,
					Consideration:                   nil,
					OrderType:                       types.PartialRestricted,
					StartTime:                       fmt.Sprintf("%v", 1234567989),
					EndTime:                         fmt.Sprintf("%v", 1234567989), // 24 hours later
					ZoneHash:                        types.DefaultZoneHash,
					Salt:                            "0",
					ConduitKey:                      types.DefaultConduitKey,
					TotalOriginalConsiderationItems: 1,
					Counter:                         "0",
				},
				chainId: 1,
			},
			wantErr: false,
		},
		{
			name: "One Offer and Consideration",
			args: args{
				order: types.OrderParameters{
					Offerer: "0xb794f5ea0ba39494ce839613fffba74279579268", // random address
					Zone:    types.DefaultOrderAddress,
					Offer: []types.OfferItem{{
						ItemType:             types.ERC20,
						Token:                "0xb794f5ea0ba39494ce839613fffba74279579268",
						IdentifierOrCriteria: "0",
						StartAmount:          "1",
						EndAmount:            "1",
					}},
					Consideration: []types.ConsiderationItem{{
						ItemType:             types.ERC20,
						Token:                "0xb794f5ea0ba39494ce839613fffba74279579268",
						IdentifierOrCriteria: "0",
						StartAmount:          "1",
						EndAmount:            "1",
						Recipient:            "0xb794f5ea0ba39494ce839613fffba74279579268",
					}},
					OrderType:                       types.PartialRestricted,
					StartTime:                       fmt.Sprintf("%v", 1234567989),
					EndTime:                         fmt.Sprintf("%v", 1234567989), // 24 hours later
					ZoneHash:                        types.DefaultZoneHash,
					Salt:                            "0",
					ConduitKey:                      types.DefaultConduitKey,
					TotalOriginalConsiderationItems: 1,
					Counter:                         "0",
				},
				chainId: 1,
			},
			wantErr: false,
		},
		{
			name: "Multiple Offer and Consideration",
			args: args{
				order: types.OrderParameters{
					Offerer: "0xb794f5ea0ba39494ce839613fffba74279579268", // random address
					Zone:    types.DefaultOrderAddress,
					Offer: []types.OfferItem{{
						ItemType:             types.ERC20,
						Token:                "0xb794f5ea0ba39494ce839613fffba74279579268",
						IdentifierOrCriteria: "0",
						StartAmount:          "1",
						EndAmount:            "1",
					},
						{
							ItemType:             types.ERC20,
							Token:                "0xb794f5ea0ba39494ce839613fffba74279579268",
							IdentifierOrCriteria: "0",
							StartAmount:          "1",
							EndAmount:            "1",
						}},
					Consideration: []types.ConsiderationItem{{
						ItemType:             types.ERC20,
						Token:                "0xb794f5ea0ba39494ce839613fffba74279579268",
						IdentifierOrCriteria: "0",
						StartAmount:          "1",
						EndAmount:            "1",
						Recipient:            "0xb794f5ea0ba39494ce839613fffba74279579268",
					},
						{
							ItemType:             types.ERC20,
							Token:                "0xb794f5ea0ba39494ce839613fffba74279579268",
							IdentifierOrCriteria: "0",
							StartAmount:          "1",
							EndAmount:            "1",
							Recipient:            "0xb794f5ea0ba39494ce839613fffba74279579268",
						}},
					OrderType:                       types.PartialRestricted,
					StartTime:                       fmt.Sprintf("%v", 1234567989),
					EndTime:                         fmt.Sprintf("%v", 1234567989), // 24 hours later
					ZoneHash:                        types.DefaultZoneHash,
					Salt:                            "0",
					ConduitKey:                      types.DefaultConduitKey,
					TotalOriginalConsiderationItems: 1,
					Counter:                         "0",
				},
				chainId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SignOrder(tt.args.order, tt.args.chainId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
