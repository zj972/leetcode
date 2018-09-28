package main

import (
	"fmt"
)

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
//func search(nums []int, target int) int {
//	for idx, val := range nums {
//		if val == target {
//			return idx
//		}
//	}
//	return -1
//}

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	leftPart := true
	if right < 0 || (target > nums[right] && target < nums[left]) {
		return -1
	}
	if nums[right] < nums[left] && nums[right] >= target {
		leftPart = false
	}
	var mid int
	for {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if left >= right-1 {
			if nums[left] == target {
				return left
			} else if nums[right] == target {
				return right
			} else {
				return -1
			}
		}
		if leftPart {
			if nums[mid] < nums[right] && nums[right] < nums[left] {
				right = mid - 1
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] > nums[left] && nums[right] < nums[left] {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
}

func main() {
	var data = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(data, 5))
}
