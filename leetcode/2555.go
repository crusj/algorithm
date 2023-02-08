package leetcode

/*
输入：prizePositions = [1,1,2,2,3,3,5], k = 2
输出：7
解释：这个例子中，你可以选择线段 [1, 3] 和 [3, 5] ，获得 7 个奖品。
*/

func maximizeWin(prizePositions []int, k int) int {
	// map存储数字对应的次数.
	m := make(map[int]int)
	// singles存储去重后的prizePositions.
	singles := make([]int, 0)
	for _, v := range prizePositions {
		m[v]++
		if m[v] == 1 {
			singles = append(singles, v)
		}
	}
	return 0
}
