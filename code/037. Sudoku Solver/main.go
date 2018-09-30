package main

import "fmt"

//Runtime: 172 ms, faster than 10.34% of Go online submissions for Sudoku Solver.
//func solveSudoku(board [][]byte) {
//	var next func(board [][]byte, x int, y int) bool
//	next = func(board [][]byte, x int, y int) bool {
//		if board[x][y] != byte(46) {
//			if x == 8 && y == 8 {
//				return true
//			}
//			if x == 8 {
//				return next(board, 0, y+1)
//			} else {
//				return next(board, x+1, y)
//			}
//		}
//		list := verify(board, x, y)
//		if len(list) == 0 {
//			return false
//		}
//		if x == 8 && y == 8 {
//			if len(list) != 1 {
//				return false
//			}
//			board[x][y] = list[0]
//			return true
//		}
//		var v bool
//		for i := 0; i < len(list); i++ {
//			board[x][y] = list[i]
//			if x == 8 {
//				v = next(board, 0, y+1)
//			} else {
//				v = next(board, x+1, y)
//			}
//			if v {
//				return true
//			}
//		}
//		if v == false {
//			board[x][y] = byte(46)
//		}
//		return v
//	}
//	next(board, 0, 0)
//}
//
//func verify(board [][]byte, x int, y int) []byte {
//	all := map[byte]bool{
//		46: false,
//		49: true,
//		50: true,
//		51: true,
//		52: true,
//		53: true,
//		54: true,
//		55: true,
//		56: true,
//		57: true,
//	}
//	for i := 0; i < len(board); i++ {
//		for j := 0; j < len(board[0]); j++ {
//			if i == x || j == y {
//				all[board[i][j]] = false
//				continue
//			}
//			xs := (x / 3) * 3
//			ys := (y / 3) * 3
//			if i >= xs && i <= xs+2 && j >= ys && j <= ys+2 {
//				all[board[i][j]] = false
//				continue
//			}
//		}
//	}
//	outPut := []byte{}
//	for k, v := range all {
//		if v {
//			outPut = append(outPut, k)
//		}
//	}
//	return outPut
//}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Sudoku Solver.
func solveSudoku(board [][]byte) {
	var colMark [9][9]bool
	var rowMark [9][9]bool
	var areaMark [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			colMark[i][j] = true
			rowMark[i][j] = true
			areaMark[i][j] = true
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j]-'0') - 1
				rowMark[i][num] = false
				colMark[j][num] = false
				index := (i/3)*3 + j/3
				areaMark[index][num] = false
			}
		}
	}

	findSolved(board, 0, colMark, rowMark, areaMark)
}

func findSolved(board [][]byte, count int, colMark, rowMark, areaMark [9][9]bool) bool {
	if count >= 81 {
		return true
	}
	row := count / 9
	col := count % 9
	area := (row/3)*3 + col/3
	if board[row][col] == '.' {
		for v := 0; v < 9; v++ {
			if colMark[col][v] && rowMark[row][v] && areaMark[area][v] {
				board[row][col] = byte('1' + v)
				colMark[col][v] = false
				rowMark[row][v] = false
				areaMark[area][v] = false
				if findSolved(board, count+1, colMark, rowMark, areaMark) {
					return true
				}
				board[row][col] = '.'
				colMark[col][v] = true
				rowMark[row][v] = true
				areaMark[area][v] = true
			}
		}
	} else {
		return findSolved(board, count+1, colMark, rowMark, areaMark)
	}

	return false
}

func main() {
	var data = [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	for i := 0; i < len(data); i++ {
		fmt.Printf("%v:", i)
		for j := 0; j < len(data[i]); j++ {
			fmt.Printf("  %v", string(data[i][j]))
			if j == 2 || j == 5 {
				fmt.Printf("  |")
			}
		}
		fmt.Printf("\n")
		if i == 2 || i == 5 {
			fmt.Printf("—  —  —  —  —  —  —  —  —  —  —  ———\n")
		}
	}
	fmt.Printf("--------------------------------------\n")
	solveSudoku(data)
	for i := 0; i < len(data); i++ {
		fmt.Printf("%v:", i)
		for j := 0; j < len(data[i]); j++ {
			fmt.Printf("  %v", string(data[i][j]))
			if j == 2 || j == 5 {
				fmt.Printf("  |")
			}
		}
		fmt.Printf("\n")
		if i == 2 || i == 5 {
			fmt.Printf("—  —  —  —  —  —  —  —  —  —  —  ———\n")
		}
	}
}
