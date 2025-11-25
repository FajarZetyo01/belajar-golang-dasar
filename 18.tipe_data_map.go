package main

import "fmt"

func main() {
	book := map[string]string{
		"nameBook": "Cerita Masa Lalu",
		"penulis":  "Fajar Setyo Pambudi",
		"tahun":    "2015",
	}
	fmt.Println(book)
	fmt.Println(book["nameBook"])
	fmt.Println(book["penulis"])
	fmt.Println(book["tahun"])

	//MEMBUAT MAP BARU

	buku := make(map[string]string)
	buku["name"] = "Cerita Cinta Senja"
	buku["penulis"] = "Fajar Setyo Pambudi"
	buku["tahun"] = "2018"

	//Merubah value key
	buku["tahun"] = "2020"

	fmt.Println(buku)
	fmt.Println(buku["name"])
	fmt.Println(buku["penulis"])
	fmt.Println(buku["tahun"])

	//Cara menghapus Key
	delete(buku, "tahun")
	fmt.Println(buku)
}
