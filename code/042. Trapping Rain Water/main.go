package main

import (
	"fmt"
)

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Trapping Rain Water.
//func trap(height []int) int {
//	var sum int
//	var i int
//	for i < len(height) {
//		if height[i] == 0 {
//			i++
//			continue
//		}
//		h := height[i]
//		var j int = i + 1
//		var right int
//		var tmp int
//		var maxJ int = i + 1
//		for j < len(height) {
//			if height[j] < h {
//				if height[maxJ] < height[j] {
//					maxJ = j
//				}
//				tmp += height[j]
//				j++
//				continue
//			}
//			right = height[j]
//			break
//		}
//		if right == 0 {
//			if maxJ == i+1 {
//				i++
//				continue
//			}
//			j = maxJ
//			h = height[j]
//			tmp = 0
//			for k := i + 1; k < j; k++ {
//				tmp += height[k]
//			}
//		}
//		sum += (j-i-1)*h - tmp
//		i = j
//	}
//	return sum
//}

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Trapping Rain Water.
func trap(height []int) int {
	leng := len(height)
	left := 0
	right := leng - 1
	sum := 0
	maxRight := 0
	maxLeft := 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				sum += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				sum += maxRight - height[right]
			}
			right--
		}
	}
	return sum
}

func main() {
	var data = []int{4, 2, 3}
	fmt.Println(trap(data))
}
