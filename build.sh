#!/bin/bash
env GOOS=windows GOARCH=amd64 go build . 
go build . && ./mtg-csv-extender -outputFormat=html -inputFileName=mtg.csv -outputFileName=result.csv  