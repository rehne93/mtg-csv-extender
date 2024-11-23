package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"time"

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
		time.Sleep(250 * time.Millisecond)
		if scryfallCard.Name == "EMPTY" {
			fmt.Println("Error while searching for " + cards.Cardname + "(line " + strconv.Itoa(idx+1) + ")")
		}

		cardsList = append(cardsList, scryfallCard)
	}

	fmt.Println("Finished scryfall")
	cardsList = sortCardlistBySet(cardsList)

	fmt.Println("Finished sort")

	if *output == "csv" {
		writeCsv(cardsList, *outputFile)
	}

	if *output == "html" {
		writeToFile(parseHtmlTemplate(cardsList))
	}
}

func sortCardlistBySet(cards []scryfall.Card) []scryfall.Card {

	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Name == "EMPTY" || cards[j].Name == "EMPTY" {
			return true
		}
		return cards[i].Set < cards[j].Set
	})
	return cards
}
