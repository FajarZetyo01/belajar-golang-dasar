package main

import (
	"errors"
	"fmt"
)

// 1️⃣ Struktur data (MODEL)
type User struct {
	ID   int
	Name string
	Age  int
}

// 2️⃣ Storage sederhana (DATABASE PALSU 😄)
var users []User //SLICE GLOBAL
var lastID int = 0

// 3️⃣ CREATE (Tambah data)
func createUser(name string, age int) User {
	lastID++
	user := User{
		ID:   lastID,
		Name: name,
		Age:  age,
	}
	users = append(users, user) //user di tambahkan ke tabel users diatas
	return user
}

// 4️⃣ READ (Ambil semua data)
func readUser() []User {
	return users
}

// 5️⃣ UPDATE (Ubah data user berdasarkan ID)
func updateUser(id int, name string, age int) (User, error) {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = name
			users[i].Age = age
			return users[i], nil
		}
	}
	return User{}, errors.New("User tidak ditemukan")
}

// 6️⃣ DELETE (Hapus user berdasarkan ID)
func deleteUser(id int) (User, error) {
	for i, user := range users {
		if user.ID == id {
			deletedUser := user // simpan dulu
			users = append(users[:i], users[i+1:]...)
			return deletedUser, nil
		}
	}
	return User{}, errors.New("User tidak ditemukan")
}
func main() {

	// CREATE
	createUser("Fajar", 25)
	createUser("Budi", 30)
	createUser("Daniel", 30)
	createUser("Joko", 20)

	// READ
	fmt.Println("Data awal:")
	fmt.Println(readUser())

	// UPDATE
	user, err := updateUser(1, "Fajar Alexander", 26)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Updated:", user)
	}

	// DELETE
	user, err = deleteUser(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Deleted:", user.Name)
	}

	// READ lagi
	fmt.Println("DATA AKHIR:")
	fmt.Println("=======================================================")
	for _, user := range readUser() {
		fmt.Println("ID:", user.ID, "Name:", user.Name, "Age:", user.Age)
	}
	fmt.Println("=======================================================")
}
