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
