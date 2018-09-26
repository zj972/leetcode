/*
1.如果确定切片长度，最好make指定长度
2.通过_,ok:=$value来判断$value是否存在
3.使用传统for;;要比for-range快
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			res[0] = m[target-nums[i]]
			res[1] = i
			break
		}
		m[nums[i]] = i
	}
	return res
}

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}
