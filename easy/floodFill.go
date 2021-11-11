package main

import "fmt"

func main() {
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2))
	fmt.Println(floodFill([][]int{{2, 2, 2}, {2, 2, 2}}, 0, 0, 3))
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	target := image[sr][sc]

	if target == newColor {
		return image
	} else {
		return floodFill2(image, sr, sc, newColor, target)

	}
}

func floodFill2(image [][]int, sr int, sc int, newColor, target int) [][]int {

	if image[sr][sc] == target {
		image[sr][sc] = newColor

		if sr-1 >= 0 {
			image = floodFill2(image, sr-1, sc, newColor, target)
		}
		if sc-1 >= 0 {
			image = floodFill2(image, sr, sc-1, newColor, target)

		}

		if sr+1 < len(image) {
			image = floodFill2(image, sr+1, sc, newColor, target)
		}
		if sc+1 < len(image[0]) {
			image = floodFill2(image, sr, sc+1, newColor, target)

		}

		return image
	} else {
		return image
	}

}
