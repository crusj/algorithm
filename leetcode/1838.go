package leetcode

import (
	"sort"
)

/*
输入：nums = [1,2,4], k = 5
输出：3
解释：对第一个元素执行 3 次递增操作，对第二个元素执 2 次递增操作，此时 nums = [4,4,4] 。
4 是数组中最高频元素，频数是 3 。
*/
func maxFrequency(nums []int, k int) int {
	// if nums is empty, return 0.
	if len(nums) == 0 {
		return 0
	}
	// sort nums.
	sort.Ints(nums)

	i, j, di, max := 0, 0, 0, 0
	for j < len(nums) {
		if j >= 1 && j-1 >= i {
			v := nums[j] - nums[j-1]
			if v > 0 {
				di += (j - i) * v
			}
		}

		for di > k {
			di -= nums[j] - nums[i]
			i++
		}

		if t := j - i + 1; t > max {
			max = t
		}

		j++
	}

	return max
}
