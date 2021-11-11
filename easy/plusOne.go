package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	digit := 1
	for j := len(digits) - 1; j >= 0; j-- {
		if digits[j]+digit > 9 {
			digits[j] = 0
		} else {
			digits[j] = digits[j] + digit
			return digits
		}
	}

	if digit == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
