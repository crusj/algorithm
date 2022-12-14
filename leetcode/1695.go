package leetcode

func maximumUniqueSubarray(nums []int) int {
	// if nums is empty return 0.
	if len(nums) == 0 {
		return 0
	}

	i, j, w, t := 0, 0, make(map[int]int, 0), 0
	sum := 0
	max := 0

	for j < len(nums) {
		// if w[nums[j]] is exists, increment w[nums[j]].

		if w[nums[j]] == 1 {
			t++
		}

		if t < j-i+1 {
			w[nums[i]]--
			if w[nums[i]] == 0 {
				t--
			}

			sum -= nums[i]
			i++
		}

		if sum > max {
			max = sum
		}

		j++
	}

	return max
}
