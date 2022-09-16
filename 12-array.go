package main

import "fmt"

func main() {

	values := [4]string{
		"Fanny",
		"Arif",
		"Nasrudin",
		"Farda",
	}
	values[1] = "Gans"
	fmt.Println(values)
}
