package main

import (
	"fmt"
	"math"
	"sort"
)

//Runtime: 696 ms
//Your runtime beats 1.71 % of golang submissions.
//func threeSumClosest(nums []int, target int) int {
//	var sum func([]int, int, int) int
//	sum = func(nums []int, target int, num int) int {
//		if num == 0 {
//			return target
//		}
//		if len(nums) < num {
//			return math.MaxInt32
//		}
//		t1 := sum(nums[1:], target, num)
//		t2 := sum(nums[1:], target+nums[0], num-1)
//		v1 := t1
//		v2 := t2
//		if v1 < 0 {
//			v1 = -v1
//		}
//		if v2 < 0 {
//			v2 = -v2
//		}
//		if v1 < v2 {
//			return t1
//		}
//		return t2
//	}
//	return sum(nums, -target, 3) + target
//}

//Runtime: 4 ms
//Your runtime beats 100.00 % of golang submissions.
func threeSumClosest(nums []int, target int) int {
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	minDistance := math.MaxInt32
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums); i++ {
		innerTarget := target - nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			s := nums[l] + nums[r]
			if s == innerTarget {
				return target
			} else if s > innerTarget {
				r--
			} else {
				l++
			}
			if abs(minDistance-target) > abs(s+nums[i]-target) {
				minDistance = s + nums[i]
			}
		}

	}
	return minDistance
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
