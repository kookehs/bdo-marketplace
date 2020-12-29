package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

type Item struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func (i Item) CSV() []string {
	csv := []string{}
	csv = append(csv, i.Name)
	csv = append(csv, strconv.Itoa(i.ID))

	return csv
}

func ParseItemList(filename string) []Item {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)

		return nil
	}

	items := []Item{}

	if err := json.Unmarshal(data, &items); err != nil {
		log.Println(err)

		return nil
	}

	return items
}
