package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/BlueMonday/go-scryfall"
)

// parses the csv in the given path
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

// Converts the records to the internal data structure
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

// writes the list of cards to csv data
func createCsv(cards []scryfall.Card) [][]string {
	var csvData [][]string

	csvData = append(csvData, createHeader())
	for _, card := range cards {
		csvData = append(csvData, convertToDataArray(card))
	}

	return csvData
}

// creates the header for csv or html
func createHeader() []string {
	return []string{
		"English Cardname",
		"German Cardname",
		"Mana Value",
		"Color Identity",
		"Rarity",
		"Set",
		"Collector-Number",
		"Value (€)",
		"URL",
		"Image URL",
	}
}

// Converts a scryfall card to a data array
func convertToDataArray(card scryfall.Card) []string {
	return []string{
		card.Name,
		getGermanName(card),
		getManaValue(card),
		getColors(card),
		strings.ToUpper(card.Rarity),
		strings.ToUpper(card.Set),
		card.CollectorNumber,
		strings.Replace(getPrice(card), ".", ",", -1),
		card.ScryfallURI,
		card.ImageURIs.Small,
	}
}

// converts a scryfall card to carddata for html
func convertToCardData(card scryfall.Card) CardData {
	if card.Name == "EMPTY" {
		return CardData{}
	}

	var imageUri = ""
	if card.ImageURIs != nil {
		imageUri = card.ImageURIs.Small
	}

	return CardData{
		Cardname:        card.Name,
		GermanCardname:  getGermanName(card),
		Manavalue:       getManaValue(card),
		Colors:          getColors(card),
		Rarity:          strings.ToUpper(card.Rarity),
		Set:             strings.ToUpper(card.Set),
		CollectorNumber: card.CollectorNumber,
		Price:           strings.Replace(getPrice(card), ".", ",", -1),
		ScryfallUri:     card.ScryfallURI,
		ImageUri:        imageUri,
	}
}

// writes a csv to the disk
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
