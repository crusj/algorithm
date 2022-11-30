package simulation

//
// dp(n,m)代表把长度为n的绳子剪成m段，每段乘积的最大值
// dp(n,m) = n - 1 // m == 2,至少两段
// dp(n,m) = n  // n == m
// dp(n,m) = max(dp(n-1, m - 1) * 1, dp(n-2, m -1) * 2, dp(n-3, m-1) * 3, dp(n-k, m -1) * k) // (n - k >= m - 1) 不能一下减没了，剩下的段数和长度至少相等才能分到平均分到一
func cutRope(number int) int {
	dp := make([][]int, number+1)
	for i := range dp {
		dp[i] = make([]int, number+1)
	}

	var fn func(n, m int) int
	fn = func(n, m int) int {
		if dp[n][m] > 0 {
			return dp[n][m]
		}

		if m == 2 { // 最小分为两段
			return n - 1
		}

		if n == m {
			return n
		}

		max := n
		for k := 1; n-k >= m-1; k++ {
			v := fn(n-k, m-1) * k
			if v > max {
				max = v
			}
		}
		dp[n][m] = max
		return max
	}

	max := 1
	for i := 2; i <= number; i++ {
		tmp := fn(number, i)
		if max < tmp {
			max = tmp
		}
	}
	return max
}
