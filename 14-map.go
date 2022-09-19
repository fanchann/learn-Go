package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Fanny Arif",
		"address": "Denmark",
	}

	fmt.Println(person)
	fmt.Println(person["address"])
	//add new data
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["title"])
	fmt.Println(person)
	delete(person, "address")
	fmt.Println(person)
}
