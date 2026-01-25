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

type PersonMan1 struct {
	Name, Address, Province string
}

func main() {
	fajar := PersonMan1{"Fajar Setyo Pambudi", "Bekasi", "Jabar"}
	fmt.Println(fajar)

	andi := &fajar
	//andi.Name = "Andi Bachdim"
	*andi = PersonMan1{"Andi A", "Bogor", "Jabar"}
	fmt.Println(andi) //Value andi ikut berubah mengikuti pointer fajar
	fmt.Println(fajar)
}
