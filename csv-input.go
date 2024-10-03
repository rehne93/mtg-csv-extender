package main

type CsvInput struct {
	Cardname string
	Language string
	Set      string
}

type CardData struct {
	Cardname        string
	GermanCardname  string
	Manavalue       string
	Colors          string
	Rarity          string
	Set             string
	CollectorNumber string
	Price           string
	ScryfallUri     string
	ImageUri        string
}
