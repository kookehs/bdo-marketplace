package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	c := new(Config)
	err = json.Unmarshal(f, c)

	if err != nil {
		fmt.Println(err)
		return
	}

	bdomc := NewBDOMarketplaceClient(
		c.URL,
		map[string]string{
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
			"Cookie":       c.Cookie,
			"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		},
		c.Token,
	)

	i := GetItemList(c.Input)
	ii := GetPrices(bdomc, i)
	DumpToCSV(c.Output, ii)
}
