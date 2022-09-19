package main

import "fmt"

func main() {
	// for counter <= 10 {
	// 	fmt.Println("Ini angka ke", counter)
	// 	counter++
	// }
	//short
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Ini angka ke", counter)
	// }

	// jika untuk mengakses tipe data slice

	slice := []string{"Fanny", "Arif", "Nasrudin"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//juga bisa memakai

	for index, name := range slice {
		fmt.Println("index", index, "=", name)
	}

	//jika data bertipe map gunakan ini

	person := map[string]string{
		"name":  "Fanny",
		"title": "Programmer",
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
