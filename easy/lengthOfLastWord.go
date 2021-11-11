package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}

	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == byte(' ') {
			return length
		} else {
			length = length + 1
		}
	}

	return length
}
