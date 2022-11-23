package simulation

// jz29 function   顺时针打印矩阵
func jz29(matrix [][]int) []int {
	topLeft := []int{0, 0}
	bottomRight := []int{len(matrix[0]) - 1, len(matrix) - 1}

	ret := []int{}
	for {
		if topLeft[0] > bottomRight[0] || topLeft[1] > bottomRight[1] {
			break
		}

		// x变y不变
		// 2变1不变
		for x := topLeft[0]; x <= bottomRight[0]; x++ {
			ret = append(ret, matrix[topLeft[1]][x])
		}
		if topLeft[1] == bottomRight[1] {
			break
		}

		// x不变y变
		// 2不变1变
		for y := topLeft[1] + 1; y <= bottomRight[1]; y++ {
			ret = append(ret, matrix[y][bottomRight[0]])
		}
		if topLeft[0] == bottomRight[0] {
			break
		}

		// x变y不变
		// 2变1不变
		for x := bottomRight[0] - 1; x >= topLeft[0]; x-- {
			ret = append(ret, matrix[bottomRight[1]][x])
		}

		// x不变y变
		// 2不变1变
		for y := bottomRight[1] - 1; y > topLeft[1]; y-- {
			ret = append(ret, matrix[y][topLeft[0]])
		}

		bottomRight[0]--
		bottomRight[1]--
		topLeft[0]++
		topLeft[1]++

	}
	return ret
}
