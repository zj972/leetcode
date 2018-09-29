package main

import "fmt"

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Valid Sudoku.
//func isValidSudoku(board [][]byte) bool {
//	colMap := make([]map[string]bool, 9)
//	divMap := make([][]map[string]bool, 3)
//	divMap[0] = make([]map[string]bool, 3)
//	divMap[1] = make([]map[string]bool, 3)
//	divMap[2] = make([]map[string]bool, 3)
//	for i := 0; i < len(board); i++ {
//		rowMap := make(map[string]bool)
//		for j := 0; j < 9; j++ {
//			if string(board[i][j]) == "." {
//				continue
//			}
//
//			if rowMap[string(board[i][j])] == true {
//				return false
//			} else {
//				rowMap[string(board[i][j])] = true
//			}
//
//			if colMap[j] == nil {
//				colMap[j] = make(map[string]bool)
//			}
//			if colMap[j][string(board[i][j])] == true {
//				return false
//			} else {
//				colMap[j][string(board[i][j])] = true
//			}
//
//			var x, y int
//			switch i {
//			case 0, 1, 2:
//				x = 0
//			case 3, 4, 5:
//				x = 1
//			case 6, 7, 8:
//				x = 2
//			}
//			switch j {
//			case 0, 1, 2:
//				y = 0
//			case 3, 4, 5:
//				y = 1
//			case 6, 7, 8:
//				y = 2
//			}
//			if divMap[x][y] == nil {
//				divMap[x][y] = make(map[string]bool)
//			}
//			if divMap[x][y][string(board[i][j])] == true {
//				return false
//			} else {
//				divMap[x][y][string(board[i][j])] = true
//			}
//		}
//	}
//	return true
//}

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Valid Sudoku.
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				// Check rows
				for k := i + 1; k < len(board); k++ {
					if board[k][j] == board[i][j] {
						return false
					}
				}
				// Check cols
				for k := j + 1; k < len(board[0]); k++ {
					if board[i][k] == board[i][j] {
						return false
					}
				}

				// Check sub
				brow := i / 3
				bcol := j / 3
				for k := 0; k < 3; k++ {
					for c := 0; c < 3; c++ {
						row := 3*brow + k
						col := 3*bcol + c

						if row == i && col == j {
							continue
						} else if board[row][col] == board[i][j] {
							return false
						}
					}
				}
			}
		}
	}
	return true
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
	fmt.Println(isValidSudoku(data))
}
