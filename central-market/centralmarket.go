package centralmarket

type BuyListing struct {
	ItemListing
	AddEnchantPrice    uint64 `json:"addEnchantPrice"`
	BuyNo              uint64 `json:"buyNo"`
	RegisterMoneyCount uint64 `json:"registerMoneyCount"`
	BoughtCount        int    `json:"boughtCount"`
	LeftCount          int    `json:"leftCount"`
}

type DetailList struct {
	ItemListing
	TotalTradeCount uint64 `json:"totalTradeCount"`
}

type Item struct {
	Name    string `json:"name"`
	MainKey int    `json:"mainKey"`
	Grade   int    `json:"grade"`
}

type ItemExtended struct {
	Item
	ChooseKey    int `json:"chooseKey"`
	Count        int `json:"count"`
	KeyType      int `json:"keyType"`
	MainCategory int `json:"mainCategory"`
	SubCategory  int `json:"subCategory"`
	SubKey       int `json:"subKey"`
}

type ItemListing struct {
	ItemExtended
	PricePerOne uint64 `json:"pricePerOne"`
}

type HotListing struct {
	ItemListing
	FluctuationPrice uint64 `json:"fluctuationPrice"`
	TotalTradeCount  uint64 `json:"totalTradeCount"`
	FluctuationType  int    `json:"fluctuationType"`
	Subtype          int    `json:"substype"`
}

type MarketCondition struct {
	PricePerOne uint64 `json:"pricePerOne"`
	BuyCount    int    `json:"buyCount"`
	SellCount   int    `json:"sellCount"`
}

type MarketListing struct {
	Item
	MinPrice uint64 `json:"minPrice"`
	SumCount int    `json:"sumCount"`
}

type PriceHistory struct {
	Days  string `json:"days"`
	Value uint64 `json:"value"`
}

type SearchListing struct {
	Name          string `json:"name"`
	Grade         int    `json:"grade"`
	MainKey       int    `json:"mainKey"`
	SumCount      int    `json:"sumCount"`
	TotalSumCount int    `json:"totalSumCount"`
}

type SellListing struct {
	ItemListing
	AccumulateMoneyCount uint64 `json:"accumulateMoneyCount"`
	AddEnchantPrice      uint64 `json:"addEnchantPrice"`
	EnchantMaterialPrice uint64 `json:"enchantMaterialPrice"`
	SellNo               uint64 `json:"sellNo"`
	EnchantNeedCount     int    `json:"enchantNeedCount"`
	LeftCount            int    `json:"leftCount"`
	SoldCount            int    `json:"soldCount"`
	IsSealed             bool   `json:"isSealed"`
	RingBuff             bool   `json:"ringBuff"`
}

type WalletListing struct {
	ItemExtended
	NationCode int  `json:"nationCode"`
	ServerNo   int  `json:"serverNo"`
	UserNo     int  `json:"userNo"`
	IsSealed   bool `json:"isSealed"`
}
