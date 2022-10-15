package main

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

func rotate(matrix [][]int) {
	p1, p2, p3, p4 := 0, 0, 0, 0
	length := len(matrix) - 1
	for i := 0; i <= length/2; i++ {
		for j := i; j <= length-i; j++ {
			p1 = matrix[i][j]
			p2 = matrix[j][length-i]
			p3 = matrix[length-i][length-j]
			p4 = matrix[length-j][i]

			matrix[i][j] = p4
			matrix[j][length-i] = p1
			matrix[length-i][length-j] = p2
			matrix[length-j][i] = p3
		}
	}
}

func transpose(matrix [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}

func reflect(matrix [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][n-1-j]
			matrix[i][n-1-j] = tmp
		}
	}
}

func rotate2(matrix [][]int) {
	n := len(matrix[0])
	transpose(matrix, n)
	reflect(matrix, n)
}
