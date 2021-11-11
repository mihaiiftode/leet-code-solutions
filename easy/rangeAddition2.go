package main

func main() {

}

func maxCount(m int, n int, ops [][]int) int {

	minX, minY := m, n
	for _, v := range ops {
		minX = min(minX, v[0])
		minY = min(minY, v[1])
	}

	return minX * minY
}

func min(a, b int) int {

	if a <= b {
		return a
	} else {
		return b
	}

}
