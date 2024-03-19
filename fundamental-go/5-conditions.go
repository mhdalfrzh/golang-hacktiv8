package main

import "fmt"

func main() {
	var year = 2024

	// kita bisa bikin variable sementara yang hanya bisa diakses di blok if
	if age := year - 2000; age < 17 {
		fmt.Println("kamu udah punya ktp")
	} else {
		fmt.Println("kamu belom punya ktp")
	}

	// switch
	var score = 10
	switch score {
	case 10:
		fmt.Println("mantap")
	case 9:
		fmt.Println("awesome")
	case 8:
		fmt.Println("anjay")
	default:
		fmt.Println("not bad lah")
	}

	// switch dengan operator perbandingan
	switch {
	case score == 10:
		fmt.Println("perfect")
	case (score < 10) && score > 3:
		fmt.Println("mayanlah")
	default:
		fmt.Println("belajar lagi yah")
	}

	// faltrough keyword : melanjutkan pengecekan walaupun case terpenuhi
	switch {
	case score == 10:
		fmt.Println("perfect")
		fallthrough
	case (score < 10) && (score > 3):
		fmt.Println("mayanlahh")
	case score > 5:
		fmt.Println("ini ikut diprint nih")
	default:
		fmt.Println("belajar lagi yah")
	}
	println(score < 10 && score > 3)

	// nested conditions
	var skor = 7
	if skor > 7 {
		switch score {
		case 10:
			fmt.Println("ntaps")
		default:
			fmt.Println("nais")
		}
	} else {
		if score == 5 {
			fmt.Println("blok")
		} else {
			fmt.Println("belajar yang rajin")
		}
	}
}
