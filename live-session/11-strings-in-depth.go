package main

import "fmt"

func main() {

	// loop over string byte by byte
	city := "jakarta"
	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte:%d\n", i, city[i])
	}

	var city2 []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println(string(city2))
}
