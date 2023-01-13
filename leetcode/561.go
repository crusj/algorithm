package leetcode

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ret := 0

	j := len(nums) - 2
	for j >= 0 {
		ret += nums[j]
		j = j - 2
	}

	return ret
}
