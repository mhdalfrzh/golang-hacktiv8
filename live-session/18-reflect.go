package main

import (
	"fmt"
	"reflect"
)

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	/*
		reflect digunakan untuk inspeksi variabel, mengambil informasi dari variabel
		atau memanipulasinya. cakupan informasinya sangat luas seperti melihat struktur variabel,
		tipe, nilai pointer, dan banyak lagi.
	*/

	/*
		fungsi reflect.valueOf() memiliki parameter yang bisa menampung segala jenis tipe data. Fungsi
		tersebut mengembalikan objek dalam tipe reflect.Value yang berisikan informasi terkait variabel yang bersangkutan
	*/
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("tipe variabel: ", reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel: ", reflectValue.Int())
	}

	// jika nilai hanya diperlukan untuk ditampilak di output, bisa menggunakan .Interface()
	fmt.Println("tipe variabel: ", reflectValue.Type())
	fmt.Println("nilai variabel: ", reflectValue.Interface())

	// Informasi mengenai method bisa diakses lewat reflect, syaratnya harus bermodifier public

	/*
	s1 merupakan instance struct student. Awalnya property name berisikan "John Wick".
	lalu refleksi nilai objek tersebut diambil, method SetName juga diambil
	*/
	var s1 = &student{Name: "John Wick", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue1 = reflect.ValueOf(s1)
	var method = reflectValue1.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama: ", s1.Name)
}
