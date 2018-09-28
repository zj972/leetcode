package main

import (
	"fmt"
)

//Runtime: 12 ms, faster than 100.00% of Go online submissions for Find First and Last Position of Element in Sorted Array.
//func searchRange(nums []int, target int) []int {
//	list := []int{-1, -1}
//	for idx, val := range nums {
//		if val == target && (idx-1 < 0 || nums[idx-1] != target) {
//			list[0] = idx
//		}
//		if val == target && (idx+1 > len(nums)-1 || nums[idx+1] != target) {
//			list[1] = idx
//		}
//	}
//	return list
//}

func searchRange(nums []int, target int) []int {
	leftStart := 0
	rightStart := len(nums) - 1
	if rightStart < 0 {
		return []int{-1, -1}
	}
	mid := (leftStart + rightStart) / 2
	start := 0
	for {
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				start = mid
				break
			} else {
				rightStart = mid - 1
			}
		} else if nums[mid] > target {
			rightStart = mid - 1
		} else {
			leftStart = mid + 1
		}
		mid = (leftStart + rightStart) / 2
		if leftStart >= rightStart {
			break
		}
	}
	if nums[mid] == target {
		start = mid
	} else {
		return []int{-1, -1}
	}

	leftEnd := 0
	rightEnd := len(nums) - 1
	mid = (leftEnd + rightEnd) / 2
	for {
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				break
			} else {
				leftEnd = mid + 1
			}
		} else if nums[mid] > target {
			rightEnd = mid - 1
		} else {
			leftEnd = mid + 1
		}
		mid = (leftEnd + rightEnd) / 2
		if leftEnd >= rightEnd {
			break
		}
	}
	if nums[mid] == target {
		return []int{start, mid}
	} else {
		return []int{-1, -1}
	}
}

func main() {
	var data = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(searchRange(data, 5))
}
