package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai, bagi int) (int, error) {
	if bagi == 0 {
		return 0, errors.New("Pembagi tidak boleh angka 0")
	} else {
		return nilai / bagi, nil
	}
}

func main() {
	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("Hasil : ", hasil)
	} else {
		fmt.Println("Errror : ", err.Error())
	}

}
