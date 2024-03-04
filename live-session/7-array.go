package main

import "fmt"

func main() {

	// cara deklarasi array pertama
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	// cara deklarasi array kedua
	var strings = [3]string{"ras", "pal", "gik"}

	// %#v untuk melihat isi panjang array dan tipe datanya
	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)
	fmt.Println(strings)

	// modify element through index
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fmt.Println(fruits)

	// range loop
	for i, v := range fruits{
		fmt.Println(i, v)
	}

	// cara kedua
	for i := 0; i < len(fruits); i++{
		fmt.Println(fruits[i])
	}

	// multidimensional array
	balances := [2][3]int{{5,6,7}, {8,9,10}}
	for _, arr := range balances{
		for _, value := range arr {
			fmt.Printf("%d", value)
		}
		fmt.Println()
	}
}
