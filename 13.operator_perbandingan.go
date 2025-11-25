package main

import "fmt"

func main() {

	var name1 = "Fajar"
	var name2 = "Andi"

	var result bool = name1 == name2
	fmt.Println(result)

	var angka1 = 10
	var angka2 = 20

	var result1 bool = angka1 > angka2
	fmt.Println(result1)

	var angka3 = 10
	var angka4 = 20

	var result2 bool = angka3 != angka4 //TIDAK SAMA DENGAN
	fmt.Println(result2)
}
