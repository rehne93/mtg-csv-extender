# MTG-CSV-Importer
Converts your mtg collection from a simple csv-format into a extender version including prices and additional informations.

## Input-Specification
The input contains three columns:
- Name of the card, in german or english
- Language (currently DE or EN)
- Set-Shorthand (eg. SOK for Savior of Kamigawa)

The input will be parsed and an extended csv will be generated. The information is gathered from the scryfall-api.
A simple example csv is included.