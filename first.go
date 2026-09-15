package main

import "fmt"

func main () {
	StudentAge := make(map[string] int)

	StudentAge["Arrya"] = 23
	StudentAge["Saurabh"] = 27
	StudentAge["Prerna"] = 27
	StudentAge["Akrati"] = 19
	StudentAge["Sahiti"] = 42
	StudentAge["Kirti"] = 22

	fmt.Println(StudentAge)
}
