package main

import (
	"fmt"
)

//func convert(s string, numRows int) string {
//	var res string
//	if numRows == 1 {
//		return s
//	}
//	for i := 0; i < numRows; i++ {
//		j := i
//		for j < len(s) {
//			res += s[j : j+1]
//			if i != 0 && i != numRows-1 && j+(numRows-i)*2-2 < len(s) {
//				res += s[j+(numRows-i)*2-2 : j+(numRows-i)*2-1]
//			}
//			j += (numRows - 1) * 2
//		}
//	}
//	return res
//}

func convert(s string, numRows int) string {
	fmt.Printf("s: %v\n", s)
	fmt.Printf("numRowss: %v\n", numRows)
	length := len(s)
	result := make([]byte, length)
	side := numRows*2 - 2
	fmt.Printf("side: %v\n", side)
	if numRows == 1 {
		return s
	}
	index := 0
	for i := 0; i < numRows; i++ {
		fmt.Printf("-----%v-----\n", i)
		j := i
		if i == 0 || i == numRows-1 {
			for j < length {
				fmt.Printf("index: %v\n", index)
				fmt.Printf("j: %v\n", j)
				result[index] = s[j]
				index++
				j += side
			}
		} else {
			top := true
			for j < length {
				fmt.Printf("index: %v\n", index)
				fmt.Printf("j: %v\n", j)
				result[index] = s[j]
				index++
				if top {
					j += side - i*2
				} else {
					j += i * 2
				}
				top = !top
			}
		}
	}
	return string(result)
}

func main() {
	fmt.Println(convert("abcdefghijklmnopqrstuvwxyz", 5))
}
