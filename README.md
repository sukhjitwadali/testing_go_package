# Go Testing and Benchmark Suite

This repository contains test and benchmark implementations for a variety of utility functions in Go. The tests cover common operations like string analysis, mathematical calculations, array manipulations, and more. Benchmarks are included to evaluate the performance of these functions.

## Functions Covered

### 1. `CountVowels`
Counts the number of vowels in a given string, case-insensitive.

- **Tests**:
  - Handles empty strings.
  - Processes strings with no vowels, all vowels, and mixed-case vowels.
  - Includes non-alphabet characters in strings.
- **Benchmark**:
  - Evaluates performance with a long lorem ipsum text.

### 2. `GCD`
Finds the greatest common divisor (GCD) of two integers using the Euclidean algorithm.

- **Tests**:
  - Handles cases where both numbers are zero, one number is zero, or the numbers are negative.
  - Tests for coprime numbers and numbers with a common divisor.
- **Benchmark**:
  - Evaluates performance with large integers.

### 3. `FindMax`
Finds the maximum value in a slice of integers.

- **Tests**:
  - Handles single-element slices, positive numbers, negative numbers, and mixed slices.
  - Includes a test for expected panic on an empty slice.
- **Benchmark**:
  - Measures performance with a large slice of integers.

### 4. `Sum`
Calculates the sum of all elements in a slice of integers.

- **Tests**:
  - Covers empty slices, single-element slices, and slices with positive or negative numbers.
- **Benchmark**:
  - Evaluates performance with a large slice of integers.

### 5. `IsPalindrome`
Determines whether a given string is a palindrome, ignoring case and spaces.

- **Tests**:
  - Handles empty strings, single-character strings, and typical palindromes.
  - Tests for case insensitivity and spaces in the input.
- **Benchmark**:
  - Evaluates performance with a long palindrome phrase.

## How to Run Tests

go test
go test -v
