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

func main() {
	fmt.Println("CountVowels: ", CountVowels("bcdfg"))
	fmt.Println("Graetest Common Divisor: ", GCD(17, 18))
	fmt.Println("Finding maximum number: ", FindMax([]int{19, 3, 2, 11, 4}))

}
