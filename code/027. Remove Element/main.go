package main

import (
	"fmt"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		fmt.Println(i, j, nums[j])
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	data := []int{1, 1, 2}
	fmt.Println(removeElement(data, 2))
	fmt.Println(data)
}
