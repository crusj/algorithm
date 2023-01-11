package leetcode

import "sort"

/*
输入：tokens = [100,200,300,400], power = 200
输出：2
解释：按下面顺序使用令牌可以得到 2 分：
1. 令牌 0 正面朝上，能量变为 100 ，分数变为 1
2. 令牌 3 正面朝下，能量变为 500 ，分数变为 0
3. 令牌 1 正面朝上，能量变为 300 ，分数变为 1
4. 令牌 2 正面朝上，能量变为 0 ，分数变为 2
*/
// 感觉像dp
func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	return fn(tokens, power, 0, len(tokens)-1)
}

// 能量换分数
// 分数换能量
func fn(tokens []int, power int, i, j int) int {
	var a, b int
	// 获取更大能量
	if i < j && tokens[i] < tokens[j] && power >= tokens[i] {
		a = fn(tokens, power-tokens[i]+tokens[j], i+1, j-1)
	}

	// 获取更大的分数
	if i <= j && tokens[i] <= power {
		b = fn(tokens, power-tokens[i], i+1, j) + 1
	}

	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
