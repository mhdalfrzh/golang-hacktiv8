package main

import "fmt"

// function yang akan selalu dieksekusi sebelum function main
func init() {
	fmt.Println("Halo ini dari function init")
}