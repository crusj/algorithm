package leetcode

import (
	"sort"
)

/*
输入：nums = [3,-1,0,2], k = 3
输出：6
解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。

输入：nums = [2,-3,-1,5,-4], k = 2
输出：13
解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。
*/
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	i := 0
	for i < len(nums) && k > 0 {
		if nums[i] < 0 {
			nums[i] *= -1
			i++
			k--
		} else if nums[i] == 0 {
			break
		} else {
			if (k-i)%2 != 0 {
				if i == 0 {
					nums[i] *= -1
					break
				}

				if nums[i-1] < nums[i] {
					nums[i-1] *= -1
				}
			}
			break
		}
	}

	if k > 0 && i == len(nums) && k%2 != 0 {
		nums[i-1] *= -1
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}

	return ans
}
