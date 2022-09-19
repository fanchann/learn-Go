package main

import "fmt"

func main() {

	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice1 := months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := months[10:]
	fmt.Println(slice2)

	//append
	slice3 := append(slice2, "Fanny")
	fmt.Println(slice3)
	slice3[1] = "Ganti bulan"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	//make

	makeSlice := make([]string, 2, 3)
	makeSlice[0] = "Kuntul"
	makeSlice[1] = "Anjay"

	fmt.Println(makeSlice, len(makeSlice), cap(makeSlice))

	//copy
	copySlice := make([]string, 1, cap(makeSlice))
	copy(copySlice, makeSlice)
	fmt.Println(copySlice)

}
