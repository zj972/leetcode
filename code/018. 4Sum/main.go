package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	tmp := []int{}
	sort.Ints(nums)
	nSum(&ret, &tmp, nums, target, 4)
	return ret
}

func nSum(ret *[][]int, tmp *[]int, nums []int, target int, N int) {
	// return
	n := len(nums)
	if n < N || N < 2 || nums[0]*N > target || nums[n-1]*N < target {
		return
	}
	// 2sum
	if N == 2 {
		l, r := 0, n-1
		for l < r {
			if nums[l]+nums[r] > target {
				r--
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				*tmp = append(*tmp, nums[l], nums[r])
				*ret = append(*ret, dup(*tmp))
				*tmp = (*tmp)[:len(*tmp)-2]
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	} else {
		for i := 0; i < n-N+1; i++ {
			if i == 0 || i > 0 && nums[i] != nums[i-1] {
				*tmp = append(*tmp, nums[i])
				nSum(ret, tmp, nums[i+1:], target-nums[i], N-1)
				*tmp = (*tmp)[:len(*tmp)-1]
			}
		}
	}
}

func dup(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func main() {
	fmt.Println(fourSum([]int{1, 23, 43}, 12))
}
