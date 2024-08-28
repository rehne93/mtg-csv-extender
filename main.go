package main

import (
	"flag"
)

// TODO Kann man csv auch sauberer erstellen
// TODO CMC statt liste?
// TODO Output filename
// TODO Dokumentieren und schauen ob man an Sprachen kommt für Name DE und ENgl
func main() {
	file := flag.String("file", "test.csv", "a string")
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

	writeFile(csvContent)

}
