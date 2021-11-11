package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(addStrings("1", "9"))
}

func addStrings(num1 string, num2 string) string {
	if num1 == "0" && num2 == "0" {
		return "0"
	}

	if len(num1) < 1 && len(num1) > 100000 && len(num2) < 1 && len(num2) < 100000 {
		return "0"
	}

	var sum string
	var carry int64
	var second int64

	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	for i, j := len(num1)-1, len(num2)-1; i >= 0; i, j = i-1, j-1 {
		first, _ := strconv.ParseInt(string(num1[i]), 10, 64)
		if j >= 0 {
			second, _ = strconv.ParseInt(string(num2[j]), 10, 64)
		} else {
			second = 0
		}

		digitSum := first + second + carry
		carry = digitSum / 10
		sum = strconv.FormatInt(digitSum%10, 10) + sum
	}

	if carry > 0 {
		sum = strconv.FormatInt(carry, 10) + sum
	}

	return sum
}
