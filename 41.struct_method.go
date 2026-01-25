package main

import "fmt"

/*
Struct pada Golang adalah kumpulan data yang digabungkan menjadi satu wadah.
Kalau diibaratkan, struct itu seperti benda nyata yang punya sifat-sifat (atribut).
🧍‍♂️ 3. Manusia
Struct = template manusia
Field = nama, umur, tinggi
*/

type Manusia1 struct {
	Nama, TempatLahir string
	Umur, Tinggi      int
}

func (manusia1 Manusia1) sayHello(name string) {
	fmt.Println("Hello", name, "My name is : ", manusia1.Nama)
}

func main() {

	//CARA MEMBUAT STRUCT 1
	var fajar Manusia1
	fmt.Println(fajar)
	fajar.Nama = "Fajar Setyo Pambudi"
	fajar.TempatLahir = "Kebumen"
	fajar.Umur = 25
	fajar.Tinggi = 170
	fmt.Println(fajar)

	//CARA MEMBUAT STRUCT 2
	Andi := Manusia1{
		Nama:        "Andi",
		TempatLahir: "Jakarta",
		Umur:        26,
		Tinggi:      165,
	}
	fmt.Println(Andi)

	//CARA MEMBUAT STRUCT 3 URUTAN FIELD NYA HARUS SESUAI TYPE STRUCT
	Adhli := Manusia1{"Adhli A", "Bogor", 25, 170}
	fmt.Println(Adhli)

	fajar.sayHello("Jono")

}
