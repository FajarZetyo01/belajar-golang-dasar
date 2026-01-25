package main

import "fmt"

/*
Struct pada Golang adalah kumpulan data yang digabungkan menjadi satu wadah.
Kalau diibaratkan, struct itu seperti benda nyata yang punya sifat-sifat (atribut).
🧍‍♂️ 3. Manusia
Struct = template manusia
Field = nama, umur, tinggi
*/

type Manusia struct {
	Nama, TempatLahir string
	Umur, Tinggi      int
}

func main() {

	//CARA MEMBUAT STRUCT 1
	var fajar Manusia
	fmt.Println(fajar)
	fajar.Nama = "Fajar Setyo Pambudi"
	fajar.TempatLahir = "Kebumen"
	fajar.Umur = 25
	fajar.Tinggi = 170
	fmt.Println(fajar)

	//CARA MEMBUAT STRUCT 2
	Andi := Manusia{
		Nama:        "Andi",
		TempatLahir: "Jakarta",
		Umur:        26,
		Tinggi:      165,
	}
	fmt.Println(Andi)

	//CARA MEMBUAT STRUCT 3 URUTAN FIELD NYA HARUS SESUAI TYPE STRUCT
	Adhli := Manusia{"Adhli A", "Bogor", 25, 170}
	fmt.Println(Adhli)

}
