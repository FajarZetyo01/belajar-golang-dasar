package main

import "fmt"

func logging() {
	fmt.Println("Memanggil logging after function done")
}

func runApplication() {
	defer logging()
	fmt.Println("Program berjalan")
}

func main() {
	fmt.Println("A")
	defer fmt.Println("B") //PROGRAM INI DI EKSEKUSI KETIKA PROGRAM A & B SUDAH SELESAI RUNNING
	fmt.Println("C")
	runApplication()

	//Defer = menunda eksekusi sebuah fungsi sampai fungsi utama selesai.
	//Ibaratnya:
	//“Sebelum keluar rumah, matikan lampu.”
	//Kamu bisa bilang dari awal: “Nanti pas mau keluar, matikan lampu ya!”
}
