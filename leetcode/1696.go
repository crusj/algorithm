package leetcode

func maxResult(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	MIM := make(map[int]int)
	max := 0

	var fn func(index int, sum int)
	fn = func(index int, sum int) {
		if index >= len(nums) {
			return
		}

		// MIM is used to keep track of the max value.
		if _, exist := MIM[index]; !exist {
			MIM[index] = sum
		} else {
			if sum <= MIM[index] {
				return
			}

			if sum > MIM[index] {
				MIM[index] = sum
			}
		}

		if index == len(nums)-1 {
			if sum > max {
				max = sum
			}

			return
		}

		for j := index + 1; j <= k+index && j < len(nums); j++ {
			fn(j, sum+nums[j])
		}
	}

	fn(0, nums[0])
	return max
}
