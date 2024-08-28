package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/BlueMonday/go-scryfall"
)

/**
* Finds a Card for the given cardname and setstring
 */
func findCard(cardname string, set string) scryfall.Card {
	cardname = strings.ToLower(cardname)
	set = strings.ToLower(set)

	ctx := context.Background()

	client, err := scryfall.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	sco := scryfall.SearchCardsOptions{
		Unique:              scryfall.UniqueModePrints,
		Order:               scryfall.OrderSet,
		Dir:                 scryfall.DirAuto,
		IncludeMultilingual: true,
	}

	searchString := cardname + " (game:paper) set:" + set

	result, err := client.SearchCards(ctx, searchString, sco)

	if err != nil {
		fmt.Println("Error while searching " + searchString)
		card := scryfall.Card{}
		return card
	}

	return result.Cards[0]
}
