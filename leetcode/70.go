package leetcode

// dp(n) = dp(n-1) + dp(n-2)
// dp(n) = 1 // n = 1
// dp(n) = 2 // n = 2
// 最优子结构
func climbStairs(n int) int {
	i, j := 1, 2
	if n == i {
		return i
	}
	if n == j {
		return j
	}

	for x := 3; x <= n; x++ {
		i, j = j, i+j
	}

	return j
}
