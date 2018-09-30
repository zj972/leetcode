package main

import (
	"fmt"
	"sort"
)

//Runtime: 408 ms, faster than 0.00% of Go online submissions for Combination Sum.
//func combinationSum(candidates []int, target int) [][]int {
//	fmt.Println(candidates)
//	if candidates[0] > target {
//		if len(candidates) == 1 {
//			return [][]int{}
//		} else {
//			return combinationSum(candidates[1:], target)
//		}
//	} else if candidates[0] == target {
//		if len(candidates) == 1 {
//			return [][]int{[]int{candidates[0]}}
//		} else {
//			return append([][]int{[]int{candidates[0]}}, combinationSum(candidates[1:], target)...)
//		}
//	} else {
//		tmp := combinationSum(candidates, target-candidates[0])
//		for i := 0; i < len(tmp); i++ {
//			tmp[i] = append(tmp[i], candidates[0])
//		}
//		if len(candidates) == 1 {
//			return tmp
//		} else {
//			return append(tmp, combinationSum(candidates[1:], target)...)
//		}
//	}
//}

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Combination Sum.
func findResult(candidates []int, targetValue int, currentResult []int, results *[][]int) {
	for i, value := range candidates {
		if value < targetValue {
			findResult(candidates[i:], targetValue-value, append(currentResult, value), results)
		} else if value == targetValue {
			currentResult = append(currentResult, value)
			tmp := make([]int, len(currentResult))
			copy(tmp, currentResult)
			*results = append(*results, tmp)
		} else {
			return
		}
	}
	return
}

func combinationSum(candidates []int, target int) [][]int {
	var results [][]int
	if target < 0 || len(candidates) == 0 {
		return results
	}
	sort.Ints(candidates)
	findResult(candidates, target, []int{}, &results)
	return results
}

func main() {
	var data = []int{2, 3, 6, 7}
	fmt.Println(combinationSum(data, 7))
}
