package main

import "fmt"

func main () {
	EvenNum := [5]int{0,2,4,6,8}

	for _, value := range EvenNum {
		fmt.Println(value)
	}
}
