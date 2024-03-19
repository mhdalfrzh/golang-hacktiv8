package main

import (
	"fmt"
)

func main() {
	/*
		empty interface adalah tipe data yang dapat menerima segala tipe data.
		Maka dari itu, variabel dengan empty interface dapat diberikan nilai dengan
		tipe data yang berbeda-beda
	*/
	var randomValues interface{}
	_ = randomValues
	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Paras"}

	var v interface{}
	v = 20

	/* terjadi error karena hanya bisa melakukan perkalian terhadap int yang asli,
	   sedangkan tipe data v adalah empty interface yang diberikan nilai int
	*/
	v = v * 9

	// untuk mengatasi error tersebut, kita bisa melakukan type assertion
	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	/*
		ketika sebuah map dideklarasikan dan valuenya diberikan empty interface, maka map
		tersebut dapat memiliki value dengan tipe data berbeda.
	*/
	rs := []interface{}{1, "Paras", true, 2, "nopal", false}
	rn := map[string]interface{}{
		"Name":   "Paras",
		"Status": true,
		"Age":    23,
	}
	_, _ = rs, rn

}
