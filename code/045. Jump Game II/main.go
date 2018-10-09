package main

import (
	"fmt"
)

//Runtime: 1720 ms, faster than 2.22% of Go online submissions for Jump Game II.
//func jump(nums []int) int {
//	var list = make([]int, len(nums))
//	var jumpFrom func(index int, nums []int, list []int) int
//	jumpFrom = func(index int, nums []int, list []int) int {
//		if list[index] != 0 || index == len(nums)-1 {
//			return list[index]
//		}
//		var min int = math.MaxInt32
//		if (index + nums[index]) >= len(nums)-1 {
//			list[index] = 1
//			return 1
//		}
//		for i := index + 1; i <= index+nums[index] && i < len(nums); i++ {
//			tmp := jumpFrom(i, nums, list)
//			fmt.Println(i, tmp, list)
//			if min > tmp {
//				min = tmp
//			}
//		}
//		list[index] = min + 1
//		return min + 1
//	}
//	return jumpFrom(0, nums, list)
//}

//Runtime: 12 ms, faster than 100.00% of Go online submissions for Jump Game II.
//func jump(nums []int) int {
//	n := len(nums)
//	if n == 0 || n == 1 || nums[0] == 0 {
//		return 0
//	}
//	start, end, jumps := 0, 0, 0
//	for dest := 0; dest < n-1; {
//		for i := start; i <= end; i++ {
//			if dest < i+nums[i] {
//				dest = i + nums[i]
//			}
//		}
//		start = end + 1
//		end = dest
//		jumps++
//	}
//	return jumps
//}

//Runtime: 12 ms, faster than 100.00% of Go online submissions for Jump Game II.
func jump(nums []int) int {
	var ret, curMax, curRch int
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i], curMax, curRch, ret)
		if curRch < i {
			ret++
			curRch = curMax
		}
		if curMax < nums[i]+i {
			curMax = nums[i] + i
		}
	}
	return ret
}

func main() {
	var data = []int{4, 5, 2, 1, 3, 4, 5, 1, 6, 2, 1}
	fmt.Println(jump(data))
}
