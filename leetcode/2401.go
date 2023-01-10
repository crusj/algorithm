package leetcode

func longestNiceSubarray(nums []int) int {
	i, j, w := 0, 0, make(map[int]int)
	max := 0

	for j < len(nums) {
		if _, exist := w[nums[j]]; !exist {
			w[nums[j]] = 1
		} else {
			w[nums[j]]++
		}

		for !isElegant(nums, i, j) {
			w[nums[i]]--
			i++
		}

		if j-i+1 > max {
			max = j - i + 1
		}

		j++
	}

	return max
}

func isElegant(nums []int, i, j int) bool {
	for i < j {
		if nums[j]&nums[i] != 0 {
			return false
		}
		i++
	}

	return true
}
