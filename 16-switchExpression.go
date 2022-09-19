package main

import "fmt"

func main() {
	name := "Nasrud"

	//switch expression and switch short
	// switch name {
	// case "Fanny":
	// 	fmt.Println("Hi Fanny")
	// case "Farda":
	// 	fmt.Println("Hi Farda")
	// default:
	// 	fmt.Println("Boleh Kenalan?")
	// }

	// switch length := len(name); length >= 7 {
	// case true:
	// 	fmt.Println("Terlalu panjang")
	// case false:
	// 	fmt.Println("Nah bener")
	// }

	//switch tanpa kondisi

	length := len(name)
	switch {
	case length > 8:
		fmt.Println("Terlalu panjang")
	case length > 6:
		fmt.Println("Masih panjang")
	default:
		fmt.Println("Nah bener")
	}
}
