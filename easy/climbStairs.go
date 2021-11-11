package main

import "fmt"

func main() {
	fmt.Println(climbStairs(10))
}

func climbStairs(n int) int {
	visited := map[int]int{}
	visited[0] = 1
	visited[1] = 1

	for i := 2; i <= n; i++ {
		visited[i] = visited[i-1] + visited[i-2]

	}
	return visited[n]
}

// func climbStairs(n int) int {
// 	return helper(n, make(map[int]int))
// }

// func helper(n int, cache map[int]int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}

// 	if val, ok := cache[n]; ok {
// 		return val
// 	} else {
// 		cache[n] = helper(n-1, cache) + helper(n-2, cache)
// 	}

// 	return cache[n]
// }
