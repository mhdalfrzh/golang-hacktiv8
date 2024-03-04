package main

import "fmt"

func main() {

	// slice mirip array tapi ga punya fixed length
	var slice = []string{"apple", "banana", "mango"}
	fmt.Println(slice)

	// make function, argumen pertama tipe data dan kedua panjang slice
	var fruits = make([]string, 3)
	fmt.Println(fruits)

	// append function : untuk menambahkan element pada slice
	var hewan = make([]string, 3)
	hewan = append(hewan, "ayam", "kambing", "bebek")
	fmt.Printf("%#v", hewan)

}
