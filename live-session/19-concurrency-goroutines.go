package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		concurrency adalah mengeksekusi sebuah proses secara independen atau berhadapan
		dengan lebih dari satu tugas dalam waktu yang sama. Berbeda dengan parallelism, kalo paralelism
		itu mengerjakan tugas yang banyak secara bersamaan. Sedangkan pada concurrency kita tidak tau siapa
		yang akan menyelesaikan tugasnya terlebih dahulu
	*/

	/*
		goroutine adalah thread pada go untuk melakukan concurrency. Goroutine memiliki ukuran
		thread yang jauh lebih ringan. Goroutine bersifat asynchronous sehingga tidak saling tunggu dengan
		Goroutine lainnya.
	*/
	go goroutine()

	fmt.Println("main execution started")
	go firstProcess(8)
	secondProcess(8)

	// mengetahui jumlah goroutine yang sedang berjalan
	fmt.Println("No of Goroutines: ", runtime.NumGoroutine())

	//Time.Sleep(time.Second * 2)
	fmt.Println("main execution ended")

}

func goroutine() {
	fmt.Println("Hello")
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second process func ended")
}

/*
output program diatas adalah berikut :

main execution started
Hello
second process func started
j= 1
j= 2
j= 3
j= 4
j= 5
j= 6
j= 7
j= 8
second process func ended
No of Goroutines:  2
main execution ended

Penjelasan :
Function firstProcess tidak menampilkan hasilnya karena setiap Goroutine bekerja secara async dan tidak saling
tunggu antar Goroutine satu dengan Goroutine lainnya

Terdapat 2 jumlah Goroutine yang sedang berjalan padahal kita hanya menjalankan satu function yang dijadikan Goroutine
Ini terjadi karena function main juga sebuah Goroutine sehingga tidak akan menunggu Goroutine lainnya untuk dieksekusi.

Gouroutine ketika dijalankan akan membutuhkan waktu lebih lama untuk memulai dibandingkan function biasa. Maka kita perlu
suatu function yang akan menahan function main untuk langsung menyelesaikan tugasnya.
*/
