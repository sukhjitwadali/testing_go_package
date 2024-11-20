package main

import (
	"fmt"
	"strings"
)

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

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func FindMax(nums []int) int {
	if len(nums) == 0 {
		panic("Cannot find max of an empty slice")
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func Sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func IsPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	// Check if the string is equal to its reverse
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println("CountVowels: ", CountVowels("bcdfg"))
	fmt.Println("Graetest Common Divisor: ", GCD(17, 18))
	fmt.Println("Finding maximum number: ", FindMax([]int{19, 3, 2, 11, 4}))
	fmt.Println("Finding Sum: ", Sum([]int{5, 4, 2, 11, 4}))
	fmt.Println("Checking Palindrome for 'madam':", IsPalindrome("madam"))

}
