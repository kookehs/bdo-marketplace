package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Item struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
	ID    int    `json:"id"`
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
