package main

import "fmt"

func CountVowels(s string) int {
	count := 0
	vowels := "aeiouAEIOU"
	for _, char := range s {
		if contains(vowels, char) {
			count++
		}
	}
	return count
}

func contains(vowels string, char rune) bool {
	for _, v := range vowels {
		if v == char {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("CountVowels: ", CountVowels("bcdfg"))

}
