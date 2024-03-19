package main

import "fmt"

func main() {
	// deklarasi variabel dengan tipe data
	var name string = "Paras"
	var age int = 23

	// kita bisa reassign variable yang udah dibuat
	name = "alfarizhi"
	age = 23

	// inisialisasi variabel tanpa tipe data / type inference
	var jenis = "laki"
	var kos = "kenzie"

	// short declaration
	hobi := "badminton"

	// multiple variable declarations
	var student1, student2, student3 = "reyhan", "nopal", "irgi"

	// underscore variabel, untuk menghindari error jika ada variabel yang ga kepake
	var first string
	_ = first

	fmt.Println("Nama saya adalah", name)
	fmt.Println("Umur saya adalah", age)
	fmt.Println("Jenis kelamin", jenis)
	fmt.Println("Tinggal di kos", kos)
	fmt.Println("Hobi saya adalah", hobi)
	fmt.Println(student1, student2, student3)
	fmt.Println(first)

	// Printf untuk mengetahui tipe data suatu variabel
	fmt.Printf("Tipe data variable ini adalah %T", hobi)
}
