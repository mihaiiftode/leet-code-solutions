package main

func main() {
	characterReplacement("ABBB", 2)
}

func characterReplacement(s string, k int) int {
	longest := 0
	left := 0
	right := 0
	maxCharCount := 0
	frequency := make(map[byte]int)

	for ; right < len(s); right = right + 1 {
		rightChar := s[right]
		frequency[rightChar]++

		maxCharCount = max(maxCharCount, frequency[rightChar])

		if right-left+1-maxCharCount > k {
			leftChar := s[left]
			frequency[leftChar]++
			left++
		}

		longest = max(right-starleft+1, longest)

	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
