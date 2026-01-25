package main

import "fmt"

// 1 QUIZ MENGHITUNG TOTAL ANGKA
func sumAll(angka ...int) int {
	total := 0

	for _, x := range angka {
		total = total + x
	}
	return total
}

// 2 QUIZ MENENTUKAN FIZZBUZZ
func fizzBuzz(angka int) int {
	for i := 0; i <= angka; i++ {
		if i%2 == 0 && i%5 == 0 {
			fmt.Println("Index : ", i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Index : ", i, "Fizz")
		} else {
			fmt.Println("Index : ", i)
		}
	}
	return angka
}

// 3 MENDETEKSI JUMLAH HURUF DI 1 KALIMAT
func countName(name string, target rune) int {
	hitungHuruf := 0
	for _, x := range name {
		if x == target {
			hitungHuruf++
		}

	}
	fmt.Println("Total hurufnya adalah : ")
	return hitungHuruf
}

// 4 MENGHITUNG JUMLAH TOTAL DISKON YANG DIDAPATKAN
func hitungDiskon(harga, diskon int) int {
	totalDiskon := harga * diskon / 100
	hargaAfterDiskon := harga - totalDiskon
	return hargaAfterDiskon
}

// 5 MENGHITUNG NILAI TERBESAR
func cariNilaiTertinggi(angka ...int) (max int, countAngka int) {
	max = angka[0]
	countAngka = 0
	for _, x := range angka {
		if x > max {
			max = x
		}
		countAngka++
	}
	fmt.Println(angka)
	fmt.Println("Jumlah Angka : ", countAngka)
	fmt.Println("Angka tersebesar adalah : ", countAngka)
	return
}

// 6 FUNCTION RETURN > 1 RETURN

// JIKA NILAI RETURN > 1 LEBIH BAIK nama variablenya di masukin di variabel awal contohnya seperti berikut
func hitungStatistik(data []int) (totalAngka int, avg float64, maxAngka int, minAngka int) {
	totalAngka = 0
	maxAngka = data[0]
	minAngka = data[0]

	for _, x := range data {
		totalAngka += x
		if x > maxAngka {
			maxAngka = x
		} else if x < minAngka {
			minAngka = x
		}
	}
	avg = float64(totalAngka) / float64(len(data))
	fmt.Println("Angkanya adalah : ", data)
	fmt.Println("Total angka : ", totalAngka)
	fmt.Println("Jumlah angka : ", len(data))
	fmt.Println("Nilai max angka : ", maxAngka)
	fmt.Println("Nilai min angka : ", minAngka)
	fmt.Println("Nilai rata-rata angka : ", avg)
	return
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4))
	fmt.Println(fizzBuzz(20))
	fmt.Println(countName("fajar setyo", 'o'))
	fmt.Println(hitungDiskon(1000000, 20))
	cariNilaiTertinggi(3, 7, 2, 10, 6)
	data := []int{10, 20, 30, 40, 50}
	hitungStatistik(data)

}
