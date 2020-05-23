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
			"Cookie":       TokenKey + "=" + c.Cookie,
			"User-Agent":   c.UserAgent,
		},
		c.Token,
		c.Timeout,
	)

	if bdomc == nil {
		return
	}

	i := GetItemList(c.Input)
	ii := GetPrices(bdomc, i)
	DumpToCSV(c.Output, ii)
}
