package main

import (
	"fmt"
)

//Runtime: 0 ms
//Your runtime beats 100.00 % of golang submissions.
func isValid(s string) bool {
	var list = []byte{}
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "(":
			list = append(list, 1)
		case ")":
			if len(list) == 0 || list[len(list)-1] != 1 {
				return false
			}
			list = list[:len(list)-1]
		case "{":
			list = append(list, 2)
		case "}":
			if len(list) == 0 || list[len(list)-1] != 2 {
				return false
			}
			list = list[:len(list)-1]
		case "[":
			list = append(list, 3)
		case "]":
			if len(list) == 0 || list[len(list)-1] != 3 {
				return false
			}
			list = list[:len(list)-1]
		}
	}
	if len(list) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("{}[])"))
}
