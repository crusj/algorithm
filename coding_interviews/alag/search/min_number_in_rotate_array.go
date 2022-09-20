package search

/**
 *
 * @param rotateArray int整型一维数组
 * @return int整型
 */
func minNumberInRotateArray(rotateArray []int) int {
	if len(rotateArray) == 0 {
		return -1
	}

	var tf func(rotateArray []int, i, j int) int
	tf = func(rotateArray []int, i, j int) int {
		if i == j {
			return rotateArray[i]
		}

		mid := (i + j) / 2
		if rotateArray[mid] > rotateArray[j] {
			return tf(rotateArray, mid+1, j)
		} else if rotateArray[mid] < rotateArray[j] {
			return tf(rotateArray, i, mid)
		} else {
			return tf(rotateArray, i, j-1)
		}
	}

	return tf(rotateArray, 0, len(rotateArray)-1)
}
