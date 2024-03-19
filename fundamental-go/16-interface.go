package main

import (
	"fmt"
	"math"
)

/*
Interface adalah sebuah tipe data pada bahasa Go yang merupakan kumpulan dari definisi satu atau lebih method. Kita
dapat menggunakan tipe data interface jika kita telah mengimplementasikan method-method yang telah didefinisikan oleh
interface tersebut melalui tipe data lainnya.
*/

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func print(t string, s shape) {
	fmt.Println(t, s.area())
	fmt.Println(t, s.perimeter())
}

// Ketika struct circle menambahkan satu method selain dari interface shape, maka variabel c1
// tidak dapat menggunakan method tersebut
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	// variabel c1 dan r1 dapat menampung struct circle dan rectangle karena telah mengimplementasikan
	// seluruh method yang didefinisikan oleh interface shape
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	// Kedua variabel ini akan memiliki 2 tipe data yaitu struct dan shape, hal inilah yang disebut polymorphism
	// dengan implementasi interface, maka suatu variabel akan mempunyai 2 tipe data
	fmt.Printf("Type of c1 %T", c1)
	fmt.Printf("Type of c1 %T", r1)

	fmt.Println(c1.area())
	fmt.Println(c1.perimeter())
	fmt.Println(r1.area())
	fmt.Println(r1.perimeter())

	// karena variabel c1 dan r1 tipe datanya interface, maka bisa jadi argumen kedua fungsi print
	print("Rectangle", c1)
	print("Circle", r1)

	// Jika pada fungsi main dipanggil maka akan terjadi error
	// c1.volume()

	// Untuk mengatasi error tersebut, diperlukan type assertion yaitu mengembalikan tipe data interface
	// ke tipe data aslinya

	// mengembalikan variable c1 dari interface ke struct circle
	c1.(circle).volume()

	// cara memeriksa apakah type assertion berhasil atau tidak
	value, ok := c1.(circle)
	if ok == true{
		fmt.Println(value)
		fmt.Println(value.volume())
	}
}
