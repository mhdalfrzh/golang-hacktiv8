package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greet("paras", 23)
	greet1("paras", "bogor")

	var names = []string{"paras", "alparis"}
	var printMsg = greet2("Heii", names)

	fmt.Println(printMsg)

	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	var area2, circumference2 = calculate2(diameter)
	fmt.Println(area, circumference)
	fmt.Println(area2, circumference2)

	studentLists := print("ras", "pal", "han", "gik")
	fmt.Printf("%v", studentLists)

	numberLists := []int{1, 2, 3, 4, 5}
	result := sum(numberLists...)
	fmt.Println(result)

	profile("ras", "sidang", "sigor")
}

func greet(name string, age int8) {
	fmt.Printf("Hello there my name is %s and I'm %d years olf", name, age)
}

// kalo parameter lebih dari satu tapi tipe datanya sama, maka bisa digabung
func greet1(name, address string) {
	fmt.Println(name, address)
}

// kalo function kita return sebuah nilai, maka kita perlu tulis tipe datanya
func greet2(msg string, names []string) string {

	// fungsi join berguna untuk menggabungkan tipe data string
	var joinStr = strings.Join(names, " ")

	// sprintf akan mereturn sebuah nilai, kalo printf engga
	var result string = fmt.Sprintf("%s %s", msg, joinStr)
	return result
}

// returning multiple values
func calculate(d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return area, circumference
}

// predefined return value
func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}

// variadic function : fungsi yang dapat menerima argumen tak terbatas
func print(names ...string) []map[string]string {
	var result []map[string]string
	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

// menggunakan slice sebagai parameter function
func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

// menggabungkan parameter biasa dengan variadic, tapi variadic harus selalu di akhir
func profile(name string, favFoods ...string){
	mergeFavFoods := strings.Join(favFoods, ",")
	fmt.Println(mergeFavFoods)
}