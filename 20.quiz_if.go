package main

import "fmt"

func main() {
	//QUIZ LEVEL PEMULA
	//Jawaban Quiz no 1
	x := 12
	if x%2 == 0 {
		fmt.Println("Ini adalah bilangan genap")
	} else if x%2 == 1 {
		fmt.Println("Ini dalah bilangan ganjil")
	}

	//Jawaban Quiz no 2
	umur := 20
	if umur >= 17 {
		fmt.Println("Boleh membuat KTP")
	} else {
		fmt.Println("Belum boleh")
	}

	//Jawaban Quiz no 3
	nilai := 85

	if nilai >= 90 {
		fmt.Println("A")
	} else if nilai >= 80 && nilai <= 89 {
		fmt.Println("B")
	} else if nilai >= 70 && nilai <= 79 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	//Jawaban Quiz no 4
	user := "admin"
	password := "1234"

	if user == "admin" && password == "1234" {
		fmt.Println("Login sukses")
	} else {
		fmt.Println("Login gagal")
	}

	//Jawaban Quiz no 5
	total := 150000
	if total > 100000 {
		fmt.Println("diskon 10%")
	} else if total > 50000 {
		fmt.Println("diskon 5%")
	} else {
		fmt.Println("tidak ada diskon")
	}

	//QUIZ LEVEL MENENGAH
	//Jawaban Quiz no 1 Soal 1 — Validasi Usia & Status
	age := 25
	punyaKTP := false

	if age >= 17 && punyaKTP == true {
		fmt.Println("Boleh masuk")
	} else if age >= 17 && punyaKTP == false {
		fmt.Println("Harus punya KTP dulu")
	} else {
		fmt.Println("Belum cukup umur")
	}

	//Jawaban Quiz no 2 Nilai Rapor dengan Catatan
	nilaiRapot := 20
	kehadiran := 20

	if nilaiRapot >= 75 && kehadiran >= 80 {
		fmt.Println("Lulus")
	} else if nilaiRapot >= 75 && kehadiran < 80 {
		fmt.Println("Lulus Bersyarat")
	} else {
		fmt.Println("Tidak lulus")
	}

	//Jawaban Quiz no 3 Cek Promo Belanja
	totalBelanja := 120000
	memberBelanja := false

	if totalBelanja > 100000 && memberBelanja == true {
		fmt.Println("Diskon 20%")
	} else if totalBelanja > 100000 {
		fmt.Println("Diskon 10%")
	} else if totalBelanja > 50000 {
		fmt.Println("Diskon 5%")
	} else {
		fmt.Println("Tidak dapat diskon")
	}

	//Jawaban Quiz no 4 Cek Login Multi-role

	usernameMultirole := "admin"
	role := "superuser"

	if usernameMultirole == "admin" && role == "superuser" {
		fmt.Println("Akses Penuh")
	} else if usernameMultirole == "admin" && role == "editor" {
		fmt.Println("Akses Editor")
	} else if usernameMultirole == "admin" {
		fmt.Println("Akses Terbatas")
	} else {
		fmt.Println("Tidak boleh masuk")
	}

	//Jawaban Quiz no 5 Cek 3 Kondisi Angka
	angka := 10
	if angka%2 == 0 && angka%5 == 0 {
		fmt.Println("Kelipatan 2 & 5")
	} else if angka%2 == 0 {
		fmt.Println("Kelipatan 2")
	} else if angka%5 == 0 {
		fmt.Println("Kelipatan 5")
	} else {
		fmt.Println("Bukan kelipatan keduanya")
	}

}
