package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(numbers []int, target int) []int {
	hash := make(map[int]int)

	for i := 0; i < len(numbers); i++ {
		diff := target - numbers[i]

		if val, ok := hash[diff]; ok {
			return []int{val + 1, i + 1}
		}

		hash[numbers[i]] = i
	}

	return []int{}
}
