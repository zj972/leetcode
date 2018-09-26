/*
1.注意类型（数值）的溢出范围
*/
package main

import (
	"fmt"
	_ "math"
)

//func reverse(x int) int {
//	rev := 0
//	max := math.MaxInt32
//	min := math.MinInt32
//	for x != 0 {
//		pop := x % 10
//		x /= 10
//		if rev > max/10 || (rev == max/10 && pop > 7) {
//			return 0
//		}
//		if rev < min/10 || (rev == min/10 && pop < -8) {
//			return 0
//		}
//		rev = rev*10 + pop
//	}
//	return rev
//}

func reverse(x int) int {
	y := int32(x)
	var ans int32
	for y != 0 {
		temp := ans*10 + y%10
		if temp/10 != ans {
			return 0
		}
		ans = temp
		y = y / 10
	}
	return int(ans)
}

func main() {
	fmt.Println(reverse(33353434424))
}
