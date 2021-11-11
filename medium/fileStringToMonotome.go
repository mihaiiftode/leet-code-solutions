package main

import "fmt"

func main() {
	fmt.Println(minFlipsMonoIncr("00011000"))
}

func minFlipsMonoIncr(s string) int {
	count, flip := 0, 0
	for _, v := range s {
		if v == '0' {
			flip = flip + 1
		} else {
			count = count + 1
		}

		flip = min(flip, count)
	}

	return flip
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
