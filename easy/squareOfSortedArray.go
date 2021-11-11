package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, 1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}

func sortedSquares(nums []int) []int {
	newInts := make([]int, len(nums))

	index := len(nums) - 1
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i]*nums[i] <= nums[j]*nums[j] {
			newInts[index] = nums[j] * nums[j]
			j--
		} else {
			newInts[index] = nums[i] * nums[i]
			i++
		}
		index--
	}

	return newInts
}
