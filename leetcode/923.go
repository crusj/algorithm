package leetcode

import (
	"sort"
)

// 给定一个整数数组 arr ，以及一个整数 target 作为目标值，返回满足 i < j < k 且 arr[i] + arr[j] + arr[k] == target 的元组 i, j, k 的数量。
// 由于结果会非常大，请返回 109 + 7 的模。

/*
输入：arr = [1,1,2,2,3,3,4,4,5,5], target = 8
输出：20
解释：
按值枚举(arr[i], arr[j], arr[k])：
(1, 2, 5) 出现 8 次；
(1, 3, 4) 出现 8 次；
(2, 2, 4) 出现 2 次；
(2, 3, 3) 出现 2 次。
*/
func threeSumMulti(arr []int, target int) int {
	ret := 0
	sort.Ints(arr)
	for i, v := range arr {
		diff := target - v
		a, b := i+1, len(arr)-1
		for a < b {
			if diff == arr[a]+arr[b] {
				if arr[a] == arr[b] {
					t := b - a + 1
					ret += t * (t - 1) / 2
					break
				}

				x, y := 1, 1
				for a+1 < b && arr[a] == arr[a+1] {
					x++
					a++
				}
				for b-1 > a && arr[b] == arr[b-1] {
					y++
					b--
				}

				a++
				b--
				ret += x * y
			} else if diff > arr[a]+arr[b] {
				a++
			} else {
				b--
			}
		}
	}

	return ret
}
