package main

import (
	"fmt"
	"math"
)

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Divide Two Integers.
func divide(dividend int, divisor int) int {
	var minus = false
	if dividend < 0 {
		dividend = -dividend
		minus = !minus
	}
	if divisor < 0 {
		divisor = -divisor
		minus = !minus
	}
	var num int
	var key = 1
	var val = divisor
	for val < dividend {
		val = val << 1
		key = key << 1
	}
	for dividend >= divisor {
		if val == dividend {
			num = num + key
			break
		} else if val > dividend {
			val = val >> 1
			key = key >> 1
			if key < 1 {
				break
			}
		} else if val < dividend {
			num = num + key
			dividend = dividend - val
			val = val >> 1
			key = key >> 1
		}
	}
	if minus {
		if -num < math.MinInt32 {
			return math.MinInt32
		}
		return -num
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	return num
}

func main() {
	fmt.Println(divide(-161, -2))
}
