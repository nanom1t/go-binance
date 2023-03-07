package main

import (
	"github.com/aiviaio/go-binance/v2"
	"fmt"
	"context"
	"log"
	"strconv"
)

const (
	apiKey = "api key"
	secretKey = "secret key"
	symbolsNumber = 5
)

func main() {
	// get exchange symbols
	symbols, err := GetExchangeSymbols(symbolsNumber)
	if err != nil {
		log.Fatal(err)
		return
	}

	// get price
	for _, symbol := range symbols {
		price, err := GetSymbolPrice(symbol)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(symbol, price)
	}

	fmt.Println("done")
}

// Get exchange symbols
func GetExchangeSymbols(count int) ([]string, error) {
	client := binance.NewClient(apiKey, secretKey)
	ctx := context.Background()

	response, err := client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, err
	}

	symbols := make([]string, 0, count)
	if response.Symbols != nil {
		for z := 0; z < count; z++ {
			if z < len(response.Symbols) {
				symbol := response.Symbols[z].Symbol
				symbols = append(symbols, symbol)
			}
		}
	}

	return symbols, nil
}

// Get symbol price
func GetSymbolPrice(symbol string) (float64, error) {
	client := binance.NewClient(apiKey, secretKey)
	ctx := context.Background()

	response, err := client.NewListPricesService().Symbol(symbol).Do(ctx)
	if err != nil {
		return 0, err
	}

	var price float64
	if len(response) > 0 {
		price, err = strconv.ParseFloat(response[0].Price, 64)
	}

	return price, nil
}
