package main

import (
	"fmt"
)

//Runtime: 76 ms, faster than 32.03% of Go online submissions for Remove Duplicates from Sorted Array.
//func removeDuplicates(nums []int) int {
//	length := len(nums)
//	if length < 2 {
//		return length
//	}
//	j := 1
//	for i := 1; i < length; i++ {
//		fmt.Println(i, j, nums)
//		if nums[j] == nums[j-1] {
//			if j == len(nums)-1 {
//				nums = nums[:j]
//			} else {
//				nums = append(nums[:j], nums[j+1:]...)
//			}
//		} else {
//			j++
//		}
//	}
//	nums = nums[:j]
//	return j
//}

//Runtime: 60 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted Array.
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		if i < 1 || nums[i-1] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	data := []int{1, 1}
	fmt.Println(removeDuplicates(data))
	fmt.Println(data)
}
