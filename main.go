package main

import (
	"fmt"
)

func main() {
	str := "selamat malam"

	charCount := make(map[string]int)

	for _, char := range str {
		fmt.Println(string(char)) 
		charCount[string(char)]++
	}

	// Mencetak hasil map
	fmt.Println(charCount)
}
