package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Demak", "Jawa Tengah", "Indonesia"}
	address2 := &address1

	address2 = &Address{"Waringin Barat", "Kalimantan", "Indonesia"}
	// *address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	address3 := new(Address)
	address3.City = "Semarang"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
