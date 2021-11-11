package main

func main() {

	// fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))

}

func rotate(nums []int, k int) {
	rotated := k % len(nums)

	copy(nums, append(nums[len(nums)-rotated:], nums[:len(nums)-rotated]...))
}
