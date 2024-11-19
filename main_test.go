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
