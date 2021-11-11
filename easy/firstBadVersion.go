package main

import "fmt"

func main() {
	// fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
	// fmt.Println(firstBadVersion(5))
	fmt.Println(firstBadVersion(7))
}

func firstBadVersion(n int) int {

	left := 0
	right := n

	for left < right {
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func isBadVersion(version int) bool {
	return version == 7
}
