package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/BlueMonday/go-scryfall"
)

// TODO: Create a formated excel instead of a csv.
// TODO: Create a HTML if possible
func main() {
	file := flag.String("input", "mtg.csv", "input csv file")
	outputFile := flag.String("output", "result.csv", "output filename")

	flag.Parse()

	records := parseCsv(*file)

	var cardsList []scryfall.Card

	for idx, cards := range records {
		scryfallCard := findCard(cards.Cardname, cards.Set)

		// if the input is german, we will look for the english version to get proper prices
		// we have to look for it again
		if cards.Language == "DE" {
			englishCard := findCard(scryfallCard.Name, cards.Set)
			// if we haven't found anything we use the former card to have some data at least
			if englishCard.Name != "EMPTY" {
				scryfallCard = englishCard
			}
		}

		if scryfallCard.Name == "EMPTY" {
			fmt.Println("Error while searching for " + cards.Cardname + "(line " + strconv.Itoa(idx+1) + ")")
		}

		cardsList = append(cardsList, scryfallCard)
	}

	writeCsv(cardsList, *outputFile)

}
