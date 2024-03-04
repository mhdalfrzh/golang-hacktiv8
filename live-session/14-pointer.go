package main

import "fmt"

func main() {

	/*
		pointer adalah variabel untuk menyimpan alamat memori variabel lainnya
		sebagai contoh variabel int bernilai 4 maka pointer adalah alamat memori
		di mana nilai 4 itu disimpan

		Variabel yang memiliki alamat memori sama, saling berhubungan satu sama lain dan nilainya
		pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain
		yang alamat memorinya sama yaitu nilainya juga ikut berubah.
	*/

	var first int = 4

	// variabel pointer yang mengandung alamat memori first
	var second *int = &first

	fmt.Println("value first", first)
	fmt.Println("memory address", &first)

	// tampilkan nilai asli pada variabel second pake * karena variabel second menyimpan alamat memori bukan nilai
	fmt.Println("value second", *second)
	fmt.Println("memory address", second)

	var firstPerson string = "ras"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson value", firstPerson)
	fmt.Println("firstPerson memory address", &firstPerson)
	fmt.Println("secondNumber value", *secondPerson)
	fmt.Println("secondNumber memory address", secondPerson)

	// karena secondPerson mengandung alamat firstPerson, maka ketika nilai secon diganti first akan ikut terganti juga
	*secondPerson = "pal"
	fmt.Println("firstPerson value", firstPerson)
	fmt.Println("firstPerson memory address", &firstPerson)
	fmt.Println("secondNumber value", *secondPerson)
	fmt.Println("secondNumber memory address", secondPerson)

	var a int = 10
	fmt.Println("Before", a)

	changeValue(&a)
	fmt.Println("After", a)
}

// pointer as a parameter
func changeValue(number *int){
	*number = 20
}
