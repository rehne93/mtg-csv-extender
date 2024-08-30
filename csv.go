package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
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
			Cardname: strings.ToLower(record[0]),
			Language: record[1],
			Set:      record[2],
		}
		csvInputs = append(csvInputs, csvInput)
	}
	return csvInputs

}

func convertToCsvLine(card scryfall.Card) string {
	csvString := ""
	csvString += "\"" + card.Name + "\""
	csvString += ";"
	csvString += "\"" + getGermanName(card) + "\""
	csvString += ";"
	csvString += strconv.FormatFloat(card.CMC, 'f', -1, 64)
	csvString += ";"
	csvString += string(card.Lang)
	csvString += ";"
	csvString += strings.Replace(card.Prices.EUR, ".", ",", -1)
	csvString += ";"
	csvString += "\"" + card.Set + "\""
	csvString += ";"
	csvString += card.Rarity
	return csvString
}
