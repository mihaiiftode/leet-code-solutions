package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("ggububgvfk"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring(" "))
}

func lengthOfLongestSubstring(s string) int {
	cache := map[byte]int{}
	count := 0
	maxCount := 0
	for i := 0; i < len(s); i++ {
		if val, ok := cache[s[i]]; ok {
			if count >= maxCount {
				maxCount = count
			}
			i = val + 1 // go on next position
			count = 1
			cache = map[byte]int{s[i]: i}
		} else {
			count++
			cache[s[i]] = i
		}
	}

	if count < maxCount {
		return maxCount
	} else {
		return count
	}
}

// func lengthOfLongestSubstring(s string) int {
//     if len(s) <= 1 {
//         return len(s)
//     }
//     n := len(s)
//     max := 0
//     left, right := 0, 0
//     counts := make([]int, 128)
//     for right < n {
//         c := byte(s[right])
//         counts[c]++
//         for counts[c] > 1 {
//             z := byte(s[left])
//             counts[z]--
//             left++
//         }
//         newLen := right-left + 1
//         if newLen > max {
//             max =newLen
//         }
//         right++
//     }

//     return max
// }
