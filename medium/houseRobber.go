package main

func main() {

}

func rob(nums []int) int {
	cache := make([]int, len(nums))
	cache[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if i < 2 {
			cache[i] = max(nums[i], cache[i-1])
		} else {
			cache[i] = max(nums[i]+cache[i-2], cache[i-1])
		}
	}

	return cache[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
