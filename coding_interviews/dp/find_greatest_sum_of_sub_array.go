package dp

import (
	"math"
)

/**
 *
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray(array []int) int {
	// 暴力法
	// return FindGreatestSumOfSubArray1(array)
	return FindGreatestSumOfSubArrayDP(array)
}

// FindGreatestSumOfSubArray1 function  
// 暴力法
func FindGreatestSumOfSubArray1(array []int) int {
	ret := math.MinInt64
	for i := 0; i < len(array); i++ {
		sum := 0
		for j := i; j < len(array); j++ {
			sum += array[j]
			if sum > ret {
				ret = sum
			}
		}
	}

	return ret
}

// FindGreatestSumOfSubArrayDP function  
// 动态规划
func FindGreatestSumOfSubArrayDP(array []int) int {
	dp := make(map[int]int, 0)

	var fn func(i int) int
	fn = func(i int) int {
		if i == 0 {
			return array[i]
		}

		dpi := 0
		if v, exists := dp[i-1]; exists {
			dpi = v
		} else {
			dpi = fn(i - 1)
			dp[i-1] = dpi
		}

		if dpi < 0 {
			return array[i]
		}

		return array[i] + dpi
	}

	max := math.MinInt32
	for i := len(array) - 1; i >= 0; i-- {
		tmp := fn(i)
		if tmp > max {
			max = tmp
		}
	}

	return max
}
