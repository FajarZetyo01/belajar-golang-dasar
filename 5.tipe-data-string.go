package main

import "fmt"

func main() {

	name := "Alexander Bautista"
	name1 := "Daniel Prasetyo"
	//LEN kalau ada spasi spasinya tetep dihitung
	fmt.Println("Fajar Setyo Pambudi")
	fmt.Println("Fajar Setyo")
	fmt.Println("Fajar")

	fmt.Println(len("Fajar Setyo Pambudi"))
	fmt.Println("Fajar Setyo"[0]) //SAAT DI RUN RESULTNYA TYPENYA MASIH BYTE
	fmt.Println("Fajar"[1])
	fmt.Println(len(name))
	fmt.Println(len(name1))
}
