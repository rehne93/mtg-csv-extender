#!/bin/bash
go build . && ./mtg-csv-extender -input=mtg.csv -output=result.csv