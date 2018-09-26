package main

import "fmt"

func maxArea(height []int) int {
	var (
		maxarea = 0
		l       = 0
		r       = len(height) - 1
	)
	for l < r {
		var long int
		if height[l] < height[r] {
			long = height[l]
			l++
		} else {
			long = height[r]
			r--
		}
		if maxarea < long*(r-l+1) {
			maxarea = long * (r - l + 1)
		}
	}
	return maxarea
}

func main() {
	fmt.Println(maxArea([]int{1, 2, 3}))
}
