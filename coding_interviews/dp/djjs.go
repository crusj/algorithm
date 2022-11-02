package dp

import "github.com/crusj/algorithm/common"

// 打家劫舍
// 相邻不能偷，求最大金额
// dp[n] = max(dp[n-1], dp[n-2] + nums[n])
func djjs(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 0, nums[0]

	for i := 2; i <= len(nums); i++ {
		dp[i] = common.Max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}
