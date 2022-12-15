package leetcode

func minSwaps(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	size := 0
	for _, v := range nums {
		if v == 1 {
			size++
		}
	}
	if size == 0 {
		return 0
	}

	nums = append(nums, nums...)
	i, j, zt := 0, 0, 0
	min := 99999

	for j < len(nums) {
		if nums[j] == 0 {
			zt++
		}

		for j-i+1 > size {
			if nums[i] == 0 {
				zt--
			}
			i++
		}

		if j-i+1 == size {
			if zt < min {
				min = zt
			}
		}

		j++
	}

	return min
}
