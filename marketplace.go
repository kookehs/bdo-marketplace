package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	CookieKey      = "Cookie"
	ContentTypeKey = "Content-Type"
	MainKeyKey     = "mainKey"
	TokenKey       = "__RequestVerificationToken"
	UserAgent      = "User-Agent"
)

type GetWorldMarketSearchSubListResponse struct {
	DetailList []WorldMarketSearchSubListItem `json:"detailList"`
	ResultCode int                            `json:"resultCode"`
	ResultMsg  string                         `json:"resultMsg"`
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

func (bdomc *BDOMarketplaceClient) GetWorldMarketSearchSubList(id string) *GetWorldMarketSearchSubListResponse {
	e := "/Home/GetWorldMarketSubList"
	b := TokenKey + "=" + bdomc.RequestVerificationToken + "&" + MainKeyKey + "=" + id + "&" + "usingCleint=0"

	u, err := url.Parse(bdomc.BaseURL)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	u.Path = path.Join(u.Path, e)
	r, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(b))

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
