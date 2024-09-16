package main

import (
	"context"
	"log"
	"reflect"
	"strings"

	"github.com/BlueMonday/go-scryfall"
)

/**
* Finds a Card for the given cardname and setstring
 */
func findCard(cardname string, set string) scryfall.Card {
	cardname = strings.ToLower(cardname)
	set = strings.ToLower(set)

	searchString := cardname + " (game:paper) set:" + set
	return executeRequest(searchString)
}

/**
* Fetches a german name for a card
 */
func getGermanName(card scryfall.Card) string {
	searchString := card.Name + " (game:paper) lang:de"
	result := executeRequest(searchString)

	if reflect.ValueOf(result.PrintedName).IsNil() {
		return ""
	}
	return *result.PrintedName
}

/**
* Executes a request to scryfall
 */
func executeRequest(searchString string) scryfall.Card {
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

	result, err := client.SearchCards(ctx, searchString, sco)

	if err != nil {
		card := scryfall.Card{}
		card.Name = "EMPTY"
		return card
	}

	if len(result.Cards) == 0 {
		empty := "EMPTY"
		return scryfall.Card{PrintedName: &empty}
	}

	return result.Cards[0]
}
