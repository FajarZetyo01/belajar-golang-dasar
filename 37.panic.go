package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Program sedang error")
	}
}

//Panic = menghentikan program secara tiba-tiba karena terjadi error serius.
//Ibaratnya:
//“Rumah kebakaran! Semua kegiatan langsung berhenti.”
//Biasanya panic muncul saat:
//- akses indeks array yang salah
//- nil pointer
//- error fatal lain

func main() {

	runApp(true)
}
