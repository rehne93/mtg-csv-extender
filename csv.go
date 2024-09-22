package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/BlueMonday/go-scryfall"
)

func parseCsv(filePath string) []CsvInput {
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
	return convertToCsvInput(records)

}

func convertToCsvInput(records [][]string) []CsvInput {
	var csvInputs []CsvInput

	for _, record := range records {
		csvInput := CsvInput{
			Cardname: strings.TrimSpace(strings.ToLower(record[0])),
			Language: strings.TrimSpace(record[1]),
			Set:      strings.TrimSpace(record[2]),
		}
		csvInputs = append(csvInputs, csvInput)
	}
	return csvInputs

}

func createCsv(cards []scryfall.Card) [][]string {
	var csvData [][]string

	csvData = append(csvData, createHeader())
	for _, card := range cards {
		csvData = append(csvData, convertToDataArray(card))
	}

	return csvData
}

func createHeader() []string {
	return []string{
		"English Cardname",
		"German Cardname",
		"Mana Value",
		"Rarity",
		"Set",
		"Collector-Number",
		"Value (€)",
		"URL",
	}
}

func convertToDataArray(card scryfall.Card) []string {
	return []string{
		card.Name,
		getGermanName(card),
		getManaValue(card),
		card.Rarity,
		card.Set,
		card.CollectorNumber,
		strings.Replace(getPrice(card), ".", ",", -1),
		card.ScryfallURI,
	}
}

func writeCsv(cards []scryfall.Card, filename string) int {
	file2, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	writer := csv.NewWriter(file2)
	defer writer.Flush()

	data := createCsv(cards)

	for _, row := range data {
		writer.Write(row)
	}

	return 0
}
