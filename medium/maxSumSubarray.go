package main

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-subarray/solutions/1595195/c-python-7-simple-solutions-w-explanation-brute-force-dp-kadane-divide-conquer/
func main() {
	// fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// fmt.Println(maxSubArray([]int{-2, 1}))
	fmt.Println(maxSubArray([]int{-2, 1, -2, -2, -2, -2, -2, -2, 6, -2, -2, -2, -2, -2}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	localMax := dp[0]
	for i := 1; i < len(nums); i++ {

		dp[i] = max(dp[i-1]+nums[i], nums[i]+nums[i-1])
		localMax = max(localMax, dp[i])
		localMax = max(localMax, nums[i])
	}

	return localMax
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
