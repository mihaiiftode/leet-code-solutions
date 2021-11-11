package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	glolbalMaximum := math.MinInt32
	sum := 0

	// for i, _ := range nums {
	// 	sum := 0
	// 	for j := i; j < len(nums); j++ {
	// 		sum = sum + nums[j]
	// 		if sum > glolbalMaximum {
	// 			glolbalMaximum = sum
	// 		}
	// 	}
	// }

	for _, v := range nums {

		sum = max(v, sum+v)
		if sum > glolbalMaximum {
			glolbalMaximum = sum
		}

	}

	return glolbalMaximum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
