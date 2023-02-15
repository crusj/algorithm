package leetcode

// dp(n) = dp(n-1) + dp(n-2)
// dp(n) = 1 // n = 1
// dp(n) = 2 // n = 2
// 最优子结构
func climbStairs(n int) int {
	dp := make(map[int]int, n)
	dp[1] = 1
	dp[2] = 2

	var fn func(n int) int
	fn = func(n int) int {
		if dp[n] > 0 {
			return dp[n]
		}

		dp[n] = fn(n-1) + fn(n-2)
		return dp[n]
	}

	return fn(n)
}
