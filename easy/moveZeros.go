package main

func main() {
	moveZeroes([]int{1, 2})
	// moveZeroes([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0})
}

func moveZeroes(nums []int) {

	for i, j := 0, 0; i < len(nums) && j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
