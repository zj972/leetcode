package main

import (
	"fmt"
)

//在一的基础上过滤相同值的递归
//Runtime: 12 ms, faster than 100.00% of Go online submissions for Permutations II.
//func permuteUnique(nums []int) [][]int {
//	sort.Ints(nums)
//	var add func(item []int, addList []int) [][]int
//	add = func(item []int, addList []int) [][]int {
//		if len(addList) == 0 {
//			return [][]int{item}
//		}
//		ret := [][]int{}
//		for i := 0; i < len(addList); i++ {
//			if i != 0 && addList[i] == addList[i-1] {
//				continue
//			}
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

//由于在基础数组进行交换操作，所以无法保证有序，使用map确认是否已经有相同值的递归
//Runtime: 12 ms, faster than 100.00% of Go online submissions for Permutations II.
func permuteUnique(nums []int) [][]int {
	var allPermutations [][]int
	var genPermutation func(nums []int, allPermutations *[][]int, index int)
	genPermutation = func(nums []int, allPermutations *[][]int, index int) {
		if index == len(nums) {
			*allPermutations = append(*allPermutations, append([]int{}, nums...))
			return
		}
		tmp := map[int]bool{}
		for i := index; i < len(nums); i++ {
			if i == index {
				genPermutation(nums, allPermutations, index+1)
			} else {
				if tmp[nums[i]] {
					continue
				}
				nums[i], nums[index] = nums[index], nums[i]
				genPermutation(nums, allPermutations, index+1)
				nums[i], nums[index] = nums[index], nums[i]
			}
			tmp[nums[i]] = true
		}
	}
	genPermutation(nums, &allPermutations, 0)
	return allPermutations
}

func main() {
	var data = []int{0, 1, 0, 0, 9}
	fmt.Println(permuteUnique(data))
}
