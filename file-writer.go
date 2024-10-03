package main

import (
	"fmt"
	"os"
)

func writeToFile(content string) {
	f, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(content)

}
