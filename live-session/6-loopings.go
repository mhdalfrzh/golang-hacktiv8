package main

import "fmt"

func main() {

	// first way
	for i := 0; i < 3; i++ {
		fmt.Println("First", i)
	}

	// second way
	var j = 0
	for j < 3 {
		fmt.Println("Second", j)
		j++
	}

	// third way
	var k = 0
	for {
		fmt.Println("Third", k)
		k++
		if k == 3 {
			break
		}
	}

	// break and continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("brecon", i)
	}

	// nested looping
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// label : melakukan perulangan tertentu
	outerLoop:
		for i := 0; i < 3; i++{
			fmt.Println("Perulangan ke - ", i + 1)
			for j := 0; j < 3; j++{
				if i == 2{
					break outerLoop
				}
				fmt.Print(j, " ")
			}
			fmt.Print("\n")
		}
}
