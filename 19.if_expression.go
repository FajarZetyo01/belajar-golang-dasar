package main

import "fmt"

func main() {
	name := "daniel"
	if name == "Fajar" {
		fmt.Println("Ya benar Fajar!")
	} else if name == "Daniel" {
		fmt.Println("Ya benar daniel")
	} else if name == "Jono" {
		fmt.Println("Ya benar Jono !!!")
	} else {
		fmt.Println("Kamu orang lain ya!")
	}

	if panjangNama := len(name); panjangNama > 5 {
		fmt.Println("Nama sudah pas")
	} else {
		fmt.Println("Nama terlalu panjang")
	}

	//Latihan IF

	username := "fajar"
	password := "12345"

	if username == "Fajar" && password == "12345" {
		fmt.Println("Login sukses!!!")
	} else {
		fmt.Println("Username atau Password salah!!!")
	}
}
