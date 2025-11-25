package main

import "fmt"

func main() {

	//CONTOH BREAK (JIKA BREAK ALL KONDISI LANGSUNG STOP)
	for i := 0; i < 20; i++ {
		if i == 15 {
			break
		}
		fmt.Println("Perulangan : ", i)
	}
	fmt.Println("-------------------------------------------------------------------------------------")
	//CONTOH CONTINUE (JIKA CONTINUE KONDISI MASIH BERJALAN)
	for z := 0; z < 50; z++ {
		if z%2 == 0 {
			continue
		}
		fmt.Println("Perulangan : ", z)

		//DISINI HANYA MENAMPILKAN BILANGAN YG GANJIL KARENA YG GENAP DI SKIP
	}
}
