package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {

	var count_map = make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		_, ok := count_map[nums[i]]
		if ok {
			return true
		} else {
			count_map[nums[i]] = 1
		}
	}

	return false
}

// func containsDuplicate(nums []int) bool {
//     s := map[int]bool{}
//     for _, n := range nums {
//         if s[n] {
//             return true
//         }
//         s[n] = true
//     }
//     return false
// }
