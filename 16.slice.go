package main

import "fmt"

func main() {

	data := []string{
		"Fajar", "Andi", "Daniel", "Joko", "Parno", "Jackie",
	}
	//SLICE ARTINYA IRISAN
	fmt.Println(data)

	slice1 := data[3:4]
	fmt.Println(slice1)
	//SLICE ADA POTONGAN CONTOH SLICE [3:4] BERARTI YG DI AMBIL
	// "Fajar", "Andi", "Daniel", |||"Joko", |||"Parno", "Jackie",
	//  0         1         2          3           4        5
	// 3 KARENA MASIH DIDALAM SLICE JADI MASIH BISA DI AMBIL BATAS RECRUSIVE ADA DI PARNO
	// fmt.Println(slice2)

	slice3 := data[:4]
	fmt.Println(slice3)

	slice4 := data[:]
	fmt.Println(slice4)

	days := []string{
		"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu",
	}
	fmt.Println(days)
	fmt.Println(days[0])
	fmt.Println(days[1])
	slice5 := days[5:]
	slice5[0] = "Sabtu Baru" //MERUBAH VALUE DATANYA
	slice5[1] = "Minggu Baru"
	slice6 := days[2:2]
	fmt.Println(slice6)
	slice7 := days[2:3]
	fmt.Println(slice7)
	fmt.Println(days)

	numberchatgpt := []int{
		10, 20, 30, 40, 50, 60,
	}
	fmt.Println(numberchatgpt)
	potong := numberchatgpt[1:4]
	fmt.Println(potong)
	potong[0] = 100
	fmt.Println(numberchatgpt)
	potong2 := numberchatgpt[2:5]
	fmt.Println(potong2[0]) // 30
	fmt.Println(potong2[1]) //40
	fmt.Println(potong2[2]) //50

	//APPEND (MENAMBAH NILAI BARU KE ARRAY
	tambahNilai := append(numberchatgpt, 30, 20) //MENAMBAHKAN NILAI 2 SEKALIGUS
	fmt.Println(tambahNilai)

}
