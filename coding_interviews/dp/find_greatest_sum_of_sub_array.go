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
	return 0
}
