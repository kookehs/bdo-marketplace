package centralmarket

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

type BuyListing struct {
	ItemListing
	AddEnchantPrice    uint64 `json:"addEnchantPrice"`
	BuyNo              uint64 `json:"buyNo"`
	RegisterMoneyCount uint64 `json:"registerMoneyCount"`
	BoughtCount        int    `json:"boughtCount"`
	LeftCount          int    `json:"leftCount"`
}

func (bl BuyListing) CSV() []string {
	csv := []string{}
	csv = append(csv, bl.Name)
	csv = append(csv, strconv.Itoa(bl.MainKey))
	csv = append(csv, strconv.Itoa(bl.Grade))
	csv = append(csv, strconv.Itoa(bl.ChooseKey))
	csv = append(csv, strconv.Itoa(bl.Count))
	csv = append(csv, strconv.Itoa(bl.KeyType))
	csv = append(csv, strconv.Itoa(bl.MainCategory))
	csv = append(csv, strconv.Itoa(bl.SubCategory))
	csv = append(csv, strconv.Itoa(bl.SubKey))
	csv = append(csv, strconv.FormatUint(bl.AddEnchantPrice, 10))
	csv = append(csv, strconv.FormatUint(bl.BuyNo, 10))
	csv = append(csv, strconv.FormatUint(bl.RegisterMoneyCount, 10))
	csv = append(csv, strconv.Itoa(bl.BoughtCount))
	csv = append(csv, strconv.Itoa(bl.LeftCount))

	return csv
}

type DetailList struct {
	ItemListing
	TotalTradeCount uint64 `json:"totalTradeCount"`
}

func (dl DetailList) CSV() []string {
	csv := []string{}
	csv = append(csv, dl.Name)
	csv = append(csv, strconv.Itoa(dl.MainKey))
	csv = append(csv, strconv.Itoa(dl.Grade))
	csv = append(csv, strconv.Itoa(dl.ChooseKey))
	csv = append(csv, strconv.Itoa(dl.Count))
	csv = append(csv, strconv.Itoa(dl.KeyType))
	csv = append(csv, strconv.Itoa(dl.MainCategory))
	csv = append(csv, strconv.Itoa(dl.SubCategory))
	csv = append(csv, strconv.Itoa(dl.SubKey))
	csv = append(csv, strconv.FormatUint(dl.TotalTradeCount, 10))

	return csv
}

type Item struct {
	Name    string `json:"name"`
	MainKey int    `json:"mainKey"`
	Grade   int    `json:"grade"`
}

func (i Item) CSV() []string {
	csv := []string{}
	csv = append(csv, i.Name)
	csv = append(csv, strconv.Itoa(i.MainKey))
	csv = append(csv, strconv.Itoa(i.Grade))

	return csv
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

func (ie ItemExtended) CSV() []string {
	csv := []string{}
	csv = append(csv, ie.Name)
	csv = append(csv, strconv.Itoa(ie.MainKey))
	csv = append(csv, strconv.Itoa(ie.Grade))
	csv = append(csv, strconv.Itoa(ie.ChooseKey))
	csv = append(csv, strconv.Itoa(ie.Count))
	csv = append(csv, strconv.Itoa(ie.KeyType))
	csv = append(csv, strconv.Itoa(ie.MainCategory))
	csv = append(csv, strconv.Itoa(ie.SubCategory))
	csv = append(csv, strconv.Itoa(ie.SubKey))

	return csv
}

type ItemListing struct {
	ItemExtended
	PricePerOne uint64 `json:"pricePerOne"`
}

func (il ItemListing) CSV() []string {
	csv := []string{}
	csv = append(csv, il.Name)
	csv = append(csv, strconv.Itoa(il.MainKey))
	csv = append(csv, strconv.Itoa(il.Grade))
	csv = append(csv, strconv.Itoa(il.ChooseKey))
	csv = append(csv, strconv.Itoa(il.Count))
	csv = append(csv, strconv.Itoa(il.KeyType))
	csv = append(csv, strconv.Itoa(il.MainCategory))
	csv = append(csv, strconv.Itoa(il.SubCategory))
	csv = append(csv, strconv.Itoa(il.SubKey))
	csv = append(csv, strconv.FormatUint(il.PricePerOne, 10))

	return csv
}

type HotListing struct {
	ItemListing
	FluctuationPrice uint64 `json:"fluctuationPrice"`
	TotalTradeCount  uint64 `json:"totalTradeCount"`
	FluctuationType  int    `json:"fluctuationType"`
	Subtype          int    `json:"substype"`
}

func (hl HotListing) CSV() []string {
	csv := []string{}
	csv = append(csv, hl.Name)
	csv = append(csv, strconv.Itoa(hl.MainKey))
	csv = append(csv, strconv.Itoa(hl.Grade))
	csv = append(csv, strconv.Itoa(hl.ChooseKey))
	csv = append(csv, strconv.Itoa(hl.Count))
	csv = append(csv, strconv.Itoa(hl.KeyType))
	csv = append(csv, strconv.Itoa(hl.MainCategory))
	csv = append(csv, strconv.Itoa(hl.SubCategory))
	csv = append(csv, strconv.Itoa(hl.SubKey))
	csv = append(csv, strconv.FormatUint(hl.PricePerOne, 10))
	csv = append(csv, strconv.FormatUint(hl.FluctuationPrice, 10))
	csv = append(csv, strconv.FormatUint(hl.TotalTradeCount, 10))
	csv = append(csv, strconv.Itoa(hl.FluctuationType))
	csv = append(csv, strconv.Itoa(hl.Subtype))

	return csv
}

type MarketCondition struct {
	PricePerOne uint64 `json:"pricePerOne"`
	BuyCount    int    `json:"buyCount"`
	SellCount   int    `json:"sellCount"`
}

func (mc MarketCondition) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.FormatUint(mc.PricePerOne, 10))
	csv = append(csv, strconv.Itoa(mc.BuyCount))
	csv = append(csv, strconv.Itoa(mc.SellCount))

	return csv
}

