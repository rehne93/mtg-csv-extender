package main

import (
	"flag"
)

// TODO: Prüfen, welche Infos da sind um zu erweitern
// TODO: Wenn DE -> nach englischen namen suchen und query für den machen statt deutsch
// TODO: Nicht printen sondern csv bauen
// TODO: toLower
func main() {
	file := flag.String("file", "test.csv", "a string")
	flag.Parse()

	records := parseCsv(*file)

	for _, cards := range records {
		scryfallCard := findCard(cards[0], cards[2])

		if cards[1] == "DE" {
			scryfallCard = findCard(scryfallCard.Name, cards[2])
		}

		convertToCsvLine(scryfallCard)
	}

}
