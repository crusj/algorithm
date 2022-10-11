package dp

// 1.穷举状态
/*
for 天数状态 in 所有天数
    for 持有状态 in [1 持有，0 不持有]
        dp[i][j] = max(选择，选择)
*/
// 2.状态转移方程
/*
dp[i][j] = max(dp[i][0] , dp[i][1])
dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + p[i])
dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - p[i])
*/
// 3.base case
/*
dp[0][0] = 0
dp[0][1] = 0
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// write code here
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, len(prices))
	}

	for i := 0; i < len(prices); i++ {
		// base case
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[0]
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
		// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
