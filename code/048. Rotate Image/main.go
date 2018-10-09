package main

import (
	"fmt"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate Image.
func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := i; j < l-1-i; j++ {
			matrix[i][j], matrix[j][l-1-i], matrix[l-1-i][l-1-j], matrix[l-1-j][i] = matrix[l-1-j][i], matrix[i][j], matrix[j][l-1-i], matrix[l-1-i][l-1-j]
		}
	}
}

func main() {
	var data = [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15},
		[]int{16, 17, 18, 19, 20},
		[]int{21, 22, 23, 24, 25},
	}
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
	fmt.Println("--------------")
	rotate(data)
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}
