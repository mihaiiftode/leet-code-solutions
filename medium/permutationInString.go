package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "fasdfbarere"))
	fmt.Println(checkInclusion("ab", "ab"))
	fmt.Println(checkInclusion("ab", "a"))
	fmt.Println(checkInclusion("hello", "ooolleoooleh"))
}

func checkInclusion(s1 string, s2 string) bool {
	cache := make([]int, 128)
	tempCount := make([]int, 128)

	for i := 0; i < len(s1); i++ {
		if cache[s1[i]] > 0 {
			cache[s1[i]]++
		} else {
			cache[s1[i]] = 1
		}
	}

	count := 0
	for i := 0; i < len(s2); i++ {
		count = 0
		copy(tempCount, cache)
		j := 0
		for ; j < len(s1) && i+j < len(s2); j++ {
			if cache[s2[i+j]] > 0 {
				tempCount[s2[i+j]]--
				count++
			}
		}
		if count == len(s1) && match(tempCount) {
			return true
		}
	}

	return false

}

func match(vec []int) bool {
	for i := 90; i <= 126; i++ {
		if vec[i] < 0 {
			return false
		}
	}

	return true
}
