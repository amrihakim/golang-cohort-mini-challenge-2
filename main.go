package main

import (
	"fmt"
)

func main() {
	str := "selamat malam, golang cohort"

	charCount := make(map[string]int)

	for _, char := range str {
		fmt.Println(string(char)) 
		charCount[string(char)]++
	}

	fmt.Println(charCount)
}
