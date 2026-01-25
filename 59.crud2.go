package main

import (
	"errors"
	"fmt"
)

type Mobil struct {
	Id        int
	NamaMobil string
	Tahun     int
	Merek     string
}

var mobils []Mobil
var lastID2 int = 0

func createData2(namaMobil string, tahun int, merek string) Mobil {
	lastID2++
	kendaraanInsert := Mobil{
		Id:        lastID2,
		NamaMobil: namaMobil,
		Tahun:     tahun,
		Merek:     merek,
	}
	mobils = append(mobils, kendaraanInsert)
	return kendaraanInsert
}

func readData2() []Mobil {
	return mobils
}

func updateData2(id int, namaMobil string, tahun int, merek string) (Mobil, error) {
	for i, kendaraanUpdate := range mobils {
		if kendaraanUpdate.Id == id {
			mobils[i].NamaMobil = namaMobil
			mobils[i].Tahun = tahun
			mobils[i].Merek = merek
			return kendaraanUpdate, nil
		}
	}
	return Mobil{}, errors.New("ID Tidak ditemukan")
}

func deleteData2(id int) (Mobil, error) {
	for i, deleteKendaraan := range mobils {
		if deleteKendaraan.Id == id {
			deleteData2 := deleteKendaraan
			mobils = append(mobils[:i], mobils[i+1:]...)
			return deleteData2, nil
		}
	}
	return Mobil{}, errors.New("ID tidak ada")
}

func main() {

	//ADD DATA
	createData2("Avanza Terios", 2020, "Honda")
	createData2("Avanza Terios1", 2021, "Honda")
	createData2("Avanza Terios2", 2022, "Honda")
	read := readData2()
	fmt.Println(read)

	//UPDATE DATA
	kendaraanUpdate, err := updateData2(2, "BYD Atto 1", 2025, "BYD")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Updated", kendaraanUpdate)
	}

	//DELETED DATA
	kendaraanDelete, err := deleteData2(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Deleted", kendaraanDelete)
	}
	//READ ALL DATA
	fmt.Println(readData2())
}
