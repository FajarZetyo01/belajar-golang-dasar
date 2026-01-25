package main

import (
	"errors"
	"fmt"
)

type Orang struct {
	Id        int
	Nama      string
	Umur      int
	Alamat    string
	Pekerjaan string
}

var orangs []Orang
var lastID1 int = 0

func createData1(nama string, umur int, alamat string, pekerjaan string) Orang {
	lastID1++
	orang := Orang{
		Id:        lastID1,
		Nama:      nama,
		Umur:      umur,
		Alamat:    alamat,
		Pekerjaan: pekerjaan,
	}
	orangs = append(orangs, orang)
	return orang
}

func readData1() []Orang {
	return orangs
}

func updateData1(id int, nama string, umur int, alamat string, pekerjaan string) (Orang, error) {
	for i, user := range orangs {
		if user.Id == id {
			orangs[i].Nama = nama
			orangs[i].Umur = umur
			orangs[i].Alamat = alamat
			orangs[i].Pekerjaan = pekerjaan
			return user, nil
		}
	}
	return Orang{}, errors.New("User tidak ditemukan")
}

func deletedData1(id int) (Orang, error) {
	for i, user := range orangs {
		if user.Id == id {
			deletedData1 := user
			orangs = append(orangs[:i], orangs[i+1:]...)
			return deletedData1, nil
		}
	}
	return Orang{}, errors.New("ID tidak ada")
}

func main() {

	//ADD DATA
	createData1("fajar", 25, "bekasi", "Programmer")
	createData1("Zikri", 20, "Bogor", "Helpdesk")
	createData1("Daniel", 22, "Jakarta", "Technical Support")
	createData1("Aji", 29, "Cianjur", "Dokter")

	//READ DATA
	read := readData1()
	fmt.Println(read[0])

	//UPDATE DATA
	user, err := updateData1(2, "Alexander", 2, "Bogor", "Helpdesk")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Updated", user)
	}

	//DELETED DATA
	user, err = deletedData1(6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Deleted", user)
	}
	//READ DATA ALL
	read = readData1()
	fmt.Println(read)
}
