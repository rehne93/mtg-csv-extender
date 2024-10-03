package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/BlueMonday/go-scryfall"
)

func main() {
	file := flag.String("inputFileName", "mtg.csv", "input csv file")
	output := flag.String("outputFormat", "html", "outputformat, e.g. csv or html")
	outputFile := flag.String("outputFileName", "result.csv", "output filename")

	flag.Parse()

	records := parseCsv(*file)

	var cardsList []scryfall.Card

	for idx, cards := range records {
		scryfallCard := findCard(cards.Cardname, cards.Set, cards.Language, true)

		if scryfallCard.Name == "EMPTY" {
			fmt.Println("Error while searching for " + cards.Cardname + "(line " + strconv.Itoa(idx+1) + ")")
		}

		cardsList = append(cardsList, scryfallCard)
	}

	if *output == "csv" {
		writeCsv(cardsList, *outputFile)
	}

	if *output == "html" {
		writeToFile(parseHtmlTemplate(cardsList))
	}
}
