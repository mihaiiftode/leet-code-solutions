package main

func main() {

}

type point struct {
	x int
	y int
}

func updateMatrix(mat [][]int) [][]int {

	queue := make([]point, 0)
	m, n := len(mat), len(mat[0])
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, point{x: i, y: j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	step := 0
	for len(queue) > 0 {
		k := len(queue)
		for k > 0 {
			item := queue[0]
			queue = queue[1:]
			for _, d := range dirs {
				x, y := item.x+d[0], item.y+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && mat[x][y] == -1 {
					mat[x][y] = step + 1
					queue = append(queue, point{x: x, y: y})
				}
			}
			k--
		}
		step++
	}
	return mat
}
