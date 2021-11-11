package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("test test east"))
}

func reverseWords(s string) string {
	splitString := strings.Split(s, " ")

	for i, v := range splitString {
		splitString[i] = reverseString([]byte(v))
	}

	return strings.Join(splitString, " ")

}

func reverseString(s []byte) string {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}

// func reverseString(s []byte) {
// 	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }
