package leetcode

import (
	"sort"
)

/*
有 N 位扣友参加了微软与力扣举办了「以扣会友」线下活动。主办方提供了 2*N 道题目，整型数组 questions 中每个数字对应了每道题目所涉及的知识点类型。
若每位扣友选择不同的一题，请返回被选的 N 道题目至少包含多少种知识点类型。

[1,2,3,1]
*/

// 哈希+排序+贪心
func halfQuestions(questions []int) int {
	m := map[int]int{}
	s := []int{}

	for _, item := range questions {
		m[item]++
		if m[item] == 1 {
			s = append(s, item)
		}
	}

	sort.Slice(s, func(i, j int) bool {
		return m[s[i]] > m[s[j]]
	})

	ans := 0
	t := len(questions) / 2
	for _, index := range s {
		ans++
		if m[index] >= t {
			return ans
		}

		t = t - m[index]
	}

	return ans
}
