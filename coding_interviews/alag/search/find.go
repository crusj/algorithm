package search

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
//  从左上角开始搜索会产生分叉，即使尽力减小搜索范围，还是会超时
func Find(target int, array [][]int) bool {
	width := len(array[0])
	depth := len(array)

	// write code here
	var tf func(array [][]int, x, y int) bool
	tf = func(array [][]int, x, y int) bool {
		if x >= width || y >= depth {
			return false
		}

		if array[y][x] == target {
			return true
		}

		if array[y][x] < target {
			if array[y][width-1] < target && array[depth-1][x] < target {
				return tf(array, x+1, y+1)
			}

			if array[y][width-1] < target {
				return tf(array, x, y+1)
			}

			if array[depth-1][x] < target {
				return tf(array, x+1, y)
			}

			return tf(array, x+1, y) || tf(array, x, y+1)
		}

		return false
	}

	return tf(array, 0, 0)
}

// 从左下开始搜索，小则向右大则向上，搜索不出现分叉
func Find2(target int, array [][]int) bool {
	// write code here
	var tf func(array [][]int, width, depth int) bool
	tf = func(array [][]int, width, depth int) bool {
		if width >= len(array[0]) || depth < 0 {
			return false
		}

		if array[depth][width] == target {
			return true
		}

		if array[depth][width] < target { // 向右
			return tf(array, width+1, depth)
		} else {
			return tf(array, width, depth-1) // 向上
		}
	}

	return tf(array, 0, len(array)-1)
}
