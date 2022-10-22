package review

import "math"

func FindGreatestSumOfSubArray(array []int) int {
	ret := math.MinInt64
	for i := 0; i < len(array); i++ {
		tmp := 0
		for j := i; j < len(array)-i; j++ {
			tmp += array[j]
		}
		if tmp > ret {
			ret = tmp
		}
	}

	return ret
}
