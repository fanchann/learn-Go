package main

import "fmt"

func greet(name string) string {
	return "Good morning " + name
}

func main() {
	selamatPagi := greet
	fmt.Println(selamatPagi("Fanny"))
}
