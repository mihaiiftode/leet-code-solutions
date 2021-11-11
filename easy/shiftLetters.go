package main

import "fmt"

func main() {
	// fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
	fmt.Println(shiftingLetters("bad", []int{10, 20, 30}))
}

func shiftingLetters(s string, shifts []int) string {
	str := []byte(s)
	sum := 0
	for index := len(s) - 1; index >= 0; index-- {
		sum = (sum + shifts[index]) % 26
		cur := int(str[index])
		str[index] = byte((cur-97+sum)%26 + 97)
	}

	return string(str)
}

// func shiftingLetters(s string, shifts []int) string {
// 	acum := 0
// 	ans := make([]byte, len(s))
// 	for i := len(s) - 1; i >= 0; i-- {
// 		acum = (acum + shifts[i]) % 26
// 		ans[i] = 97 + (s[i]+byte(acum)-97)%26
// 	}
// 	return string(ans)
// }
