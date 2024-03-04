package main

import (
	"fmt"
	"strings"
)

func main() {

	/*
		closure adalah fungsi tanpa nama yang dapat disimpan sebagai sebuah variabel
		dan dijadikan sebagai nilai return sebuah function
	*/

	// declare closure in variable
	var evenNumbers = func(numbers ...int) []int {
		var result []int
		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(evenNumbers(numbers...))

	// IIFE : closure yang langsung dieksekusi ketika pertama kali dideklarasikan
	var isPalindrome = func(str string) bool {
		var temp string = ""
		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)

	var studentLists = []string{"ras", "pal", "gik", "han"}
	var find = findStudent(studentLists)
	fmt.Println(find("han"))

	var numbers1 = []int{1, 2, 3, 4, 5}
	var find1 = findOddNumbers(numbers1, func(number int) bool {
		return number%2 != 0
	})
	fmt.Println(find1)

}

// closure as a return value
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int
		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s gaada", s)
		}
		return fmt.Sprintf("ada %s di %d", s, position+1)
	}
}

// callback : closure yang jadi parameter di function
func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

/*
kita bisa pake alias untuk mempersingkat parameter callbacknya
type isOddNum = func(int) bool

func findOddNumbers(numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
*/
