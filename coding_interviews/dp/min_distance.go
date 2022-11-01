package dp

import "github.com/crusj/algorithm/common"

// 最小距离
// dp
// 将word1转化为word2的最小距离
// dp[i][j]  word1长度为i，word2长度为j时需要次数
func minDistance(word1, word2 string) int {
	dp := make([][]int, len(word1))

	// base
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2), len(word2))
	}

	dp[0][0] = 0
	for j := 1; j < len(word2); j++ {
		dp[0][j] = dp[0][j-1] + 1
	}
	for i := 1; i < len(word1); i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	// dp
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			dp[i][j] = common.Min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
		}
	}

	return dp[len(word1)-1][len(word2)-1]
}
