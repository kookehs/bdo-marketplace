package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
	"sync"

	"github.com/kookehs/bdo-marketplace/centralmarket"
)

const (
	defaultConfigFilename = "./config.json"
	logFilename           = "error.log"
)

func main() {
	logFile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err != nil {
		log.Println(err)
	}

	defer func() {
		if err := logFile.Close(); err != nil {
			log.Println(err)
		}
	}()

	log.SetOutput(logFile)

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

	if config.Headers.CookieToken == "" || config.FormToken == "" {
		log.Println("CookieToken and FormToken must be present in config file")

		return
	}

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

	waitGroup := sync.WaitGroup{}

	for _, pair := range config.Files {
		waitGroup.Add(1)

		go func(client *centralmarket.Client, pair FilePair) {
			defer waitGroup.Done()
			GetMarketData(client, pair)
		}(client, pair)
	}

	waitGroup.Wait()
}

func CSVHeaders(structType reflect.Type) []string {
	headers := []string{}

	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	if structType.Kind() != reflect.Struct {
		return nil
	}

	queue := []reflect.Type{structType}
	visited := map[string]bool{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if _, ok := visited[node.Name()]; ok {
			continue
		}

		visited[node.Name()] = true

		for index := 0; index < node.NumField(); index++ {
			field := node.Field(index)
			tag := field.Tag.Get("json")

			if tag != "" {
				headers = append(headers, tag)
			}

			if field.Anonymous {
				queue = append(queue, field.Type)
			}
		}
	}

	return headers
}

func DumpToCSV(filename string, rows [][]string) {
	if rows == nil {
		return
	}

	directory := path.Dir(filename)

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err := os.MkdirAll(directory, os.ModePerm); err != nil {
			log.Println(err)

			return
		}
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

func GetMarketData(client *centralmarket.Client, pair FilePair) {
	rows := [][]string{}

	labels := append([]string{}, CSVHeaders(reflect.TypeOf(centralmarket.Item{}))...)
	labels = append(labels, CSVHeaders(reflect.TypeOf(centralmarket.GetItemSellBuyInfoOutput{}))...)
	labels = append(labels, CSVHeaders(reflect.TypeOf(centralmarket.ItemPrices{}))...)
	labels = append(labels, "count")
	rows = append(rows, labels)

	for _, item := range ParseItemList(pair.Input) {
		detailList := client.GetDetailList(client.RequestVerificationToken, item.ID)

		for _, detail := range detailList {
			info := client.GetItemInfo(client.RequestVerificationToken, detail.MainKey, detail.SubKey)
			prices := info.Prices()

			fields := append([]string{}, detail.Item.CSV()...)
			fields = append(fields, info.CSV()...)
			fields = append(fields, prices.CSV()...)
			fields = append(fields, strconv.Itoa(detail.Count))
			rows = append(rows, fields)

			history := info.PriceHistory()

			if len(history) > 0 {
				historyCSV := [][]string{CSVHeaders(reflect.TypeOf(history[0]))}
				historyCSV = append(historyCSV, history.CSV()...)
				DumpToCSV("history/"+detail.Item.Name+"."+strconv.Itoa(detail.SubKey)+".csv", historyCSV)
			}
		}
	}

	DumpToCSV(pair.Output, rows)
}
