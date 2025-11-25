package main

import "fmt"

func main() {
	newSlice := make([]string, 2, 5)
	//2 (Len panjang), 5 (kapasitas jumlah)
	newSlice[0] = "Fajar"
	newSlice[1] = "Andi"
	//newSlice[2] = "Jono" error karena kita sudah set len nya max 2
	//Jika kita ingin menambah datanya bisa juga menggunakan append

	fmt.Println(newSlice)
	fmt.Println("Nilai maksimal adalah : ", len(newSlice))
	fmt.Println("Kapasitas maksimal adalah : ", cap(newSlice))
	newSlice1 := append(newSlice, "Dono")

	fmt.Println(newSlice1)
	fmt.Println("Nilai maksimal adalah : ", len(newSlice1))
	fmt.Println("Kapasitas maksimal adalah : ", cap(newSlice1))

	//COPY SLICE
	dataSiswa := []string{
		"Andi", "Jono", "Daniel", "Zacky", "Johanes",
	}
	fmt.Println(dataSiswa)

	//Case Copy Slice
	fromSiswa := dataSiswa[:]
	toSiswa := make([]string, len(fromSiswa), cap(fromSiswa))
	copy(toSiswa, fromSiswa)
	fmt.Println(fromSiswa)
	fmt.Println(toSiswa)

	//PERBEDAAN SLICE & ARRAY
	numbers1 := [...]int{1, 2, 3, 4, 5} //Ini adalah array
	numbers2 := []int{1, 2, 3, 4, 5}    //Ini adalah slice

	fmt.Println(numbers1)
	fmt.Println(numbers2)

	//Note : Jika kita menggunakan pemrograman Golang kita menggunakan Slice bukan Array

}
