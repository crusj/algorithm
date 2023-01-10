package leetcode

import "sort"

// [1,2] 3
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i, j := 0, len(people)-1
	ret := 0
	for i <= j {
		for j > i && (people[i]+people[j]) > limit {
			ret++
			j--
		}

		ret++
		i++
		j--
	}

	return ret
}
