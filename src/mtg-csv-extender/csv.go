package main

import (
	"encoding/csv"
	"log"
	"os"
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
