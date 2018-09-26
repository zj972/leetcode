package main

import "fmt"

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	if l := len(nums1) + len(nums2); l%2 == 0 {
//		return (findKth(nums1, nums2, l/2) + findKth(nums1, nums2, l/2+1)) / 2.0
//	} else {
//		return findKth(nums1, nums2, l/2+1)
//	}
//}
//
//func findKth(nums1 []int, nums2 []int, key int) float64 {
//	if len(nums1) == 0 && len(nums2) == 0 {
//		return 0
//	}
//	if len(nums1) == 0 {
//		return float64(nums2[key-1])
//	}
//	if len(nums2) == 0 {
//		return float64(nums1[key-1])
//	}
//	for i, j := 0, 0; i < len(nums1); i++ {
//		for {
//			if nums1[i] < nums2[j] {
//				break
//			} else {
//				if i+j+1 == key {
//					return float64(nums2[j])
//				}
//				j++
//				if j >= len(nums2) {
//					return float64(nums1[key-len(nums2)-1])
//				}
//			}
//		}
//		if i+j+1 == key {
//			return float64(nums1[i])
//		}
//	}
//	return float64(nums2[key-len(nums1)-1])
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var total = len(nums1) + len(nums2)
	var mid = total / 2
	var median, previous, i, j, idx int

	for {
		var nums1Reached = i == len(nums1)
		var nums2Reached = j == len(nums2)

		if nums1Reached {
			median = nums2[j]
			j++
		} else if nums2Reached {
			median = nums1[i]
			i++
		} else {
			if nums1[i] > nums2[j] {
				median = nums2[j]
				j++
			} else {
				median = nums1[i]
				i++
			}
		}

		if idx == mid {
			break
		}
		idx++
		previous = median
	}

	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(median+previous) / 2.0
	} else {
		return float64(median)
	}
}

func main() {
	a := []int{1, 3}
	b := []int{2}
	c := findMedianSortedArrays(a, b)
	fmt.Println(c)
}
