package main

import (
	"fmt"
	"sort"
)

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Combination Sum II.
func findResult(candidates []int, targetValue int, currentResult []int, results *[][]int) {
	for i, value := range candidates {
		if i == 0 || i != 0 && candidates[i] != candidates[i-1] {
			if value < targetValue {
				findResult(candidates[i+1:], targetValue-value, append(currentResult, value), results)
			} else if value == targetValue {
				currentResult = append(currentResult, value)
				tmp := make([]int, len(currentResult))
				copy(tmp, currentResult)
				*results = append(*results, tmp)
			} else {
				return
			}
		}
	}
	return
}

func combinationSum2(candidates []int, target int) [][]int {
	var results [][]int
	if target < 0 || len(candidates) == 0 {
		return results
	}
	sort.Ints(candidates)
	findResult(candidates, target, []int{}, &results)
	return results
}

func main() {
	var data = []int{2, 2, 3, 6, 7}
	fmt.Println(combinationSum2(data, 7))
}
