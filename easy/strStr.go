package main

import (
	"fmt"
)

func main() {
	// fmt.Println(strStr("hello", "ll"))
	// fmt.Println(strStr("aaaaa", "bba"))
	// fmt.Println(strStr("mississippi", "issip"))
	fmt.Println(strStr("mississippi", "pi"))
}

func strStr(haystack string, needle string) int {

	haystackLength := len(haystack)
	needleLength := len(needle)

	if needleLength == 0 {
		return 0
	}

	if haystackLength < needleLength {
		return -1
	}

	for i := 0; i <= haystackLength-needleLength; i++ {
		if haystack[i:i+needleLength] == needle {
			return i
		}
	}

	return -1

}
