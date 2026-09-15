package main

import (
	"log"
	"os"
)

func main () {
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hi, my name is Jacob and this file was created using GO!")
	file.Close()
}


