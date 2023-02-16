package leetcode

/*
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。

*/

func getRow(rowIndex int) []int {
	rows := make([][]int, 0)
	rows = append(rows, []int{1}, []int{1, 1})
	for i := 2; i <= rowIndex; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
				continue
			}
			row[j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows = append(rows, row)
	}

	return rows[rowIndex]
}
