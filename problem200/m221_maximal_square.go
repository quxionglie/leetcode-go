package problem200

/*
*
221. 最大正方形
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

示例 1：
输入：matrix = [
["1","0","1","0","0"],
["1","0","1","1","1"],
["1","1","1","1","1"],
["1","0","0","1","0"]]
输出：4
示例 2：

输入：matrix = [
["0","1"],
["1","0"]]
输出：1
示例 3：

输入：matrix = [["0"]]
输出：0
*/
func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide
	}
	rows, columns := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				maxSide = max(maxSide, 1)
				curMaxSide := min(rows-i, columns-j)
				for k := 1; k < curMaxSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for m := 0; m < k; m++ {
						if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
