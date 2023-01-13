package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	ret := 0
	i, j := 0, 0

	for j < len(s) && i < len(g) {
		for i < len(g) && g[i] > s[j] {
			i++
		}

		if i == len(g) {
			break
		}

		i++
		j++
		ret++
	}
	return ret
}
