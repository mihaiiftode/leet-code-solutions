package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}

func letterCasePermutation(s string) []string {
	result := make([]string, 0)
	backtrack([]rune(s), 0, &result)

	return result
}

func backtrack(s []rune, start int, result *[]string) {
	for i := start; i < len(s); i++ {
		if unicode.IsLetter(s[i]) {
			backtrack(s, i+1, result)
			if unicode.IsUpper(s[i]) {
				s[i] = unicode.ToLower(s[i])
			} else {
				s[i] = unicode.ToUpper(s[i])
			}

		}
	}
	*result = append(*result, string(s))

}

// func letterCasePermutation(s string) []string {
//     result := make([]string, 0)
//     helper(s, "", &result)
//     return result
// }

// func helper(data string, slate string, result *[]string) {
//     if len(data) == 0 {
//         *result = append(*result, slate)
//         return
//     }

//     crune := rune(data[0])

//     if unicode.IsLetter(crune) {
//         var tmp rune

//         if unicode.IsUpper(crune) {
//             tmp = unicode.ToLower(crune)
//         } else {
//             tmp = unicode.ToUpper(crune)
//         }

//         helper(data[1:], slate + string(tmp), result)
//     }

//     helper(data[1:], slate + string(crune), result)
// }
