package backTracking

import "strings"

func solveNQueens(n int) [][]string {
	var result [][]string

	//初始化二维数组，即初始化棋盘
	chessbd := make([][]string, n)
	for i := 0; i < n; i++ {
		chessbd[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessbd[i][j] = "."
		}
	}

	var _deth func(int)
	_deth = func(row int) {
		//递归终止条件
		if row == n {
			//数据填装到返回结果中
			r := make([]string, n)
			for i, rowStr := range chessbd {
				r[i] = strings.Join(rowStr, "")
			}
			result = append(result, r)
			return
		}

		//对当前一行中的每一列进行放皇后
		for i := 0; i < n; i++ {
			//判断当前位置是否合法
			if isVaild(row, n, i, chessbd) {
				//中当前位置放置皇后，然后递归进入下一层
				chessbd[row][i] = "Q"
				_deth(row + 1)
				//当前递归遍历到底后要继续回撤，进入本行队其他列中国
				chessbd[row][i] = "."
			}
		}
	}
	_deth(0)
	return result
}

func isVaild(row int, n int, col int, chessbd [][]string) bool {
	//判断棋盘当前列是否有皇后
	for i := 0; i < row; i++ {
		if chessbd[i][col] == "Q" {
			return false
		}
	}

	//向后判断45度位置是否皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessbd[i][j] == "Q" {
			return false
		}
	}

	//向后判断135度方向位置是否存在皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessbd[i][j] == "Q" {
			return false
		}
	}

	return true
}
