package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	fmt.Println(helper.Application)
	fmt.Println(helper.SayHello("Fajar"))
	//fmt.Println(helper.version) tidak bisa diakses dari luar package
}
