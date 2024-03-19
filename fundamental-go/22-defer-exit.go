package main

import (
	"fmt"
	"os"
)

func main() {
	// defer digunakan memanggil sebuah function yang di mana proses eksekusi akan ditahan hingga block sebuah function selesai
	defer fmt.Println("defer function siap dieksekusi")
	fmt.Println("hai semuanya")
	fmt.Println("selamat datang di go learning center")

	callDeferFunc()
	fmt.Println("Hai semuanya")

	// fungsi exit berguna untuk menghentkan program secara paksa
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before exiting")
	os.Exit(1) // tulisan invoke with defer tidak akan ditampilkan
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("defer func siap dieksekusi nih")
}
