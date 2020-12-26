package centralmarket

import (
	"encoding/json"
	"log"
	"strings"
)

const (
	// Form data keys.
	BuyChooseKeyString  = "buyChooseKey"
	BuyCountString      = "buyCount"
	BuyMainKeyString    = "buyMainKey"
	BuyNoString         = "buyNo"
	BuyPriceString      = "buyPrice"
	BuySubKeyString     = "buySubKey"
	ChooseKeyString     = "chooseKey"
	CountString         = "count"
	IsRingBuffString    = "isRingBuff"
	IsSealedString      = "isSealed"
	IsUpString          = "isUp"
	KeyTypeString       = "keyType"
	MainKeyString       = "mainKey"
	SearchTextString    = "searchText"
	SellChooseKeyString = "sellChooseKey"
	SellCountString     = "sellCount"
	SellKeyTypeString   = "sellKeyType"
	SellNoString        = "sellNo"
	SellPriceString     = "sellPrice"
	SellSubKeyString    = "sellSubKey"
	SubKeyString        = "subKey"
	TokenString         = "__RequestVerificationToken"

	// API Endpoints.
	BuyItemEndpoint                  = "/GameTradeMarket/BuyItem"
	CalculateSellBiddingEndpoint     = "/GameTradeMarket/CalculateSellBidding"
	GetItemSellBuyInfoEndpoint       = "/Home/GetItemSellBuyInfo"
	GetMyBiddingListEndpoint         = "/Home/GetMyBiddingList"
	GetMyWalletListEndpoint          = "/Home/GetMyWalletList"
	GetWorldMarketHotListEndpoint    = "/Home/GetWorldMarketHotList"
	GetWorldMarketListEndpoint       = "/Home/GetWorldMarketList"
	GetWorldMarketSearchListEndpoint = "/Home/GetWorldMarketSearchList"
	GetWorldMarketSubListEndpoint    = "/Home/GetWorldMarketSubList"
	SellItemEndpoint                 = "/GameTradeMarket/SellItem"
	WithdrawBuyBiddingEndpoint       = "/GameTradeMarket/WithdrawBuyBidding"
	WithdrawSellBiddingEndpoint      = "/GameTradeMarket/WithdrawSellBidding"
)

type BuyItemInput struct {
	BuyPrice     uint64 `json:"buyPrice"`
	BuyChooseKey int    `json:"buyChooseKey"`
	BuyCount     int    `json:"buyCount"`
	BuyMainKey   int    `json:"buyMainKey"`
	BuySubKey    int    `json:"buySubKey"`
}

type BuyItemOutput struct {
	ResultOutput
}

type CalculateSellBiddingInput struct {
	SellNo    uint64 `json:"sellNo"`
	ChooseKey int    `json:"chooseKey"`
	KeyType   int    `json:"keyType"`
	MainKey   int    `json:"mainKey"`
	SubKey    int    `json:"subKey"`
	IsSealed  bool   `json:"isSealed"`
}

type CalculateSellBiddingOutput struct {
	ResultOutput
}

type GetItemSellBuyInfoInput struct {
	Token   string `json:"__RequestVerificationToken"`
	KeyType int    `json:"keyType"`
	MainKey int    `json:"mainKey"`
	SubKey  int    `json:"subKey"`
	IsUp    bool   `json:"isUp"`
}

type GetItemSellBuyInfoOutput struct {
	ResultOutput
	MarketConditionList       []MarketCondition `json:"marketConditionList"`
	PriceList                 []uint64          `json:"priceList"`
	BasePrice                 uint64            `json:"basePrice"`
	EnchantMaterialPrice      uint64            `json:"enchantMaterialPrice"`
	BuyMaxCount               int               `json:"buyMaxCount"`
	CountValue                int               `json:"countValue"`
	EnchantGroup              int               `json:"enchantGroup"`
	EnchantGroupMax           int               `json:"enchantGroupMax"`
	EnchantMaterialKey        int               `json:"enchantMaterialKey"`
	EnchantNeedCount          int               `json:"enchantNeedCount"`
	MaxRegisterForWorldMarket int               `json:"maxRegisterForWorldMarket"`
	SellMaxCount              int               `json:"sellMaxCount"`
}

