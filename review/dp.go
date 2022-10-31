package review

import (
	"math"
)

func FindGreatestSumOfSubArray(array []int) int {
	ret := math.MinInt64
	for i := 0; i < len(array); i++ {
		tmp := 0
		for j := i; j < len(array)-i; j++ {
			tmp += array[j]
		}
		if tmp > ret {
			ret = tmp
		}
	}

	return ret
}

// 回溯法
// 选择范围+有效选择+回退
func MinSumPath(paths [][]int) int {
	min := math.MaxInt64

	var tf func(i, s int, path []int)
	tf = func(i, s int, path []int) {
		if len(path) == len(paths) { // 计算最小值
			tmp := 0
			for _, v := range path {
				tmp += v
			}
			if tmp < min {
				min = tmp
			}

			return
		}

		selects := make([]int, 0)
		if i == 0 {
			for x := 0; x < len(paths[i]); x++ {
				selects = append(selects, x)
			}
		} else {
			// 左下
			if s-1 >= 0 {
				selects = append(selects, s-1)
			}
			// 正下
			selects = append(selects, s)

			// 右下
			if s+1 < len(paths[i]) {
				selects = append(selects, s+1)
			}
		}

		for _, item := range selects {
			path = append(path, paths[i][item])
			tf(i+1, item, path)
			path = path[0 : len(path)-1]
		}
	}

	tf(0, 0, nil)
	return min
}

func min(args ...int) int {
	if len(args) == 1 {
		return args[0]
	}

	a := args[0]
	b := min(args[1:]...)
	if a > b {
		return b
	}

	return a
}

// 动态规划
func MinFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	res := math.MaxInt64
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i < 0 || j > len(matrix)-1 || j < 0 {
			return math.MaxInt64
		}

		if i == 0 {
			return matrix[0][j]
		}

		return matrix[i][j] + min(dp(i-1, j), dp(i-1, j-1), dp(i-1, j+1))
	}

	l := len(matrix)
	for i := 0; i < len(matrix[0]); i++ {
		res = min(res, min(dp(l-1, i-1), dp(l-1, i), dp(l-1, i+1)))
	}

	return res
}

func max(args ...int) int {
	if len(args) == 1 {
		return args[0]
	}

	tmp := max(args[1:]...)
	if tmp > args[0] {
		return tmp
	}

	return args[0]
}

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid) == 1 {
		tmp := 0
		for i := 0; i < len(grid[0]); i++ {
			tmp += grid[0][i]
		}

		return tmp
	}

	var vars [200][200]*int

	// write code here
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if vars[i][j] != nil {
			return *vars[i][j]
		}

		if i == 0 {
			tmp := 0
			for j >= 0 {
				tmp += grid[0][j]
				j--
			}
			return tmp
		}

		if j == 0 {
			tmp := 0
			for i >= 0 {
				tmp += grid[i][0]
				i--
			}

			return tmp
		}

		t := grid[i][j] + max(dp(i-1, j), dp(i, j-1))
		vars[i][j] = &t

		return t
	}

	return dp(len(grid)-1, len(grid[0])-1)
}
