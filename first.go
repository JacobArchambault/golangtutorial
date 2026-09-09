package main

import "fmt"

func main () {
	var name string = "Arrya Paul"
	const pi float64 = 3.14159
	fmt.Println(len(name))
	fmt.Println(name + "is a chill dude")
	fmt.Printf("%.3f \n",pi)
	fmt.Printf("%T \n", name);
}
