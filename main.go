package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

// TODO Kann man csv auch sauberer erstellen
func main() {
	file := flag.String("input", "mtg.csv", "input csv file")
	outputFile := flag.String("output", "result.csv", "output filename")

	flag.Parse()

	records := parseCsv(*file)
	csvContent := ""

	line := 0

	for _, cards := range records {
		scryfallCard := findCard(cards.Cardname, cards.Set)
		fmt.Print(scryfallCard.Name + " " + cards.Language)
		fmt.Println()

		if cards.Language == "DE" {
			englishCard := findCard(scryfallCard.Name, cards.Set)
			if englishCard.Name != "EMPTY" {
				scryfallCard = englishCard
			}
		}

		if scryfallCard.Name == "EMPTY" {
			fmt.Println("Error while searching for " + cards.Cardname + "(line " + strconv.Itoa(line) + ")")
		}

		time.Sleep(10 * time.Millisecond)
		csvContent += convertToCsvLine(scryfallCard)
		csvContent += "\n"
		line++
	}

	writeFile(*outputFile, csvContent)

}
