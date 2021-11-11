package main

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{{0, 0}}))
}

type point struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	freshCount := 0

	queue := []point{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, point{x: i, y: j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 && len(queue) == 0 {
		return 0
	}

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	step := -1

	for len(queue) > 0 {
		k := len(queue)
		for ; k > 0; k-- {
			item := queue[0]
			queue = queue[1:]
			for _, d := range dirs {
				x, y := item.x+d[0], item.y+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					freshCount--
					queue = append(queue, point{x: x, y: y})
				}
			}

		}
		step++
	}

	if freshCount > 0 {
		return -1
	}

	return step

}
