package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman
type Teman struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Function untuk mendapatkan data teman berdasarkan absen
func getDataByAbsen(absen int) Teman {
	dataTeman := map[int]Teman{
		1: {1, "Paras", "Jakarta", "Software Engineer", "Suka dengan performa dan kemudahan penggunaan Go"},
		2: {2, "Nopal", "Bandung", "Data Scientist", "Ingin mempelajari bahasa pemrograman Go untuk proyek data science"},
		3: {3, "Reyhan", "Bali", "Backend Developer", "Belajar untuk membuat sebuah backend service menggunakan Go"},
	}

	// Mengembalikan data teman berdasarkan absen
	return dataTeman[absen]
}

func main() {
	// Mendapatkan argument dari terminal
	args := os.Args

	// Mendapatkan nomor absen dari argumen
	absen := args[1]

	// Konversi nomor absen dari string ke integer
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	// Mendapatkan data teman berdasarkan absen
	teman := getDataByAbsen(absenInt)

	// Menampilkan data teman
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
