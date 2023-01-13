package leetcode

import "sort"

func longestPalindrome(s string) int {
	sl := []byte(s)
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	ret := 0
	i, j := 0, 0
	for i < len(sl) {
		for j < len(sl) && sl[j] == sl[i] {
			j++
		}

		diff := j - i

		if diff%2 == 0 {
			ret += diff
		} else {
			ret += diff - 1
		}

		i = j
	}

	if ret < len(sl) {
		ret += 1
	}

	return ret
}
