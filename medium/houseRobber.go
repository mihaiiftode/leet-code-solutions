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

func rob1(nums []int) int {
	memo := make([]int, len(nums)+1)

	memo[0] = 0
	memo[1] = nums[0]

	for i := 1; i < len(nums); i++ {
		memo[i+1] = max(memo[i-1]+nums[i], memo[i])
	}

	return memo[len(nums)]
}
