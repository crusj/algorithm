package leetcode

import "sort"

func minimumCost(cost []int) int {
	sort.Ints(cost)
	ans := 0

	j := len(cost) - 1
	for j >= 0 {
		ans += cost[j]
		if j-1 >= 0 {
			ans += cost[j-1]
		}

		j -= 3
	}

	return ans
}
