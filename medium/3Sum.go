package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)
	if len(nums) < 3 {
		return [][]int{}
	}

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				j++
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

// func threeSum(nums []int) [][]int {
// 	solution := make(map[[3]int][]int)

// 	sort.Ints(nums)
// 	if len(nums) < 3 {
// 		return [][]int{}
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				triplet := [3]int{nums[i], nums[j], nums[k]}

// 				if nums[i]+nums[j]+nums[k] == 0 {
// 					solution[triplet] = []int{nums[i], nums[j], nums[k]}
// 				}
// 			}
// 		}
// 	}

// 	result := [][]int{}
// 	for _, triplet := range solution {
// 		result = append(result, triplet)
// 	}
// 	return result
// }
