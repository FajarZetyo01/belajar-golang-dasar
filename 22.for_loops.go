package main

import "fmt"

func main() {
	angka := 1

	for angka <= 10 {
		fmt.Println("Perulangan angka ke : ", angka)
		angka++
	}
	fmt.Println("Perulangan selesai")

	for number := 0; number <= 100; number++ {
		fmt.Println("Perulangan number : ", number)
	}
	fmt.Println("Done")

	names := []string{"Fajar", "Andi", "Jodi", "Andreas", "Ronald"}

	//FOR MANUAL
	for i := 0; i < len(names); i++ {
		fmt.Println("Namanya adalah : ", names[i])
	}

	//FOR RANGE
	for index, name := range names {
		fmt.Println("Index : ", index, "Name : ", name)
	}

	//FOR TANPA INDEX
	for _, name := range names {
		fmt.Println("Namanya adalah : ", name)
	}

	//SOAL 1 FOR LOOPS MENGHITUNG JUMLAH ANGKA GENAP & GANJIL
	numbers := []int{4, 7, 9, 10, 13, 22}

	//CARA MENGGUNAKAN FOR LOOP MANUAL
	totalGenap := 0
	totalGanjil := 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			totalGenap++

		} else {
			totalGanjil++
		}
	}
	fmt.Println("Total bilangan Genap : ", totalGenap)
	fmt.Println("Total bilangan Genap : ", totalGanjil)

	//CARA MENGGUNAKAN FOR RANGE
	totalGenap1 := 0
	totalGanjil1 := 0

	for _, angka := range numbers {
		if angka%2 == 0 {
			totalGenap1++
		} else {
			totalGanjil1++
		}
	}
	fmt.Println("Total bilangan Genap1 : ", totalGenap1)
	fmt.Println("Total bilangan Genap1 : ", totalGanjil1)

	//SOAL 2 MENGHITUNG JUMLAH HURUF PADA KALIMAT
	namaFajar := "fajar setyo pambudi"
	target := 'a'

	jumlahHuruf := 0

	for _, ch := range namaFajar {
		if ch == target {
			jumlahHuruf++
		}
	}
	fmt.Println("Jumlah huruf A : ", jumlahHuruf)

	//SOAL 3 MENGHITUNG FIZZBUZ

	countFizzBuzz := 0
	countFizz := 0
	countBuzz := 0
	countNoFizzBuzz := 0

	for angkaFizz := 1; angkaFizz <= 100; angkaFizz++ {
		if angkaFizz%3 == 0 && angkaFizz%5 == 0 {
			fmt.Println("Indeks :", angkaFizz, " : FizzBuzz")
			countFizzBuzz++
		} else if angkaFizz%3 == 0 {
			fmt.Println("Indeks :", angkaFizz, " : Fizz")
			countFizz++
		} else if angkaFizz%5 == 0 {
			fmt.Println("Indeks :", angkaFizz, " : Buzz")
			countBuzz++
		} else {
			fmt.Println("Indeks :", angkaFizz)
			countNoFizzBuzz++
		}
	}
	fmt.Println("Total FizzBuzz : ", countFizzBuzz)
	fmt.Println("Total Fizz : ", countFizz)
	fmt.Println("Total Buzz : ", countBuzz)
	fmt.Println("Total NoFizzBuzz : ", countNoFizzBuzz)
	totalAll := countFizzBuzz + countFizz + countBuzz + countNoFizzBuzz
	fmt.Println("Total All : ", totalAll)
}
