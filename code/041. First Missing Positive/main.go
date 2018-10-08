package main

import (
	"fmt"
	"sort"
)

//use sort
//Runtime: 0 ms, faster than 100.00% of Go online submissions for First Missing Positive.
//func firstMissingPositive(nums []int) int {
//	sort.Ints(nums)
//	var j = 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] < 1 {
//			continue
//		}
//		if j != nums[i] {
//			if j+1 == nums[i] {
//				j++
//				continue
//			}
//			return j + 1
//		}
//	}
//	return j + 1
//}

//no use sort
//Runtime: 0 ms, faster than 100.00% of Go online submissions for First Missing Positive.
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	var data = []int{1, 2, 2, 3, 6, 7}
	fmt.Println(firstMissingPositive(data))
}
