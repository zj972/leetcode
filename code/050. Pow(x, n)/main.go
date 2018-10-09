package main

import (
	"fmt"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Pow(x, n).
//func myPow(x float64, n int) float64 {
//	return math.Pow(x, float64(n))
//}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Pow(x, n).
func myPow(x float64, n int) float64 {
	v := 1.0
	for i := n; i != 0; i /= 2 {
		if i%2 != 0 {
			v *= x
		}
		x *= x
	}
	if n < 0 {
		return 1 / v
	}
	return v
}

func main() {
	fmt.Println(myPow(2.00000, -2))
}
