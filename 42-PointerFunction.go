package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

func ChangeAlamat(alamat *Alamat) {
	alamat.Country = "Indonesia"
}

func main() {

	var alamat = Alamat{
		City:     "Demak",
		Province: "Jawa Tengah",
		Country:  "",
	}
	ChangeAlamat(&alamat)
	fmt.Println(alamat)

}