func (gisbio *GetItemSellBuyInfoOutput) PriceHistory() []PriceHistory {
	history := []PriceHistory{}
	data := strings.ReplaceAll(gisbio.ResultMsg, "\\", "")

	if err := json.Unmarshal([]byte(data), &history); err != nil {
		log.Println(err)
	}

	return history
}

type GetMyBiddingListOutput struct {
	ResultOutput
	BuyList  []BuyListing  `json:"buyList"`
	SellList []SellListing `json:"sellList"`
}

type GetMyWalletListOutput struct {
	ResultOutput
	MyWalletList     []WalletListing `json:"myWalletList"`
	FeeRate          uint64          `json:"feeRate"`
	RingBuffCount    int             `json:"ringBuffCount"`
	AddWeight        int             `json:"addWeight"`
	MaxWeight        int             `json:"maxWeight"`
	TotalWeight      int             `json:"totalWeight"`
	UseAddWeightBuff bool            `json:"useAddWeightBuff"`
	UseValuePackage  bool            `json:"useValuePackage"`
}

type GetWorldMarketHotListOutput struct {
	ResultOutput
	HotList []HotListing `json:"hotList"`
}

type GetWorldMarketListInput struct {
	Token   string `json:"__RequestVerificationToken"`
	MainKey int    `json:"mainKey"`
	SubKey  int    `json:"subKey"`
}

type GetWorldMarketListOutput struct {
	ResultOutput
	MarketList []MarketListing `json:"marketList"`
}

type GetWorldMarketSearchListInput struct {
	SearchText string `json:"searchText"`
	Token      string `json:"__RequestVerificationToken"`
}

type GetWorldMarketSearchListOutput struct {
	ResultOutput
	List []SearchListing `json:"list"`
}

type GetWorldMarketSubListInput struct {
	MainKey string `json:"mainKey"`
	Token   string `json:"__RequestVerificationToken"`
}

type GetWorldMarketSubListOutput struct {
	ResultOutput
	DetailList []DetailList `json:"detailList"`
}

type ResultOutput struct {
	ResultMsg  string `json:"resultMsg"`
	ResultCode int    `json:"resultCode"`
}

type SellItemInput struct {
	SellPrice     uint64 `json:"sellPrice"`
	SellChooseKey int    `json:"sellChooseKey"`
	SellCount     int    `json:"sellCount"`
	SellKeyType   int    `json:"sellKeyType"`
	SellMainKey   int    `json:"sellMainKey"`
	SellSubKey    int    `json:"sellSubKey"`
	IsRingBuff    bool   `json:"isRingBuff"`
	IsSealed      bool   `json:"isSealed"`
}

type SellItemOutput struct {
	ResultOutput
}

type WithdrawBuyBiddingInput struct {
	BuyNo     uint64 `json:"buyNo"`
	ChooseKey int    `json:"chooseKey"`
	Count     int    `json:"count"`
	KeyType   int    `json:"keyType"`
	MainKey   int    `json:"mainKey"`
	SubKey    int    `json:"subKey"`
}

type WithdrawBuyBiddingOutput struct {
	ResultOutput
}

type WithdrawSellBiddingInput struct {
	SellNo    uint64 `json:"sellNo"`
	ChooseKey int    `json:"chooseKey"`
	Count     int    `json:"count"`
	KeyType   int    `json:"keyType"`
	MainKey   int    `json:"mainKey"`
	SubKey    int    `json:"subKey"`
	IsSealed  bool   `json:"isSealed"`
}

type WithdrawSellBiddingOutput struct {
	ResultOutput
}
