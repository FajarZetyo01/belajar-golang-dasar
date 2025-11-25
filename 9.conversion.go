package main

import "fmt"

func main() {

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	name := "Fajar Setyo Pambudi" //Var Name
	e := name[0]                  //Var e
	estring := string(e)          // Var estring

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(estring)
}
