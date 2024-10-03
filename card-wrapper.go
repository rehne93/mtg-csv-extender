package main

import (
	"fmt"
	"strconv"

	"github.com/BlueMonday/go-scryfall"
)

// gets the price of the card. uses USD and converts it, if no euro price was found.
func getPrice(card scryfall.Card) string {

	if card.Prices.EUR != "" {
		return card.Prices.EUR
	}

	priceInDollar, e := strconv.ParseFloat(card.Prices.USD, 32)

	if e != nil {
		return "0"
	}

	// TODO: Get factor from an api
	priceInEuros := priceInDollar * 0.89

	return fmt.Sprintf("%.2f", priceInEuros)
}

// Returns the color identity of the card for edh

func getColors(card scryfall.Card) string {
	var colorIdentity = ""
	for _, color := range card.ColorIdentity {
		colorIdentity += string(color)
	}
	return colorIdentity
}

// Returns the formatted value for manavalue
func getManaValue(card scryfall.Card) string {

	return fmt.Sprintf("%.0f", card.CMC)
}
