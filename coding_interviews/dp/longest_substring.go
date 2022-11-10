package dp

import (
	"github.com/crusj/algorithm/common"
)

// lengthOfLongestSubstring function  
// dp[n] 为长度为n的字符串最长不含重复字符的子字符串长度.
// base
// dp[1] = 1, dp[0] = 0
// dp[n] = max(dp[n - 1] + 1, dp[n - 1]) // 包含s[n]此时必须保证dp[n-1]包含最后一个字符在子串中才能保证连续，或者不包含s[n]
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	sb := []byte(s)

	dp := make([]int, len(sb))
	// base
	dp[0] = 1

	// 第二个字符开始
	for i := 1; i < len(sb); i++ {
		// 包含第i个字符
		tmp := make(map[byte]struct{}, 0)
		for j := i; j >= 0; j-- {
			if _, exists := tmp[sb[j]]; exists {
				break
			}
			tmp[sb[j]] = struct{}{}
		}

		// 不包含第i个字符
		dp[i] = common.Max(len(tmp), dp[i-1])
	}

	return dp[len(sb)-1]
}