type MarketListing struct {
	Item
	MinPrice uint64 `json:"minPrice"`
	SumCount int    `json:"sumCount"`
}

func (ml MarketListing) CSV() []string {
	csv := []string{}
	csv = append(csv, ml.Name)
	csv = append(csv, strconv.Itoa(ml.MainKey))
	csv = append(csv, strconv.Itoa(ml.Grade))
	csv = append(csv, strconv.FormatUint(ml.MinPrice, 10))
	csv = append(csv, strconv.Itoa(ml.SumCount))

	return csv
}

type PriceHistory struct {
	Days  string `json:"days"`
	Value uint64 `json:"value"`
}

func (ph PriceHistory) CSV() []string {
	csv := []string{}
	csv = append(csv, strconv.FormatUint(ph.Value, 10))
	csv = append(csv, ph.Days)

	return csv
}

type SearchListing struct {
	Item
	SumCount      int `json:"sumCount"`
	TotalSumCount int `json:"totalSumCount"`
}

func (sl SearchListing) CSV() []string {
	csv := []string{}
	csv = append(csv, sl.Name)
	csv = append(csv, strconv.Itoa(sl.MainKey))
	csv = append(csv, strconv.Itoa(sl.Grade))
	csv = append(csv, strconv.Itoa(sl.SumCount))
	csv = append(csv, strconv.Itoa(sl.TotalSumCount))

	return csv
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

func (sl SellListing) CSV() []string {
	csv := []string{}
	csv = append(csv, sl.Name)
	csv = append(csv, strconv.Itoa(sl.MainKey))
	csv = append(csv, strconv.Itoa(sl.Grade))
	csv = append(csv, strconv.Itoa(sl.ChooseKey))
	csv = append(csv, strconv.Itoa(sl.Count))
	csv = append(csv, strconv.Itoa(sl.KeyType))
	csv = append(csv, strconv.Itoa(sl.MainCategory))
	csv = append(csv, strconv.Itoa(sl.SubCategory))
	csv = append(csv, strconv.Itoa(sl.SubKey))
	csv = append(csv, strconv.FormatUint(sl.AccumulateMoneyCount, 10))
	csv = append(csv, strconv.FormatUint(sl.AddEnchantPrice, 10))
	csv = append(csv, strconv.FormatUint(sl.EnchantMaterialPrice, 10))
	csv = append(csv, strconv.FormatUint(sl.SellNo, 10))
	csv = append(csv, strconv.Itoa(sl.EnchantNeedCount))
	csv = append(csv, strconv.Itoa(sl.LeftCount))
	csv = append(csv, strconv.Itoa(sl.SoldCount))
	csv = append(csv, strconv.FormatBool(sl.IsSealed))
	csv = append(csv, strconv.FormatBool(sl.RingBuff))

	return csv
}

type WalletListing struct {
	ItemExtended
	NationCode int  `json:"nationCode"`
	ServerNo   int  `json:"serverNo"`
	UserNo     int  `json:"userNo"`
	IsSealed   bool `json:"isSealed"`
}

func (wl WalletListing) CSV() []string {
	csv := []string{}
	csv = append(csv, wl.Name)
	csv = append(csv, strconv.Itoa(wl.MainKey))
	csv = append(csv, strconv.Itoa(wl.Grade))
	csv = append(csv, strconv.Itoa(wl.ChooseKey))
	csv = append(csv, strconv.Itoa(wl.Count))
	csv = append(csv, strconv.Itoa(wl.KeyType))
	csv = append(csv, strconv.Itoa(wl.MainCategory))
	csv = append(csv, strconv.Itoa(wl.SubCategory))
	csv = append(csv, strconv.Itoa(wl.SubKey))
	csv = append(csv, strconv.Itoa(wl.NationCode))
	csv = append(csv, strconv.Itoa(wl.ServerNo))
	csv = append(csv, strconv.Itoa(wl.UserNo))
	csv = append(csv, strconv.FormatBool(wl.IsSealed))

	return csv
}

func ParametersToBody(parameters map[string]string) *bytes.Buffer {
	delimited := false
	body := &bytes.Buffer{}

	for key, value := range parameters {
		if delimited {
			body.WriteByte('&')
		}

		delimited = true

		body.WriteString(key)
		body.WriteByte('=')
		body.WriteString(value)
	}

	return body
}

type Client struct {
	*http.Client
	Headers                  map[string]string
	BaseURL                  string
	RequestVerificationToken string
}

func NewClient(baseURL string, headers map[string]string, token, timeout, proxy string) *Client {
	duration, err := time.ParseDuration(timeout)

	if err != nil {
		log.Println(err)

		return nil
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}

	if proxy != "" {
		proxyURL, err := url.Parse(proxy)

		if err != nil {
			log.Println(err)

			return nil
		}

		transport.Proxy = http.ProxyURL(proxyURL)
	}

	return &Client{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout:   duration,
			Transport: transport,
		},
		Headers:                  headers,
		RequestVerificationToken: token,
	}
}

