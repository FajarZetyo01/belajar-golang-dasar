package main

import "fmt"

func main() {
	var nilai1 = 71
	var nilai2 = 70

	var hasilakhirnilai1 = nilai1 > 70
	var hasilakhirnilai2 = nilai2 > 60

	var hasilakhir bool = hasilakhirnilai1 && hasilakhirnilai2
	fmt.Println(hasilakhir)
}
