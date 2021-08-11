package isValidSudoku

import "strconv"

func isValidSudoku(board [][]string) bool  {
	var row [9][10]int
	var col [9][10]int
	var box [9][10]int
	for i := 0; i < 9;i++ {
		for j := 0; j < 9;j++ {
			if board[i][j] == "." {
				continue
			}
			curNum := toInt(board[i][j]) - toInt("0")
			// 判斷所在的行是否有該數字出現
			if row[i][curNum] == 1 {
				return false
			}
			// 判斷所在列是否出現該數字
			if col[j][curNum] == 1 {
				return false
			}
			if box[ j / 3 + (i / 3) * 3][curNum] == 1 {
				return false
			}
			row[i][curNum] = 1
			col[j][curNum] = 1
			box[j / 3 + (i / 3) * 3][curNum] = 1
		}
	}
	return true
}

func toInt(s string)int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}