func (c Client) Request(method, endpoint string, body io.Reader) []byte {
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		log.Println(err)

		return nil
	}

	baseURL.Path = path.Join(baseURL.Path, endpoint)
	request, err := http.NewRequest(method, baseURL.String(), body)

	if err != nil {
		log.Println(err)

		return nil
	}

	for key, value := range c.Headers {
		request.Header.Set(key, value)
	}

	response, err := c.Do(request)

	if err != nil {
		log.Println(err)

		return nil
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)

		return nil
	}

	return data
}

func (c Client) BuyItem(input BuyItemInput) *BuyItemOutput {
	parameters := map[string]string{
		BuyPriceString:     strconv.FormatUint(input.BuyPrice, 10),
		BuyChooseKeyString: strconv.Itoa(input.BuyChooseKey),
		BuyCountString:     strconv.Itoa(input.BuyCount),
		MainKeyString:      strconv.Itoa(input.BuyMainKey),
		BuySubKeyString:    strconv.Itoa(input.BuySubKey),
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, BuyItemEndpoint, body)

	if data == nil {
		return nil
	}

	output := &BuyItemOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) CalculateSellBidding(input CalculateSellBiddingInput) *CalculateSellBiddingOutput {
	parameters := map[string]string{
		SellNoString:    strconv.FormatUint(input.SellNo, 10),
		ChooseKeyString: strconv.Itoa(input.ChooseKey),
		KeyTypeString:   strconv.Itoa(input.KeyType),
		MainKeyString:   strconv.Itoa(input.MainKey),
		SubKeyString:    strconv.Itoa(input.SubKey),
		IsSealedString:  strconv.FormatBool(input.IsSealed),
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, CalculateSellBiddingEndpoint, body)

	if data == nil {
		return nil
	}

	output := &CalculateSellBiddingOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) GetDetailList(token string, id int) []DetailList {
	input := GetWorldMarketSubListInput{
		Token:   token,
		MainKey: id,
	}

	output := c.GetWorldMarketSubList(input)

	if output == nil {
		return nil
	}

	return output.DetailList
}

func (c Client) GetItemInfo(token string, mainKey, subKey int) *GetItemSellBuyInfoOutput {
	input := GetItemSellBuyInfoInput{
		Token:   token,
		KeyType: 0,
		MainKey: mainKey,
		SubKey:  subKey,
		IsUp:    true,
	}

	return c.GetItemSellBuyInfo(input)
}

