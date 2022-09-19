package main

import "fmt"

func main() {
	name := "Fard"

	//if with short statement
	if length := len(name); length >= 5 {
		fmt.Println("Nama terlalu panjang")
		fmt.Println(len(name))
	} else {
		fmt.Println("Hi", name)
	}
}
