package main

import "fmt"

type Address1 struct {
	Name, Country, Province string
}

func (a *Address1) ChangeCountry(newCountry string) error {
	if newCountry == "" {
		return fmt.Errorf("Data tidak boleh kosong")
	}
	a.Country = newCountry
	return nil
}

func main() {

	//DATA AWAL
	address1 := Address1{"Fajar", "Indonesia", "Jakarta"}
	address2 := Address1{"Andi", "Malaysia", "Penang"}

	//CALL DATA
	fmt.Println(address1)
	fmt.Println(address2, "Sebelum")

	//RUBAH DATA
	address2.ChangeCountry("Singapura")
	fmt.Println(address2, "Sesudah") //Walaupun address2 bukan pointer, Go akan otomatis mengubahnya menjadi pointer jika method membutuhkan pointer receiver.

	var err = address2.ChangeCountry("Thailand")
	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println(address2)
}
