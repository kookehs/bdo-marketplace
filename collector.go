package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Grade string `json:"grade"`
}

type Items []Item

type ItemInfo struct {
	ID string
	Name string
	Grade string
	Price string
}

type ItemInfos []ItemInfo

func GetItemList(path string) *Items {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	i := new(Items)
	err = json.Unmarshal(f, i)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return i
}

func GetPrices(client *BDOMarketplaceClient, items *Items) ItemInfos {
	ret := make(ItemInfos, 0)

	for _, v := range *items {
		gwmsslr := client.GetWorldMarketSearchSubList(v.ID)

		if gwmsslr == nil {
			fmt.Println("Failed to get price for " + v.Name)
			continue
		}

		if len(gwmsslr.DetailList) > 0 {
			ii := ItemInfo{
				ID: v.ID,
				Name: v.Name,
				Grade: v.Grade,
				Price: strconv.Itoa(gwmsslr.DetailList[0].PricePerOne),
			}

			ret = append(ret, ii)
		}
	}

	return ret
}

func DumpToCSV(path string, items ItemInfos) {
	f, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	w := csv.NewWriter(f)
	r := []string{"id", "name", "grade", "price"}
	w.Write(r)

	for _, v := range items {
		r = []string{v.ID, v.Name, v.Grade, v.Price}
		w.Write(r)
	}

	w.Flush()
}
