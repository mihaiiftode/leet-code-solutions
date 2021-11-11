package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	return backtrack(nums, 0)
}

func backtrack(nums []int, start int) [][]int {
	if start == len(nums) {
		cp := make([]int, len(nums))
		copy(cp, nums)
		return [][]int{cp}
	}

	result := make([][]int, 0)

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]

		result = append(result, backtrack(nums, i+1)...)

		nums[start], nums[i] = nums[i], nums[start]

	}

	return result
}
