package main

import "fmt"

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
}

func maxAreaOfIsland(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				newCount := floodFill(grid, i, j)
				if newCount > count {
					count = newCount
				}
			}

		}
	}

	return count
}

func floodFill(grid [][]int, sr int, sc int) int {
	count := 0

	if grid[sr][sc] == 1 {
		grid[sr][sc] = 2
		count++

		if sr-1 >= 0 {
			count += floodFill(grid, sr-1, sc)
		}
		if sc-1 >= 0 {
			count += floodFill(grid, sr, sc-1)
		}

		if sr+1 < len(grid) {
			count += floodFill(grid, sr+1, sc)
		}
		if sc+1 < len(grid[0]) {
			count += floodFill(grid, sr, sc+1)

		}

		return count
	} else {
		return count
	}

}
