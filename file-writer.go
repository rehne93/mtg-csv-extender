package main

import "os"

func writeFile(content string) {
	err := os.WriteFile("output.csv", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
