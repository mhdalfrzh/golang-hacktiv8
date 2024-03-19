package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		Waitgroup adalah struct dari package sync untuk sinkronisasi go routine. Pada materi sebelumnya kita perlu
		menahan function main agar go routine menyelesaikan prosesnya menggunakan function sleep dari package time.
		Cara ini bukan merupakan cara yang baik menunggu go routine menyelesaikan prosesnya.
	*/
	fruits := []string{"Apple", "Manggo", "Durian", "Rambutan"}
	var wg sync.WaitGroup
	for index, fruit := range fruits {
		// method add untuk menambah counter dari waitgroup untuk memberitahu jumlah go routine yang harus ditunggu
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}
	// menunggu seluruh go routine menyelesaikan prosesnya sehingga method ini akan menahan function main hingga seluruh proses go routine selesai
	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	// Untuk memberitahu waitgroup terkait go routine yang telah menyelesaikan prosesnya
	wg.Done()
}
