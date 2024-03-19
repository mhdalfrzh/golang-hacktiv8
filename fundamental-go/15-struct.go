package main

import "fmt"

/*
	bahasa go ga menyediakan tipe data class tapi go punya struct
	struct adalah tipe data berupa kumpulan dari berbagai macam property
*/
type Employee struct {
	name     string
	age      int
	division string
}

type Person struct {
	name string
	age  int
}

// struct dapat mengandung struct lainnya
type Animal struct {
	species string
	habitat string
}

type Pet struct {
	animal Animal
	name   string
}

func main() {
	// simpan struct ke sebuah variabel
	var employee Employee
	employee.name = "ras"
	employee.age = 23
	employee.division = "it"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	// cara kedua
	var employee2 = Employee{name: "pal", age: 22, division: "it"}

	fmt.Println(employee)
	fmt.Println(employee2)

	// pointer to a struct
	var employee3 *Employee = &employee
	fmt.Println("Employee 1", employee.name)
	fmt.Println("Employee 3", employee3.name)

	employee3.name = "Gik"
	fmt.Println("Employee 1", employee.name)
	fmt.Println("Employee 3", employee3.name)

	var pet = Pet{}
	pet.animal.species = "ras"
	pet.animal.habitat = "rumah"
	pet.name = "oyen"
	fmt.Println(pet)

	// anonymous struct : langsung dideklarasikan sebagai variabel
	var barca = struct {
		name   string
		nopung int
	}{}
	barca.name = "lewandowski"
	barca.nopung = 9
	fmt.Println(barca)

	// slice of struct
	var person = []Person{
		{name: "ras", age: 23},
		{name: "pal", age: 23},
	}

	for _, v := range person {
		fmt.Println(v)
	}

	// slice of anonymous struct
	var employee5 = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "ras", age: 23}, division: "it"},
		{person: Person{name: "gik", age: 23}, division: "it"},
	}

	for _, v := range employee5 {
		fmt.Println(v)
	}
}
