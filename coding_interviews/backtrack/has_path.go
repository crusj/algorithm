package backtrack

func hasPath(matrix [][]byte, word string) bool {
	if len(matrix) == 0 {
		return false
	}
	// write code here

	var fn func(x, y int, used [][]int, path []byte) bool
	fn = func(x, y int, used [][]int, path []byte) bool {
		if len(path) == len(word) {
			print(string(path))
			return true
		}

		// 上
		if x-1 >= 0 && used[x-1][y] == 0 && matrix[x-1][y] == word[len(path)] {
			used[x-1][y] = 1
			path = append(path, matrix[x-1][y])
			if fn(x-1, y, used, path) {
				return true
			}

			used[x-1][y] = 0
			path = path[0 : len(path)-1]
		}
		// 下
		if x+1 < len(matrix) && used[x+1][y] == 0 && matrix[x+1][y] == word[len(path)] {
			used[x+1][y] = 1
			path = append(path, matrix[x+1][y])
			if fn(x+1, y, used, path) {
				return true
			}

			used[x+1][y] = 0
			path = path[0 : len(path)-1]
		}
		// 左
		if y-1 >= 0 && used[x][y-1] == 0 && matrix[x][y-1] == word[len(path)] {
			used[x][y-1] = 1
			path = append(path, matrix[x][y-1])
			if fn(x, y-1, used, path) {
				return true
			}

			used[x][y-1] = 0
			path = path[0 : len(path)-1]
		}
		// 右
		if y+1 < len(matrix[0]) && used[x][y+1] == 0 && matrix[x][y+1] == word[len(path)] {
			used[x][y+1] = 1
			path = append(path, matrix[x][y+1])
			if fn(x, y+1, used, path) {
				return true
			}

			used[x][y+1] = 0
			path = path[0 : len(path)-1]
		}

		return false
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != word[0] {
				continue
			}
			path := []byte{
				matrix[i][j],
			}
			tmp := make([][]int, len(matrix))
			for i := range tmp {
				tmp[i] = make([]int, len(matrix[0]))
			}

			if fn(i, j, tmp, path) {
				return true
			}
		}
	}

	return false
}
