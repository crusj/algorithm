package leetcode

func getAverages(nums []int, k int) []int {
	ret := make([]int, 0, len(nums))
	for range nums {
		ret = append(ret, -1)
	}

	size := 2*k + 1

	if len(nums) < size {
		return ret
	}

	i, j, sum := 0, 0, 0
	avgs := make([]int, 0, len(nums)-size+1)

	for j < len(nums) {
		sum += nums[j]
		if j-i+1 == size {
			avgs = append(avgs, sum/size)
			sum -= nums[i]
			i++
		}
		j++
	}

	p := k
	for _, v := range avgs {
		ret[p] = v
		p++
	}

	return ret
}
