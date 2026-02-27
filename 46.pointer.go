package main

import "fmt"

/*
Pointer di Golang itu variabel yang menyimpan alamat memori dari variabel lain, bukan nilai langsungnya.
Ibaratnya:
Bayangkan kamu punya rumah dan alamat rumah.
Variabel biasa = rumahnya (isi di dalam rumah, misalnya meja, kasur, TV = datanya)
Pointer = alamat rumahnya (misalnya Jl. Mawar No. 10)
Kalau kamu punya alamat rumah, kamu bisa masuk ke rumah itu dan mengubah isinya.
*/

type PersonMan struct {
	Name, Address, Province string
}

func main() {
	var address1 PersonMan = PersonMan{"Fajar Setyo Pambudi", "Bekasi", "Jabar"}
	fmt.Println(address1)
	//andi := fajar

	var address2 *PersonMan = &address1
	//Jika menggunakan ini value fajar juga ikut terganti
	address2.Name = "Andi Bachdim"
	fmt.Println(address1)
	fmt.Println(address2)

}
