package main

import (
	"fmt"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Multiply Strings.
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	list := make([]uint8, len(num1)+len(num2))
	for l := 0; l < len(list); l++ {
		list[l] = 48
	}
	for i := len(num1) - 1; i >= 0; i-- {
		var up uint8
		for j := len(num2) - 1; j >= 0; j-- {
			tmp := (num1[i]-48)*(num2[j]-48) + (list[i+j+1] - 48) + up
			up = tmp / 10
			list[i+j+1] = tmp%10 + 48
		}
		head := i
		for up != 0 {
			tmp := list[head] - 48 + up
			up = tmp / 10
			list[head] = tmp%10 + 48
			head--
		}
	}
	if list[0] == 48 {
		list = list[1:]
	}
	return string(list)
}

func main() {
	fmt.Println(multiply("123", "456"))
}
