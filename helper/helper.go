package helper

var version = "1.0.0"

/*
Variabel version tidak bisa diakses dari luar package karena nama variabel
berawalan huruf kecil
*/
var Application = "Golang"

/*
Variabel version bisa diakses dari luar package karena nama variabel
berawalan huruf besar
*/

func sayHai(name string) string {
	return "Hai " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
