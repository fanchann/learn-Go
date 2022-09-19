package main

import "fmt"

func main() {
	name := "Nasrudin"

	//switch expression and switch short
	// switch name {
	// case "Fanny":
	// 	fmt.Println("Hi Fanny")
	// case "Farda":
	// 	fmt.Println("Hi Farda")
	// default:
	// 	fmt.Println("Boleh Kenalan?")
	// }

	switch length := len(name); length >= 7 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Nah bener")
	}
}
