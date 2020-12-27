package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/kookehs/bdo-marketplace/centralmarket"
)

const (
	defaultConfigFilename = "./config.json"
)

func main() {
	configFilename := flag.String("config", "config.json", "location of config.json")
	flag.Parse()

	filename := defaultConfigFilename

	if configFilename != nil {
		filename = *configFilename
	}

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)

		return
	}

	config := &Config{}

	if err := json.Unmarshal(data, config); err != nil {
		log.Println(err)

		return
	}

	cookieToken := centralmarket.TokenString + "=" + config.Headers.CookieToken
	session := centralmarket.SessionString + "=" + config.Headers.Session
	cookie := cookieToken + ";" + session

	client := centralmarket.NewClient(
		config.URL,
		map[string]string{
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
			"Cookie":       cookie,
			"User-Agent":   config.Headers.UserAgent,
		},
		config.FormToken,
		config.Timeout,
		config.Proxy,
	)

	if client == nil {
		return
	}

	for _, pair := range config.Files {
		headers := false
		rows := [][]string{}
		items := ParseItemList(pair.Input)

		for _, item := range items {
			detailList := client.GetDetailList(client.RequestVerificationToken, item.ID)

			for _, detail := range detailList {
				info := client.GetItemInfo(client.RequestVerificationToken, detail.MainKey, detail.SubKey)

				if !headers {
					headers = true
					rows = append(rows, CSVHeaders(reflect.TypeOf(info)))
				}

				rows = append(rows, info.CSV())
			}
		}

		DumpToCSV(pair.Output, rows)
	}
}

func CSVHeaders(structType reflect.Type) []string {
	headers := []string{}

	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	if structType.Kind() != reflect.Struct {
		return nil
	}

	for index := 0; index < structType.NumField(); index++ {
		field := structType.Field(index)
		tag := field.Tag.Get("json")

		if tag != "" {
			headers = append(headers, tag)
		}

		if field.Anonymous {
			headers = append(headers, CSVHeaders(field.Type)...)
		}
	}

	return headers
}

func DumpToCSV(filename string, rows [][]string) {
	if rows == nil {
		return
	}

	file, err := os.Create(filename)

	if err != nil {
		log.Println(err)

		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()

	writer := csv.NewWriter(file)

	if err := writer.WriteAll(rows); err != nil {
		log.Println(err)
	}

	writer.Flush()
}
