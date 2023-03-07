package main

import (
	"github.com/aiviaio/go-binance/v2"
	"fmt"
	"context"
	"log"
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

	fmt.Println(symbols)
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
