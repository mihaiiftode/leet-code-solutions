package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	var result [][]int
	buffer := make([]int, 0, k)
	backtrack(&result, buffer, k, n, 1)
	return result
}

func backtrack(result *[][]int, buffer []int, k, n, start int) {
	if k == 0 {
		cp := make([]int, len(buffer))
		copy(cp, buffer)
		*result = append(*result, cp)
		return
	}

	for i := start; i <= n; i++ {
		backtrack(result, append(buffer, i), k-1, n, i+1)
	}
}
