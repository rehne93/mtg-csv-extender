#!/bin/bash
env GOOS=windows GOARCH=amd64 go build . 
go build . && ./mtg-csv-extender -input=mtg.csv -output=result.csv