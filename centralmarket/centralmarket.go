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

func NewClient(url string, headers map[string]string, token, timeout string) *Client {
	duration, err := time.ParseDuration(timeout)

	if err != nil {
		log.Println(err)

		return nil
	}

	return &Client{
		BaseURL: url,
		Client: &http.Client{
			Timeout: duration,
		},
		Headers:                  headers,
		RequestVerificationToken: token,
	}
}

func (c *Client) Request(method, endpoint string, body io.Reader) []byte {
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

func (c *Client) BuyItem(input BuyItemInput) *BuyItemOutput {
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

func (c *Client) CalculateSellBidding(input CalculateSellBiddingInput) *CalculateSellBiddingOutput {
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

func (c *Client) GetItemSellBuyInfo(input GetItemSellBuyInfoInput) *GetItemSellBuyInfoOutput {
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

func (c *Client) GetMyBiddingList() *GetMyBiddingListOutput {
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

func (c *Client) GetMyWalletList() *GetMyWalletListOutput {
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

func (c *Client) GetWorldMarketHotList() *GetWorldMarketHotListOutput {
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

func (c *Client) GetWorldMarketList(input GetWorldMarketListInput) *GetWorldMarketListOutput {
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

func (c *Client) GetWorldMarketSearchList(input GetWorldMarketSearchListInput) *GetWorldMarketSearchListOutput {
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

func (c *Client) GetWorldMarketSubList(input GetWorldMarketSubListInput) *GetWorldMarketSearchListOutput {
	parameters := map[string]string{
		MainKeyString: input.MainKey,
		TokenString:   input.Token,
	}

	body := ParametersToBody(parameters)
	data := c.Request(http.MethodPost, GetWorldMarketSubListEndpoint, body)

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

func (c *Client) SellItem(input SellItemInput) *SellItemOutput {
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

func (c *Client) WithdrawBuyBidding(input WithdrawBuyBiddingInput) *WithdrawBuyBiddingOutput {
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

func (c *Client) WithdrawSellBidding(input WithdrawSellBiddingInput) *WithdrawSellBiddingOutput {
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
