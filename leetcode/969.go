package leetcode

/*
输入：[3,2,4,1]
输出：[4,2,4,3]
解释：
我们执行 4 次煎饼翻转，k 值分别为 4，2，4，和 3。
初始状态 arr = [3, 2, 4, 1]
第一次翻转后（k = 4）：arr = [1, 4, 2, 3]
第二次翻转后（k = 2）：arr = [4, 1, 2, 3]
第三次翻转后（k = 4）：arr = [3, 2, 1, 4]
第四次翻转后（k = 3）：arr = [1, 2, 3, 4]，此时已完成排序。
*/
/*
3
4231

4
1324

2
3124

3
2134

1
1

2
1234

*/

func pancakeSort(arr []int) []int {
	path := make([]int, 0)
	help(arr, len(arr)-1, &path)
	return path
}

func help(arr []int, j int, path *[]int) {
	if j == 0 {
		return
	}

	maxIndex := 0
	i := 0
	for i <= j {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
		i++
	}
	// 最大值是最后一个跳过
	if maxIndex == j {
		help(arr, j-1, path)
		return
	}

	a, b := 0, maxIndex
	*path = append((*path), maxIndex+1)
	// switch
	for {
		if a >= b {
			break
		}
		arr[a], arr[b] = arr[b], arr[a]
		a++
		b--
	}

	a, b = 0, j
	*path = append((*path), j+1)
	for {
		if a >= b {
			break
		}
		arr[a], arr[b] = arr[b], arr[a]
		a++
		b--
	}

	help(arr, j-1, path)
}
