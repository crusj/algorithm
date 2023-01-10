package leetcode

// 直接暴力
func smallestSubarrays(nums []int) []int {
	ret := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		max := nums[i]
		minL := 1
		for j := i + 1; j < len(nums); j++ {
			if tmp := max | nums[j]; tmp > max {
				minL = j - i + 1
				max = tmp
			}
		}
		ret = append(ret, minL)
	}

	return ret
}
