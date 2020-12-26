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
	ID          int
	Name        string
	Grade       int
	Enhancement int
	Price       int64
	Count       int64
	Maximum     int64
	Minimum     int64
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

	if items == nil || len(*items) == 0 {
		fmt.Println("No items to fetch.")
		return nil
	}

	for _, v := range *items {
		gwmsslr := client.GetWorldMarketSearchSubList(v.ID)

		if gwmsslr == nil {
			fmt.Println("Failed to get price for " + v.Name)
			continue
		}

		for _, dl := range gwmsslr.DetailList {
			id, err := strconv.Atoi(v.ID)

			if err != nil {
				fmt.Println(err)
				continue
			}

			grade, err := strconv.Atoi(v.Grade)

			if err != nil {
				fmt.Println(err)
				continue
			}

			gisbi := client.GetItemSellBuyInfo(id, dl.SubKey)

			if gisbi == nil {
				fmt.Println("Failed to get detailed price info for " + v.Name)
			}

			max := dl.PricePerOne
			min := dl.PricePerOne

			if gisbi != nil {
				for _, c := range gisbi.MarketConditionList {
					if c.PricePerOne > max {
						max = c.PricePerOne
					}

					if c.PricePerOne < min {
						min = c.PricePerOne
					}
				}
			}

			ii := ItemInfo{
				ID:          id,
				Name:        dl.Name,
				Grade:       grade,
				Enhancement: dl.SubKey,
				Price:       dl.PricePerOne,
				Count:       dl.Count,
				Maximum:     max,
				Minimum:     min,
			}

			ret = append(ret, ii)
		}

		fmt.Println("Retrieved price for " + v.Name)
	}

	return ret
}

func DumpToCSV(path string, items ItemInfos) {
	if items == nil {
		fmt.Println("Nothing to dump.")
		return
	}

	f, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	w := csv.NewWriter(f)
	r := []string{"id", "name", "grade", "enhancement", "maximum", "minimum", "price", "count"}

	if err := w.Write(r); err != nil {
		fmt.Println(err)
	}

	for _, v := range items {
		r = []string{
			strconv.Itoa(v.ID),
			v.Name,
			strconv.Itoa(v.Grade),
			strconv.Itoa(v.Enhancement),
			strconv.FormatInt(v.Maximum, 10),
			strconv.FormatInt(v.Minimum, 10),
			strconv.FormatInt(v.Price, 10),
			strconv.FormatInt(v.Count, 10),
		}

		if err := w.Write(r); err != nil {
			fmt.Println(err)
		}
	}

	w.Flush()
}
