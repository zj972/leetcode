package main

import "fmt"

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Next Permutation.
func nextPermutation(nums []int) {
	var key int = -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			key = i
			break
		}
	}
	if key != -1 {
		//替换最近的值
		var s int
		for i := key + 1; i < len(nums); i++ {
			if nums[i] <= nums[key] {
				s = i - 1
				break
			}
		}
		if s == 0 {
			s = len(nums) - 1
		}
		nums[key], nums[s] = nums[s], nums[key]
	}
	//反转顺序（key---len-1）
	for i, j := key+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}

func main() {
	var data = []int{5, 4, 7, 5, 3, 2}
	nextPermutation(data)
	fmt.Println(data)
}
