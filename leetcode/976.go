package leetcode

import "sort"

/*
输入：nums = [1,2,1,10]
输出：0
解释：
你不能用边长 1,1,2 来组成三角形。
不能用边长 1,1,10 来构成三角形。
不能用边长 1、2 和 10 来构成三角形。
因为我们不能用任何三条边长来构成一个非零面积的三角形，所以我们返回 0。
// 任意两边大于第三边
*/
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for j := len(nums) - 1; j > 1; j-- {
		if nums[j-1]+nums[j-2] > nums[j] {
			return nums[j-1] + nums[j-2] + nums[j]
		}
	}

	return 0
}
