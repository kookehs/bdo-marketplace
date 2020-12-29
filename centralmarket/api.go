package centralmarket

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
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
	SessionString       = "ASP.NET_SessionId"
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

func (bio BuyItemOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(bio.ResultCode))
	csv = append(csv, bio.ResultMsg)

	return csv
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

func (csbo CalculateSellBiddingOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(csbo.ResultCode))
	csv = append(csv, csbo.ResultMsg)

	return csv
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

func (gisbio GetItemSellBuyInfoOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(gisbio.ResultCode))
	csv = append(csv, fmt.Sprintf("%q", gisbio.ResultMsg))

	bytes, err := json.Marshal(gisbio.MarketConditionList)

	if err != nil {
		return nil
	}

	csv = append(csv, fmt.Sprintf("%q", string(bytes)))

	bytes, err = json.Marshal(gisbio.PriceList)

	if err != nil {
		return nil
	}

	csv = append(csv, fmt.Sprintf("%q", string(bytes)))
	csv = append(csv, strconv.FormatUint(gisbio.BasePrice, 10))
	csv = append(csv, strconv.FormatUint(gisbio.EnchantMaterialPrice, 10))
	csv = append(csv, strconv.Itoa(gisbio.BuyMaxCount))
	csv = append(csv, strconv.Itoa(gisbio.CountValue))
	csv = append(csv, strconv.Itoa(gisbio.EnchantGroup))
	csv = append(csv, strconv.Itoa(gisbio.EnchantGroupMax))
	csv = append(csv, strconv.Itoa(gisbio.EnchantMaterialKey))
	csv = append(csv, strconv.Itoa(gisbio.EnchantNeedCount))
	csv = append(csv, strconv.Itoa(gisbio.MaxRegisterForWorldMarket))
	csv = append(csv, strconv.Itoa(gisbio.SellMaxCount))

	return csv
}

func (gisbio GetItemSellBuyInfoOutput) MinMaxPrices() (uint64, uint64) {
	max := uint64(0)
	min := uint64(math.MaxUint64)

	for _, condition := range gisbio.MarketConditionList {
		if condition.PricePerOne > max {
			max = condition.PricePerOne
		}

		if condition.PricePerOne < min {
			min = condition.PricePerOne
		}
	}

	return min, max
}

func (gisbio GetItemSellBuyInfoOutput) PriceHistory() PriceHistories {
	history := PriceHistories{}
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

func (gmblo GetMyBiddingListOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(gmblo.ResultCode))
	csv = append(csv, gmblo.ResultMsg)

	bytes, err := json.Marshal(gmblo.BuyList)

	if err != nil {
		return nil
	}

	csv = append(csv, fmt.Sprintf("%q", string(bytes)))

	bytes, err = json.Marshal(gmblo.SellList)

	if err != nil {
		return nil
	}

	csv = append(csv, fmt.Sprintf("%q", string(bytes)))

	return csv
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

func (gmwlo GetMyWalletListOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(gmwlo.ResultCode))
	csv = append(csv, gmwlo.ResultMsg)

	bytes, err := json.Marshal(gmwlo.MyWalletList)

	if err != nil {
		return nil
	}

	csv = append(csv, fmt.Sprintf("%q", string(bytes)))

	csv = append(csv, strconv.FormatUint(gmwlo.FeeRate, 10))
	csv = append(csv, strconv.Itoa(gmwlo.RingBuffCount))
	csv = append(csv, strconv.Itoa(gmwlo.AddWeight))
	csv = append(csv, strconv.Itoa(gmwlo.MaxWeight))
	csv = append(csv, strconv.Itoa(gmwlo.TotalWeight))
	csv = append(csv, strconv.FormatBool(gmwlo.UseAddWeightBuff))
	csv = append(csv, strconv.FormatBool(gmwlo.UseValuePackage))

	return csv
}

type GetWorldMarketHotListOutput struct {
	ResultOutput
	HotList []HotListing `json:"hotList"`
}

func (gwmhlo GetWorldMarketHotListOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(gwmhlo.ResultCode))
	csv = append(csv, gwmhlo.ResultMsg)

	bytes, err := json.Marshal(gwmhlo.HotList)

	if err != nil {
		return nil
	}

	csv = append(csv, fmt.Sprintf("%q", string(bytes)))

	return csv
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

func (gwmlo GetWorldMarketListOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(gwmlo.ResultCode))
	csv = append(csv, gwmlo.ResultMsg)

	bytes, err := json.Marshal(gwmlo.MarketList)

	if err != nil {
		return nil
	}

	csv = append(csv, fmt.Sprintf("%q", string(bytes)))

	return csv
}

type GetWorldMarketSearchListInput struct {
	SearchText string `json:"searchText"`
	Token      string `json:"__RequestVerificationToken"`
}

type GetWorldMarketSearchListOutput struct {
	ResultOutput
	List []SearchListing `json:"list"`
}

func (gwmslo GetWorldMarketSearchListOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(gwmslo.ResultCode))
	csv = append(csv, gwmslo.ResultMsg)

	bytes, err := json.Marshal(gwmslo.List)

	if err != nil {
		return nil
	}

	csv = append(csv, fmt.Sprintf("%q", string(bytes)))

	return csv
}

type GetWorldMarketSubListInput struct {
	Token   string `json:"__RequestVerificationToken"`
	MainKey int    `json:"mainKey"`
}

type GetWorldMarketSubListOutput struct {
	ResultOutput
	DetailList []DetailList `json:"detailList"`
}

func (gwmslo GetWorldMarketSubListOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(gwmslo.ResultCode))
	csv = append(csv, gwmslo.ResultMsg)

	bytes, err := json.Marshal(gwmslo.DetailList)

	if err != nil {
		return nil
	}

	csv = append(csv, fmt.Sprintf("%q", string(bytes)))

	return csv
}

type ResultOutput struct {
	ResultMsg  string `json:"resultMsg"`
	ResultCode int    `json:"resultCode"`
}

func (ro ResultOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(ro.ResultCode))
	csv = append(csv, ro.ResultMsg)

	return csv
}

type SellItemInput struct {
	Token         string `json:"__RequestVerificationToken"`
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

func (sio SellItemOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(sio.ResultCode))
	csv = append(csv, sio.ResultMsg)

	return csv
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

func (wbbo WithdrawBuyBiddingOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(wbbo.ResultCode))
	csv = append(csv, wbbo.ResultMsg)

	return csv
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

func (wsbo WithdrawSellBiddingOutput) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.Itoa(wsbo.ResultCode))
	csv = append(csv, wsbo.ResultMsg)

	return csv
}
