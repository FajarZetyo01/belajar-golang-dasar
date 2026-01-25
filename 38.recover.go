package main

import "fmt"

func endApp1() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApplication1(error bool) {
	defer endApp1()
	if error {
		panic("Ups Error")
	}
}

/*✅ 3. RECOVER
Recover = memulihkan program dari panic agar tidak langsung mati.
Ibaratnya:
“Pemadam kebakaran datang, jadi api dipadamkan dan aktivitas bisa jalan lagi.”
Recover hanya jalan di dalam fungsi yang dipanggil dengan defer
*/

func main() {
	runApplication1(true)

}
