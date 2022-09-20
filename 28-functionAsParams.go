package main

import "fmt"

type Filter func(string) string

func wordWithFilter(word string, filter Filter) {
	fmt.Println("hello ", filter(word))
}

func filteredWord(word string) string {
	if word == "Anjing" {
		return "..."
	} else {
		return word
	}
}

func main() {
	wordWithFilter("Farda", filteredWord)
	wordWithFilter("Anjing", filteredWord)
}
