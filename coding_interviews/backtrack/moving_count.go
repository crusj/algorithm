package backtrack

func movingCount(threshold int, rows int, cols int) int {
	// write code here
	if rows == 0 || cols == 0 {
		return 0
	}

	var sum func(int) int
	sum = func(a int) int {
		if a/10 > 0 {
			return a%10 + sum(a/10)
		}

		return a
	}

	enable := make([][]int, rows)
	for i := range enable {
		enable[i] = make([]int, cols)
	}
	enable[0][0] = 1

	passed := make([][]int, rows)
	for i := range enable {
		passed[i] = make([]int, cols)
	}

	var fn func(y, x int, used [][]int)
	fn = func(y, x int, used [][]int) {
		if passed[y][x] == 1 {
			return
		}
		passed[y][x] = 1

		// 上
		if y-1 >= 0 && used[y-1][x] == 0 && (sum(y-1)+sum(x)) <= threshold {
			used[y-1][x] = 1
			enable[y-1][x] = 1
			fn(y-1, x, used)
			used[y-1][x] = 0
		}

		// 下
		if y+1 < rows && used[y+1][x] == 0 && (sum(y+1)+sum(x)) <= threshold {
			used[y+1][x] = 1
			enable[y+1][x] = 1
			fn(y+1, x, used)
			used[y+1][x] = 0
		}

		// 左
		if x-1 >= 0 && used[y][x-1] == 0 && (sum(y)+sum(x-1)) <= threshold {
			used[y][x-1] = 1
			enable[y][x-1] = 1
			fn(y, x-1, used)
			used[y][x-1] = 0
		}

		// 右
		if x+1 < cols && used[y][x+1] == 0 && (sum(y)+sum(x+1)) <= threshold {
			used[y][x+1] = 1
			enable[y][x+1] = 1
			fn(y, x+1, used)
			used[y][x+1] = 0
		}
	}

	used := make([][]int, rows)
	for i := range used {
		used[i] = make([]int, cols)
	}
	used[0][0] = 1

	// call
	fn(0, 0, used)

	enableCount := 0
	for i := range enable {
		for j := range enable[i] {
			if enable[i][j] == 1 {
				enableCount++
			}
		}
	}

	return enableCount
}
