package main

import (
	"fmt"
)

//Runtime: 8 ms, faster than 83.59% of Go online submissions for Permutations.
//func permute(nums []int) [][]int {
//	var add func(item []int, addList []int) [][]int
//	add = func(item []int, addList []int) [][]int {
//		if len(addList) == 0 {
//			return [][]int{item}
//		}
//		ret := [][]int{}
//		for i := 0; i < len(addList); i++ {
//			tmpItem := make([]int, len(item))
//			tmpAddList := make([]int, len(addList))
//			copy(tmpItem, item)
//			copy(tmpAddList, addList)
//			tmpItem = append(tmpItem, tmpAddList[i])
//			if i == len(addList)-1 {
//				tmpAddList = tmpAddList[:i]
//			} else {
//				tmpAddList = append(tmpAddList[:i], tmpAddList[i+1:]...)
//			}
//			tmp := add(tmpItem, tmpAddList)
//			ret = append(ret, tmp...)
//		}
//		return ret
//	}
//	return add([]int{}, nums)
//}

//只使用原数组切片，没有tmp中间空间
//Runtime: 4 ms, faster than 100.00% of Go online submissions for Permutations.
func permute(nums []int) [][]int {
	var allPermutations [][]int
	var genPermutation func(nums []int, allPermutations *[][]int, index int)
	genPermutation = func(nums []int, allPermutations *[][]int, index int) {
		if index == len(nums) {
			*allPermutations = append(*allPermutations, append([]int{}, nums...))
			return
		}
		for i := index; i < len(nums); i++ {
			if i == index {
				genPermutation(nums, allPermutations, index+1)
			} else {
				nums[i], nums[index] = nums[index], nums[i]
				genPermutation(nums, allPermutations, index+1)
				nums[i], nums[index] = nums[index], nums[i]
			}
		}
	}
	genPermutation(nums, &allPermutations, 0)
	return allPermutations
}

func main() {
	var data = []int{4, 5, 2, 1}
	fmt.Println(permute(data))
}
