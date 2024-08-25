package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/BlueMonday/go-scryfall"
)

func parseCsv(filePath string) [][]string {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("unable to read input")
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Could not read csv")
	}
	return records

}

func convertToCsvLine(card scryfall.Card) string {
	fmt.Print(card.Name + " ")
	fmt.Print(card.Lang + " ")
	fmt.Print(card.ColorIdentity)
	fmt.Print(card.Prices.EUR + "€ ")
	fmt.Print(card.Prices.USD + "$")
	fmt.Println()
	return ""
}
