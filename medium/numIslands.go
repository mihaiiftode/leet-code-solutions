package main

func main() {
	numIslands([][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}})
}

func numIslands(grid [][]byte) int {
	visited := make([][]byte, len(grid))
	count := 0
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]byte, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if visited[i][j] == 1 || grid[i][j] == 0 {
				continue
			} else {
				count += search(grid, i, j, visited)
			}
		}
	}

	return count
}

func search(grid [][]byte, row int, col int, visited [][]byte) int {

	h := len(grid)
	l := len(grid[0])

	if row < 0 || col < 0 || row >= h || col >= l || (visited)[row][col] == 1 || grid[row][col] == 0 {
		return 1
	}

	(visited)[row][col] = 1
	search(grid, row+1, col, visited) // go right
	search(grid, row-1, col, visited) //go left
	search(grid, row, col+1, visited) //go down
	search(grid, row, col-1, visited) // go up

	return 1
}
