package main

import "fmt"

/*Dalam Golang, interface adalah kumpulan definisi perilaku (method) tanpa implementasi.
Interface hanya bilang: “struct yang ingin dianggap tipe ini harus punya method-method berikut.”

IBARAT INTERFACE DI GOLANG
🎤 1. Kontrak kerja
Interface = kontrak
Struct = pekerja
Kontrak bilang:
“Kalau mau dianggap pegawai CustomerService, kamu harus bisa:
- ngomong (Speak)
- melayani (Serve)”
Siapa pun yang memenuhi kontrak → dianggap CustomerService.
*/

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {

	return animal.Name
}

func main() {

	fajar := Person{Name: "Fajar"}
	SayHello(fajar)

	kucing := Animal{Name: "Kucing"}
	SayHello(kucing)

}
