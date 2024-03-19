package main

import "fmt"

func main(){
	var first = 89
	var second = -12

	// menentukan tipe data sesuai ukuran
	var first1 uint8 = 89
	var second1 int8 = -12

	// float
	var decimal float32 = 3.63

	// bool
	var condition bool = true

	// string
	var message string = "Halo"
	var message1 string = `halo halo bandung`

	fmt.Printf("tipe data first adalah %T\n", first)
	fmt.Printf("tipe data second adalah %T\n", second)
	fmt.Printf("tipe data first adalah %T\n", first1)
	fmt.Printf("tipe data second adalah %T\n", second1)
	fmt.Printf("is it permitted? %t\n", condition)
	fmt.Println(message)
	fmt.Println(message1)

	// default %.f adalah 6 angka di belakang koma
	fmt.Printf("decimal number : %.3f", decimal)
}