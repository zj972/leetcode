package main

import (
	"fmt"
)

//Time Limit Exceeded
func canJump(nums []int) bool {
	var list = make([]bool, len(nums))
	var jump func(index int, nums []int, list []bool) bool
	jump = func(index int, nums []int, list []bool) bool {
		if list[index] || index == len(nums)-1 {
			return true
		}
		if (index + nums[index]) >= len(nums)-1 {
			return true
		}
		for i := index + 1; i <= index+nums[index] && i < len(nums); i++ {
			list[i] = jump(i, nums, list)
			if list[i] {
				return true
			}
		}
		return false
	}
	return jump(0, nums, list)
}

func main() {
	var data = []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(data))
}
