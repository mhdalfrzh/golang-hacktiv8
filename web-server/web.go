package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main() {
	http.HandleFunc("/", greet)

	// menjalankan server aplikasi
	http.ListenAndServe(PORT, nil)
}

// http.ResoponseWriter digunakan untuk mengirim response balik ke si yang mengirim request
func greet(w http.ResponseWriter, r *http.Request){
	msg := "Hello World"
	fmt.Fprint(w, msg)
}