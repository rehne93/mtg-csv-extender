package main

import (
	"flag"
)

// TODO Kann man csv auch sauberer erstellen
// TODO: Optimierung Deutscher Namen fetchen?
func main() {
	file := flag.String("input", "mtg.csv", "input csv file")
	outputFile := flag.String("output", "result.csv", "output filename")

	flag.Parse()

	records := parseCsv(*file)
	csvContent := ""

	for _, cards := range records {
		scryfallCard := findCard(cards.Cardname, cards.Set)

		if cards.Language == "DE" {
			scryfallCard = findCard(scryfallCard.Name, cards.Set)
		}

		csvContent += convertToCsvLine(scryfallCard)
		csvContent += "\n"
	}

	writeFile(*outputFile, csvContent)

}
