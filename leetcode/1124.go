package leetcode

/*
给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
我们认为当员工一天中的工作小时数大于 8小时的时候，那么这一天就是「劳累的一天」。
所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格大于「不劳累的天数」。
请你返回「表现良好时间段」的最大长度。

输入：hours = [9,9,6,0,6,6,9]
输出：3
解释：最长的表现良好时间段是 [9,9,6]。

输入：hours = [6,6,6]
输出：0

1 <= hours.length <= 104
0 <= hours[i] <= 16
*/

// 暴力窗口穷举，可以
// 滑动窗口，不行
// 贪心连续劳累最多的，不行
// 动态规划，也不行
// 回溯，也不行
// 前缀和
func longestWPI(hours []int) int {
	ans := 0
	// fmt.Println(hours)
	preSum := make([]int, len(hours)+1)
	preSum[0] = 0
	for i := 1; i < len(preSum); i++ {
		if hours[i-1] > 8 {
			hours[i-1] = 1
		} else {
			hours[i-1] = -1
		}
		preSum[i] = preSum[i-1] + hours[i-1]
	}

	j := len(preSum) - 1
	// fmt.Println(" ", hours)
	// fmt.Println(preSum)
	for j > 0 {
		if preSum[j] > 0 {
			ans = max(ans, j)
			j--
			continue
		}

		for i := 0; i < j; i++ {
			if preSum[i] < preSum[j] {
				ans = max(ans, j-i)
				break
			}
		}

		j--
	}

	return ans
}

// [1,3,5,7,9] 1-4
// [0,1,4,9,16,25]
// i,j的和包括i,j
func getSumPrefix(arr []int, i, j int) (ans int) {
	sumPres := make([]int, len(arr)+1)
	sumPres[0] = 0
	for i := range arr {
		sumPres[i+1] = arr[i] + sumPres[i]
	}

	ans = sumPres[j] - sumPres[i-1]
	return
}

// [1,2,3]      3
// [0,1,3,6]
// [1:1,3:1,6:1]
// 连续指数组和的个数
func subarraySum(nums []int, k int) (ans int) {
	preSum := map[int]int{}
	sumi := 0
	preSum[0] = 1
	for i := 0; i < len(nums); i++ {
		sumi += nums[i]
		sumj := sumi - k
		if preSum[sumj] > 0 {
			ans += preSum[sumj]
		}

		preSum[sumi]++
	}

	return ans
}