func (c Client) GetItemSellBuyInfo(input GetItemSellBuyInfoInput) *GetItemSellBuyInfoOutput {
	parameters := map[string]string{
		TokenString:   input.Token,
		MainKeyString: strconv.Itoa(input.MainKey),
		SubKeyString:  strconv.Itoa(input.SubKey),
		IsUpString:    strconv.FormatBool(input.IsUp),
		KeyTypeString: strconv.Itoa(input.KeyType),
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, GetItemSellBuyInfoEndpoint, body)

	if data == nil {
		return nil
	}

	output := &GetItemSellBuyInfoOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) GetMyBiddingList() *GetMyBiddingListOutput {
	data := c.Request(http.MethodPost, GetMyBiddingListEndpoint, nil)

	if data == nil {
		return nil
	}

	output := &GetMyBiddingListOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) GetMyWalletList() *GetMyWalletListOutput {
	data := c.Request(http.MethodPost, GetMyWalletListEndpoint, nil)

	if data == nil {
		return nil
	}

	output := &GetMyWalletListOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) GetWorldMarketHotList() *GetWorldMarketHotListOutput {
	data := c.Request(http.MethodPost, GetWorldMarketHotListEndpoint, nil)

	if data == nil {
		return nil
	}

	output := &GetWorldMarketHotListOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) GetWorldMarketList(input GetWorldMarketListInput) *GetWorldMarketListOutput {
	parameters := map[string]string{
		TokenString:   input.Token,
		MainKeyString: strconv.Itoa(input.MainKey),
		SubKeyString:  strconv.Itoa(input.SubKey),
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, GetWorldMarketListEndpoint, body)

	if data == nil {
		return nil
	}

	output := &GetWorldMarketListOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) GetWorldMarketSearchList(input GetWorldMarketSearchListInput) *GetWorldMarketSearchListOutput {
	parameters := map[string]string{
		SearchTextString: input.SearchText,
		TokenString:      input.Token,
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, GetWorldMarketSearchListEndpoint, body)

	if data == nil {
		return nil
	}

	output := &GetWorldMarketSearchListOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) GetWorldMarketSubList(input GetWorldMarketSubListInput) *GetWorldMarketSubListOutput {
	parameters := map[string]string{
		TokenString:   input.Token,
		MainKeyString: strconv.Itoa(input.MainKey),
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, GetWorldMarketSubListEndpoint, body)

	if data == nil {
		return nil
	}

	output := &GetWorldMarketSubListOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) SellItem(input SellItemInput) *SellItemOutput {
	parameters := map[string]string{
		TokenString:         c.RequestVerificationToken,
		SellPriceString:     strconv.FormatUint(input.SellPrice, 10),
		SellChooseKeyString: strconv.Itoa(input.SellChooseKey),
		SellCountString:     strconv.Itoa(input.SellCount),
		SellKeyTypeString:   strconv.Itoa(input.SellKeyType),
		SellSubKeyString:    strconv.Itoa(input.SellSubKey),
		IsRingBuffString:    strconv.FormatBool(input.IsRingBuff),
		IsSealedString:      strconv.FormatBool(input.IsSealed),
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, SellItemEndpoint, body)

	if data == nil {
		return nil
	}

	output := &SellItemOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) WithdrawBuyBidding(input WithdrawBuyBiddingInput) *WithdrawBuyBiddingOutput {
	parameters := map[string]string{
		BuyNoString:     strconv.FormatUint(input.BuyNo, 10),
		ChooseKeyString: strconv.Itoa(input.ChooseKey),
		CountString:     strconv.Itoa(input.Count),
		KeyTypeString:   strconv.Itoa(input.KeyType),
		MainKeyString:   strconv.Itoa(input.MainKey),
		SubKeyString:    strconv.Itoa(input.SubKey),
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, WithdrawBuyBiddingEndpoint, body)

	if data == nil {
		return nil
	}

	output := &WithdrawBuyBiddingOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}

func (c Client) WithdrawSellBidding(input WithdrawSellBiddingInput) *WithdrawSellBiddingOutput {
	parameters := map[string]string{
		SellNoString:    strconv.FormatUint(input.SellNo, 10),
		ChooseKeyString: strconv.Itoa(input.ChooseKey),
		CountString:     strconv.Itoa(input.Count),
		KeyTypeString:   strconv.Itoa(input.KeyType),
		MainKeyString:   strconv.Itoa(input.MainKey),
		SubKeyString:    strconv.Itoa(input.SubKey),
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, WithdrawSellBiddingEndpoint, body)

	if data == nil {
		return nil
	}

	output := &WithdrawSellBiddingOutput{}

	if err := json.Unmarshal(data, output); err != nil {
		log.Println(err)

		return nil
	}

	return output
}
