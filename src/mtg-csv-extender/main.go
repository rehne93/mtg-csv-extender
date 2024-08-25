package main

import (
	"flag"
	"fmt"
)

// TODO: Prüfen, welche Infos da sind um zu erweitern
// TODO: Nicht printen sondern csv bauen
func main() {
	file := flag.String("file", "test.csv", "a string")
	flag.Parse()

	records := parseCsv(*file)

	for _, cards := range records {
		scryfallCard := findCard(cards[0], cards[2])

		fmt.Print(scryfallCard.Name)
		fmt.Print(" ")
		fmt.Print(scryfallCard.Prices)
		fmt.Print(" ")
		fmt.Print(scryfallCard.Set)

		fmt.Println()
	}

}
