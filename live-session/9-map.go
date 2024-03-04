package main

import "fmt"

func main() {
	// map pada Go mirip seperti object pada javascript yaitu berisi pasangan key dan value
	var person map[string]string // artinya key dari map harus string dan value dari map juga string
	person = map[string]string{}
	person["name"] = "ras"
	person["age"] = "23"
	person["address"] = "jalan sudirman"
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])

	/* kasi key dan value langsung ketika inisialisasi
	var person = map[string]string{
		"name": "ras"
		"age": "23"
		"address": "jalan sudirman"
	}
	*/

	// looping map with range
	for key, value := range person {
		fmt.Println(key, value)
	}

	// menghapus value dari map
	delete(person, "age")
	fmt.Println(person)

	// detecting a value
	// value untuk mengembalikan nilainya, exist untuk cek apakah ada atau engga (bool)
	value, exist := person["agee"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("gaada")
	}

	// combining slice with map
	var people = []map[string]string{
		{"name": "lewa", "age": "23"},
		{"name": "pedri", "age": "23"},
	}

	for _, person := range people{
		fmt.Println(person["name"], person["age"])
	}

}
