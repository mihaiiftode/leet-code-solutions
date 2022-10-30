package main

import "fmt"

func main() {
	fmt.Println(search([]int{1}, 0))
}

// func search(nums []int, target int) int {
// 	for index, v := range nums {
// 		if v == target {
// 			return index
// 		}
// 	}
// }

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for left <= right {
		mid = left + (right-left)/2
		if target == nums[mid] {
			return nums[mid]
		}

		// left
		if nums[left] <= nums[mid] {
			if target > nums[mid] || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
			// right
		} else {
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
