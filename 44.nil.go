package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

/*
TIPE DATA YG BISA DIPAKAI NIL
Pointer,Slice,Map,Chanel,Function,Interface,Error
*/

func main() {
	data := NewMap("Fajar")
	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
