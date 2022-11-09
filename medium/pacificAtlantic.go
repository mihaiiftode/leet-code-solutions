package main

import "fmt"

func main() {
	fmt.Println(pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
}

func pacificAtlantic(heights [][]int) [][]int {
	n, m := len(heights), len(heights[0])
	canReachP := make([][]bool, len(heights))
	canReachA := make([][]bool, len(heights))

	for i := 0; i < len(heights); i++ {
		canReachP[i] = make([]bool, len(heights[0]))
		canReachA[i] = make([]bool, len(heights[0]))
	}

	for i := 0; i < n; i++ {
		dfs(heights, &canReachP, 0, i, 0, n, m)
		dfs(heights, &canReachA, 0, i, m-1, n, m)
	}

	for j := 0; j < m; j++ {
		dfs(heights, &canReachP, 0, 0, j, n, m)
		dfs(heights, &canReachA, 0, n-1, j, n, m)
	}

	results := [][]int{}

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if canReachA[i][j] && canReachP[i][j] {
				results = append(results, []int{i, j})
			}
		}
	}

	return results
}

func dfs(heights [][]int, visited *[][]bool, height, i, j, n, m int) {
	if i >= 0 && i < n && j >= 0 && j < m && !(*visited)[i][j] {
		if heights[i][j] < height {
			return
		}
		(*visited)[i][j] = true

		dfs(heights, visited, heights[i][j], i+1, j, n, m)
		dfs(heights, visited, heights[i][j], i, j+1, n, m)
		dfs(heights, visited, heights[i][j], i-1, j, n, m)
		dfs(heights, visited, heights[i][j], i, j-1, n, m)
	}
}

// func pacificAtlantic(heights [][]int) [][]int {
// 	n, m := len(heights), len(heights[0])
// 	canReachP := make([][]bool, len(heights))
// 	canReachA := make([][]bool, len(heights))

// 	for i := 0; i < len(heights); i++ {
// 		canReachP[i] = make([]bool, len(heights[0]))
// 		canReachA[i] = make([]bool, len(heights[0]))
// 	}

// 	for i := 0; i < len(heights); i++ {
// 		canReachP[i][0] = true
// 		canReachA[i][m-1] = true
// 	}
// 	for i := 0; i < len(heights[0]); i++ {
// 		canReachP[0][i] = true
// 		canReachA[n-1][i] = true
// 	}

// 	queue := [][]int{}

// 	queue = append(queue, []int{0, 0})

// 	for len(queue) != 0 {
// 		node := queue[0]
// 		i := node[0]
// 		j := node[1]

// 		if i > 0 && j > 0 && (canReachP[i-1][j] || canReachP[i][j-1]) {
// 			if heights[i][j] >= heights[i-1][j] || heights[i][j] >= heights[i][j-1] {
// 				canReachP[i][j] = true
// 			}
// 		}

// 		if i < n-1 && canReachP[i][j] {
// 			queue = append(queue, []int{i + 1, j})
// 		}

// 		if j < m-1 && canReachP[i][j] {
// 			queue = append(queue, []int{i, j + 1})
// 		}

// 		queue = queue[1:]

// 	}

// 	queue = append(queue, []int{n - 1, m - 1})
// 	for len(queue) != 0 {
// 		node := queue[0]
// 		i := node[0]
// 		j := node[1]

// 		if i < n-1 && j < m-1 && (canReachA[i+1][j] || canReachA[i][j+1]) {
// 			if heights[i][j] >= heights[i+1][j] || heights[i][j] >= heights[i][j+1] {
// 				canReachA[i][j] = true
// 			}
// 		}

// 		if i > 0 && canReachA[i][j] {
// 			queue = append(queue, []int{i - 1, j})
// 		}

// 		if j > 0 && canReachA[i][j] {
// 			queue = append(queue, []int{i, j - 1})
// 		}

// 		queue = queue[1:]

// 	}

// 	results := [][]int{}

// 	for i := 0; i < len(heights); i++ {
// 		for j := 0; j < len(heights[0]); j++ {
// 			if canReachA[i][j] && canReachP[i][j] {
// 				results = append(results, []int{i, j})
// 			}
// 		}
// 	}

// 	return results
// }
