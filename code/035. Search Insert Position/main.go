package main

import (
	"fmt"
)

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Search Insert Position.
//func searchInsert(nums []int, target int) int {
//	for idx, val := range nums {
//		if val == target || val > target {
//			return idx
//		}
//	}
//	return len(nums)
//}

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Search Insert Position.
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	if right < 0 {
		return 0
	}
	mid := (left + right) / 2
	for {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
		if left >= right {
			break
		}
	}
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return mid
	} else {
		return mid + 1
	}
}

func main() {
	var data = []int{0, 1, 2, 4, 5, 6, 7}
	fmt.Println(searchInsert(data, 5))
}
