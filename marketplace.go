package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

const (
	MainKeyKey     = "mainKey"
	SubKeyKey      = "subKey"
	IsUpKey        = "isUp"
	KeyTypeKey     = "keyType"
	UsingClientKey = "usingCleint"
	TokenKey       = "__RequestVerificationToken"
)

type GetItemSellBuyInfoResponse struct {
	PriceList                 []int                      `json:"priceList"`
	MarketConditionList       []ItemSellBuyInfoCondition `json:"marketConditionList"`
	BasePrice                 int                        `json:"basePrice"`
	EnchantGroup              int                        `json:"enchantGroup"`
	EnchantMaxGroup           int                        `json:"enchantMaxGroup"`
	EnchantMaterialKey        int                        `json:"enchantMaterialKey"`
	EnchantMaterialPrice      int                        `json:"enchantMaterialPrice"`
	EnchantNeedCount          int                        `json:"enchantNeedCount"`
	MaxRegisterForWorldMarket int                        `json:"maxRegisterForWorldMarket"`
	CountValue                int                        `json:"countValue"`
	SellMaxCount              int                        `json:"sellMaxCount"`
	BuyMaxCount               int                        `json:"buyMaxCount"`
	ResultCode                int                        `json:"resultCode"`
	ResultMsg                 string                     `json:"resultMsg"`
}

type GetWorldMarketSearchSubListResponse struct {
	DetailList []WorldMarketSearchSubListItem `json:"detailList"`
	ResultCode int                            `json:"resultCode"`
	ResultMsg  string                         `json:"resultMsg"`
}

type ItemSellBuyInfoCondition struct {
	SellCount   int `json:"sellCount"`
	BuyCount    int `json:"buyCount"`
	PricePerOne int `json:"pricePerOne"`
}

type WorldMarketSearchSubListItem struct {
	PricePerOne     int    `json:"pricePerOne"`
	TotalTradeCount int    `json:"totalTradeCount"`
	KeyType         int    `json:"keyType"`
	MainKey         int    `json:"mainKey"`
	SubKey          int    `json:"subKey"`
	Count           int    `json:"count"`
	Name            string `json:"name"`
	Grade           int    `json:"grade"`
	MainCategory    int    `json:"mainCategory"`
	SubCategory     int    `json:"subCategory"`
	ChooseKey       int    `json:"chooseKey"`
}

type BDOMarketplaceClient struct {
	BaseURL                  string
	Client                   *http.Client
	Headers                  map[string]string
	RequestVerificationToken string
}

func NewBDOMarketplaceClient(url string, headers map[string]string, token, timeout string) *BDOMarketplaceClient {
	d, err := time.ParseDuration(timeout)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &BDOMarketplaceClient{
		BaseURL: url,
		Client: &http.Client{
			Timeout: d,
		},
		Headers:                  headers,
		RequestVerificationToken: token,
	}
}

func (bdomc *BDOMarketplaceClient) GetItemSellBuyInfo(main, sub int) *GetItemSellBuyInfoResponse {
	e := "/Home/GetItemSellBuyInfo"

	f := map[string]string{
		TokenKey:   bdomc.RequestVerificationToken,
		MainKeyKey: strconv.Itoa(main),
		SubKeyKey:  strconv.Itoa(sub),
		IsUpKey:    "true",
		KeyTypeKey: "0",
	}

	b := new(bytes.Buffer)
	a := false

	for k, v := range f {
		if a {
			b.WriteByte('&')
		}

		a = true

		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(v)
	}

	u, err := url.Parse(bdomc.BaseURL)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	u.Path = path.Join(u.Path, e)
	r, err := http.NewRequest("POST", u.String(), b)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	d := bdomc.Post(r)

	if d == nil {
		return nil
	}

	gwmsslr := new(GetItemSellBuyInfoResponse)
	err = json.Unmarshal(d, gwmsslr)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return gwmsslr
}

func (bdomc *BDOMarketplaceClient) GetWorldMarketSearchSubList(id string) *GetWorldMarketSearchSubListResponse {
	e := "/Home/GetWorldMarketSubList"

	f := map[string]string{
		TokenKey:       bdomc.RequestVerificationToken,
		MainKeyKey:     id,
		UsingClientKey: "0",
	}

	b := new(bytes.Buffer)
	a := false

	for k, v := range f {
		if a {
			b.WriteByte('&')
		}

		a = true

		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(v)
	}

	u, err := url.Parse(bdomc.BaseURL)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	u.Path = path.Join(u.Path, e)
	r, err := http.NewRequest("POST", u.String(), b)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	d := bdomc.Post(r)

	if d == nil {
		return nil
	}

	gwmsslr := new(GetWorldMarketSearchSubListResponse)
	err = json.Unmarshal(d, gwmsslr)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return gwmsslr
}

func (bdomc *BDOMarketplaceClient) Post(request *http.Request) []byte {
	for k, v := range bdomc.Headers {
		request.Header.Set(k, v)
	}

	r, err := bdomc.Client.Do(request)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return b
}
