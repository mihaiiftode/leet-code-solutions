package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x = x / 10
	}

	return x == reversed || x == reversed/10

}
