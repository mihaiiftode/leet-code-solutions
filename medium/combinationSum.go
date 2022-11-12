package main

func main() {
	// combinationSum([]int{2, 3, 6, 7}, 7)
	combinationSum([]int{2, 3, 5}, 8)
}

func combinationSum(candidates []int, target int) [][]int {

	result := combinationSumRecursive(candidates, make([]int, 0), [][]int{}, target, 0)

	return result
}

func combinationSumRecursive(candidates, currentCandidate []int, result [][]int, target, sum int) [][]int {
	for i := 0; i < len(candidates); i++ {
		if sum+candidates[i] <= target {
			currentCandidate = append(currentCandidate, candidates[i])
			sum += candidates[i]
			if sum == target {
				result = append(result, append([]int{}, currentCandidate...))
			} else {
				result = combinationSumRecursive(candidates[i:], currentCandidate, result, target, sum)
			}

			sum -= currentCandidate[len(currentCandidate)-1]
			currentCandidate = currentCandidate[:len(currentCandidate)-1]
		}
	}

	return result
}

func combinationSumRecursive2(candidates, currentCandidate []int, result [][]int, target, sum int) [][]int {
	for i := 0; i < len(candidates); i++ {
		if sum+candidates[i] <= target {
			currentCandidate = append(currentCandidate, candidates[i])
			sum += candidates[i]
			if sum == target {
				result = append(result, append([]int{}, currentCandidate...))
			} else {
				result = combinationSumRecursive(candidates[i:], currentCandidate, result, target, sum)
			}

			sum -= currentCandidate[len(currentCandidate)-1]
			currentCandidate = currentCandidate[:len(currentCandidate)-1]
		}
	}

	return result
}
