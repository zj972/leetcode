package main

import (
	"fmt"
)

//. 46
//Q 81
//Runtime: 56 ms, faster than 18.52% of Go online submissions for N-Queens.
//func solveNQueens(n int) [][]string {
//	ret := [][]string{}
//	tmp := make([][]uint8, n)
//	row := make([]bool, n)
//	for i := 0; i < n; i++ {
//		tmp[i] = make([]uint8, n)
//		for j := 0; j < n; j++ {
//			tmp[i][j] = 46
//		}
//		row[i] = true
//	}
//
//	lBias := make([]bool, 2*n-1)
//	rBias := make([]bool, 2*n-1)
//	for i := 0; i < 2*n-1; i++ {
//		lBias[i] = true
//		rBias[i] = true
//	}
//
//	addQ(tmp, 0, 0, row, lBias, rBias, &ret, false)
//	return ret
//}
//
//func addQ(data [][]uint8, x int, y int, row []bool, lBias []bool, rBias []bool, ret *[][]string, back bool) {
//	fmt.Printf("data: %v, x: %v, y: %v, row: %v, lBias: %v, rBias: %v\n", data, x, y, row, lBias, rBias)
//	//校验当前位置是否可以插入Q
//	if back {
//		//最上面回滚，结束
//		if x == 0 {
//			return
//		}
//		i := 0
//		for ; i < len(data); i++ {
//			if data[x-1][i] == 81 {
//				data[x-1][i] = 46
//				setCheck(x-1, i, row, lBias, rBias, true)
//				break
//			}
//		}
//		if i == len(data)-1 {
//			addQ(data, x-1, i, row, lBias, rBias, ret, true)
//		} else {
//			addQ(data, x-1, i+1, row, lBias, rBias, ret, false)
//		}
//	} else {
//		if row[y] && lBias[x+y] && rBias[(len(rBias)-1)/2+x-y] {
//			data[x][y] = 81
//			if x == len(data)-1 {
//				tmp := []string{}
//				for i := 0; i < len(data); i++ {
//					tmp = append(tmp, string(data[i]))
//				}
//				*ret = append(*ret, tmp)
//
//				data[x][y] = 46
//				addQ(data, x, y, row, lBias, rBias, ret, true)
//			} else {
//				setCheck(x, y, row, lBias, rBias, false)
//				addQ(data, x+1, 0, row, lBias, rBias, ret, false)
//			}
//		} else {
//			if y+1 < len(data) {
//				addQ(data, x, y+1, row, lBias, rBias, ret, false)
//			} else {
//				addQ(data, x, y, row, lBias, rBias, ret, true)
//			}
//		}
//	}
//}
//
//func setCheck(x int, y int, row []bool, lBias []bool, rBias []bool, key bool) {
//	row[y] = key
//	lBias[x+y] = key
//	rBias[(len(rBias)-1)/2+x-y] = key
//}

//Runtime: 8 ms, faster than 100.00% of Go online submissions for N-Queens.
func solveNQueens(n int) [][]string {
	return dfs(nil, nil, n)
}

func dfs(solutions [][]string, xs []int, n int) [][]string {
	qy := len(xs)
	if qy == n {
		return append(solutions, solution(xs, n))
	}

qx:
	for qx := 0; qx < n; qx++ {
		for y, x := range xs {
			if x == qx || qy-y == x-qx || qy-y == qx-x {
				continue qx
			}
		}
		solutions = dfs(solutions, append(xs, qx), n)
	}
	return solutions
}

func solution(xs []int, n int) []string {
	b := make([]byte, n)
	for i := range b {
		b[i] = '.'
	}

	s := make([]string, n)
	for y, x := range xs {
		b[x] = 'Q'
		s[y] = string(b)
		b[x] = '.'
	}
	return s
}

func main() {
	fmt.Println(solveNQueens(4))
}
