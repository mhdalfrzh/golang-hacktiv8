package main

import "fmt"

func main(){

	// const adalah variabel yang gabisa diubah nilainya, harus inisialisasi dari awal
	const fullName string = "Parasi"

	// Arithmetic operators
	var value = 2 + 2 * 3

	// Comparison operators
	var first bool = 2 < 3
	var second bool = "joey" == "Joey"

	// Logical operators
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right

	fmt.Printf("Hello %s\n", fullName)
	fmt.Println(value)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(wrongAndRight)
}