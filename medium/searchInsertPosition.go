package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 2))
}

func searchInsert(nums []int, target int) int {
	return insert(nums, target, 0, len(nums))

}

func insert(nums []int, target, left, right int) int {
	if right > left {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			return insert(nums, target, left, mid)
		} else {
			return insert(nums, target, mid+1, right)

		}

	}

	return left

}
