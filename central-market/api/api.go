package api

import (
	"encoding/json"
	"log"
	"strings"

	centralmarket "github.com/kookehs/bdo-marketplace/central-market"
)

type BuyItemResponse struct {
	ResultResponse
}

type CalculateSellBiddingResponse struct {
	ResultResponse
}

type GetItemSellBuyInfoResponse struct {
	ResultResponse
	MarketConditionList       []centralmarket.MarketCondition `json:"marketConditionList"`
	PriceList                 []uint64                        `json:"priceList"`
	BasePrice                 uint64                          `json:"basePrice"`
	EnchantMaterialPrice      uint64                          `json:"enchantMaterialPrice"`
	BuyMaxCount               int                             `json:"buyMaxCount"`
	CountValue                int                             `json:"countValue"`
	EnchantGroup              int                             `json:"enchantGroup"`
	EnchantGroupMax           int                             `json:"enchantGroupMax"`
	EnchantMaterialKey        int                             `json:"enchantMaterialKey"`
	EnchantNeedCount          int                             `json:"enchantNeedCount"`
	MaxRegisterForWorldMarket int                             `json:"maxRegisterForWorldMarket"`
	SellMaxCount              int                             `json:"sellMaxCount"`
}

func (gisbir *GetItemSellBuyInfoResponse) PriceHistory() []centralmarket.PriceHistory {
	history := []centralmarket.PriceHistory{}
	data := strings.ReplaceAll(gisbir.ResultMsg, "\\", "")

	if err := json.Unmarshal([]byte(data), &history); err != nil {
		log.Println(err)
	}

	return history
}

type GetMyBiddingListResponse struct {
	ResultResponse
	BuyList  []centralmarket.BuyListing  `json:"buyList"`
	SellList []centralmarket.SellListing `json:"sellList"`
}

type GetMyWalletList struct {
	ResultResponse
	MyWalletList     []centralmarket.WalletListing `json:"myWalletList"`
	FeeRate          uint64                        `json:"feeRate"`
	RingBuffCount    int                           `json:"ringBuffCount"`
	AddWeight        int                           `json:"addWeight"`
	MaxWeight        int                           `json:"maxWeight"`
	TotalWeight      int                           `json:"totalWeight"`
	UseAddWeightBuff bool                          `json:"useAddWeightBuff"`
	UseValuePackage  bool                          `json:"useValuePackage"`
}

type GetWorldMarketHotListResponse struct {
	ResultResponse
	HotList []centralmarket.HotListing `json:"hotList"`
}

type GetWorldMarketListResponse struct {
	ResultResponse
	MarketList []centralmarket.MarketListing `json:"marketList"`
}

type GetWorldMarketSearchListResponse struct {
	ResultResponse
	List []centralmarket.SearchListing `json:"list"`
}

type GetWorldMarketSubListResponse struct {
	ResultResponse
	DetailList []centralmarket.DetailList `json:"detailList"`
}

type ResultResponse struct {
	ResultMsg  string `json:"resultMsg"`
	ResultCode int    `json:"resultCode"`
}

type SellItemResponse struct {
	ResultResponse
}

type WithdrawBuyBiddingResponse struct {
	ResultResponse
}
