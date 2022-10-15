package main

func main() {
	insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 11}, {12, 16}}, []int{4, 6})
}

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}

	index := 0

	for ; index < len(intervals) && intervals[index][1] < newInterval[0]; index++ {
		result = append(result, intervals[index])
	}

	for ; index < len(intervals) && intervals[index][0] <= newInterval[1]; index++ {
		newInterval[0] = min(intervals[index][0], newInterval[0])
		newInterval[1] = max(intervals[index][1], newInterval[1])
	}

	result = append(result, newInterval)

	for ; index < len(intervals); index++ {
		result = append(result, intervals[index])
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func insert(intervals [][]int, newInterval []int) [][]int {
// 	start := newInterval[0]
// 	startIndex := -1
// 	end := newInterval[1]
// 	endIndex := -1

// 	if len(intervals) == 0 || (start <= intervals[0][0] && end >= intervals[len(intervals)-1][1]) {
// 		return [][]int{newInterval}
// 	}

// 	if end < intervals[0][0] {
// 		return append([][]int{newInterval}, intervals...)
// 	} else if start > intervals[len(intervals)-1][1] {
// 		return append(intervals, newInterval)
// 	}
// 	myInterval := newInterval

// 	for i := 0; i < len(intervals); i++ {

// 		if start >= intervals[i][0] && start <= intervals[i][1] {
// 			myInterval[0] = intervals[i][0]
// 			startIndex = i
// 		}

// 		if end >= intervals[i][0] && end <= intervals[i][1] {
// 			myInterval[1] = intervals[i][1]
// 			endIndex = i
// 		}

// 	}

// 	if startIndex < 0 && endIndex < 0 {

// 		for i := 0; i < len(intervals); i++ {
// 			if i > 0 && start >= intervals[i-1][1] && end <= intervals[i][0] {
// 				intervals = append(intervals[:i+1], (intervals[i:])...)
// 				intervals[i] = newInterval
// 				return intervals
// 			}
// 		}
// 	}

// 	if startIndex >= 0 && endIndex >= 0 {
// 		intervals = append(intervals[:startIndex], (intervals[endIndex:])...)
// 	}
// 	if startIndex >= 0 {
// 		intervals[startIndex] = myInterval
// 	} else {
// 		intervals[endIndex] = myInterval
// 	}
// 	return intervals

// }
