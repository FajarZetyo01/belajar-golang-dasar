package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Andi"
	names[1] = "Doni"
	names[2] = "Daniel"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//MENAMBAHKAN ARRAY SEKALIGUS
	var angka = [3]int{
		10, 20, 100,
	}

	fmt.Println(angka)
	fmt.Println(angka[0])
	fmt.Println(angka[1])
	fmt.Println(angka[2])

	var angka2 = [...]int{
		90, 100, 120, 110, 1000,
	}
	fmt.Println(angka2)
	fmt.Println(len(angka2))
	fmt.Println(angka2[1])

	names1 := []string{
		"Fajar",
		"Aul",
		"Daniel",
		"Joni",
	}

	fmt.Println(names1)
	fmt.Println(len(names1))
	names1[0] = "Jono"
	fmt.Println(names1)

}
