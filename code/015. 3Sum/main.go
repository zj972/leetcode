package main

import (
	"fmt"
	"sort"
)

//func threeSum(nums []int) [][]int {
//	data := [][]int{}
//	if len(nums) < 3 {
//		return data
//	}
//	sum := map[int]map[int]int{}
//	s := make(map[int]bool)
//	for i := 0; i < len(nums); i++ {
//		if _, ok := s[nums[i]]; ok {
//			continue
//		} else {
//			s[nums[i]] = true
//		}
//		m := make(map[int]int)
//		target := 0 - nums[i]
//		for j := 0; j < len(nums); j++ {
//			if i == j || m[nums[j]] == -1 {
//				continue
//			}
//			if m[target-nums[j]] == 1 {
//				x := nums[i]
//				y := nums[j]
//				z := target - nums[j]
//				if x > y {
//					x, y = y, x
//				}
//				if y > z {
//					y, z = z, y
//				}
//				if x > y {
//					x, y = y, x
//				}
//				if _, ok := sum[x]; ok {
//					if _, ok := sum[x][y]; ok {
//						continue
//					} else {
//						data = append(data, []int{x, y, z})
//						m[nums[j]] = -1
//						sum[x][y] = z
//					}
//				} else {
//					data = append(data, []int{x, y, z})
//					m[nums[j]] = -1
//					sum[x] = map[int]int{}
//					sum[x][y] = z
//				}
//				continue
//			}
//			m[nums[j]] = 1
//		}
//	}
//	return data
//}

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	l := len(nums)

	sort.Ints(nums)
	if l < 3 {
		return ret
	}
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j, k := i+1, l-1; j < k; {
			s := nums[i] + nums[j] + nums[k]
			if s == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if s < 0 {
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			} else {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0}))
}
