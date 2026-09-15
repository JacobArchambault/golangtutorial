package main

import (
	"log"
	"os"
	"io/ioutil"
	"fmt"
)

func main () {
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hi, my name is Jacob and this file was created using GO!")
	file.Close()
	
	stream, err := ioutil.ReadFile("sample.txt")

	s1 := string(stream)

	fmt.Println(s1)
}


