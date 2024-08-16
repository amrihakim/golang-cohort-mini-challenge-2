package main

import (
	"fmt"
)

func main() {
	// Inisialisasi string
	str := "selamat malam"

	// Inisialisasi map untuk menghitung frekuensi kemunculan setiap karakter
	charCount := make(map[string]int)

	// Loop untuk memecah string menjadi karakter dan menghitung frekuensinya
	for _, char := range str {
		fmt.Println(string(char)) // Mencetak karakter per baris
		charCount[string(char)]++ // Menambahkan frekuensi karakter ke dalam map
	}

	// Mencetak hasil map
	fmt.Println(charCount)
}
