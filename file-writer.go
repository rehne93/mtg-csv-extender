package main

import (
	"fmt"
	"os"
)

func writeToFile(content string) {
	f, err := os.Create("test.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(content)

}
