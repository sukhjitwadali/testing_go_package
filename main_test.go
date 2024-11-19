package main

import "testing"

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result int
	}{
		{"Empty String", "", 0},
		{"No Vowels", "bcdfg", 0},
		{"All Vowels", "aeiou", 5},
		{"Mixed Case Vowels", "AeIoU", 5},
		{"With Non-Alphabet", "Hello, World!", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountVowels(tt.input)
			if got != tt.result {
				t.Errorf("CountVowels(%q) = %d; want %d", tt.input, got, tt.result)
			}
		})
	}
}
func BenchmarkCountVowels(b *testing.B) {
	input := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum varius."
	for i := 0; i < b.N; i++ {
		CountVowels(input)
	}
}
func TestGCD(t *testing.T) {
	tests := []struct {
		name   string
		a, b   int
		result int
	}{
		{"Both Zero", 0, 0, 0},
		{"One Zero", 0, 10, 10},
		{"Equal Numbers", 5, 5, 5},
		{"Coprime Numbers", 17, 13, 1},
		{"Common Divisor", 56, 98, 14},
		{"Negative Numbers", -48, 18, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GCD(tt.a, tt.b)
			if got != tt.result {
				t.Errorf("GCD(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.result)
			}
		})
	}
}
func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(98765432, 12345678)
	}
}

func TestFindMax(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		result int
	}{
		{"Single Element", []int{5}, 5},
		{"All Positive Numbers", []int{1, 2, 3, 4, 5}, 5},
		{"Mixed Positive and Negative", []int{-1, 0, 3, 7, -5}, 7},
		{"All Negative Numbers", []int{-9, -7, -5, -3}, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMax(tt.input)
			if got != tt.result {
				t.Errorf("FindMax(%v) = %d; want %d", tt.input, got, tt.result)
			}
		})
	}
}

func TestFindMaxPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty slice")
		}
	}()
	FindMax([]int{})
}
func BenchmarkFindMax(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindMax(nums)
	}
}
