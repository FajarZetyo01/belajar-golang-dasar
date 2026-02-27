package main

import (
	"fmt"
)

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

//KETIKA DIFUNCTION ADA METHOD LEBIH DISARANKAN MENGGUNAKAN POINTER pada type Struct
/*SANGAT DI REKOMENDASIKAN MENGGUNAKAN POINTER DI METHOD, SEHINGGA TIDAK BOROS MEMORY KARENA
HARUS SELALU DUPLIKASI KETIKA MEMANGGIL METHOD
*/

/*
🎯 Kapan Lebih Baik Pakai Function Biasa?
Gunakan function biasa jika:
Logic tidak terikat ke satu struct
Utility/helper function
Operasi matematis umum
Shared tool

contoh :
type User struct {}
type Admin struct {}
type Customer struct {}

func (u *User) Login()
func (a *Admin) DeleteUser()
func (c *Customer) Checkout()

Lebih scalable untuk project besar.
*/

func main() {
	m := Man{Name: "Fajar"}
	m.Married()
	fmt.Println(m)
}
