package leetcode

import "sort"

/*
输入：words = ["a","b","ba","bca","bda","bdca"]
输出：4
解释：最长单词链之一为 ["a","ba","bda","bdca"]
*/
func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	dp := make([]int, len(words))
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if fnn(words[j], words[i]) {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	max := dp[0]
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

func fnn(prev, next string) bool {
	if len(prev)+1 != len(next) {
		return false
	}

	j := 0
	for _, ch1 := range []byte(prev) {
		for j < len(next) && ch1 != next[j] {
			j++
		}

		if j == len(next) {
			return false
		}
	}

	return true
}
