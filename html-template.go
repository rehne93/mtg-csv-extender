package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/BlueMonday/go-scryfall"
)

type HtmlData struct {
	Header []string
	Cards  []CardData
}

func parseHtmlTemplate(cards []scryfall.Card) string {
	tmpl, err := template.ParseFiles("cards.html")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var headers []string = createHeader()

	var cardData []CardData

	for _, card := range cards {
		cardData = append(cardData, convertToCardData(card))
	}

	data := HtmlData{
		Header: headers,
		Cards:  cardData,
	}
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, data)

	return tpl.String()
}
