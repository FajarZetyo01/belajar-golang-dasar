package main

import "fmt"

func main() {

	name := "Daniel"

	switch name {
	case "Fajar":
		fmt.Println("Nama kamu Fajar")
	case "Eko":
		fmt.Println("Nama kamu bukan Fajar ")
	default:
		fmt.Println("Hai salam kenal!")

	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama terlalu pendek")
	}

	name = "Daniel"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sesuai sesuai")
	}

}
