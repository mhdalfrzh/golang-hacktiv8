package main

import "fmt"

func main() {
	// aliase adalah nama alternative dari tipe data yang sudah ada
	// tipe data byte merupakan aliase dari uint8, rune aliase dari uint32
	var a uint8 = 15
	var b byte
	b = a // no error karena byte tipe data yang sama dengan uint8
	_ = b

	type second = uint
	var hour second = 3600
	fmt.Printf("hour type %T\n", hour)
}
