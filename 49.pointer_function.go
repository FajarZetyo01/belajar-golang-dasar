package main

import "fmt"

type Alamat struct {
	Name, Country string
}

func ChangeCountryToIndonesia(addres *Alamat) {
	addres.Country = "Indonesia"
}

func main() {

	addres := Alamat{}
	ChangeCountryToIndonesia(&addres)
	fmt.Println(addres)
}
