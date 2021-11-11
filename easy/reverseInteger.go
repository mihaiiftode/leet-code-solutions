package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-321))
}

func reverse(x int) int {
	// sign := math.Signbit(float64(x))

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	var y int
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}

	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y

}
