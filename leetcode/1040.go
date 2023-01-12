package leetcode

import (
	"fmt"
	"sort"
)

/*
输入：[7,4,9]
输出：[1,2]
解释：
我们可以移动一次，4 -> 8，游戏结束。
或者，我们可以移动两次 9 -> 5，4 -> 6，游戏结束。
*/
// 4 7 9

// 不能在移动 最小值不能再移动 max - next + 1 = len(arr) - 1，最大值不能在移动 next - min + 1 = len(arr) - 1
func numMovesStonesII(stones []int) []int {
	ret := make([]int, 2)
	sort.Ints(stones)

	copy1 := make([]int, len(stones))
	copy2 := make([]int, len(stones))
	copy(copy1, stones)
	copy(copy2, stones)

	l := len(stones)
	for {
		sort.Ints(copy1)
		left := copy1[l-1] - copy1[1] + 1
		right := copy1[l-2] - copy1[0] + 1
		fmt.Println(l, left, right, copy1)
		if left == l-1 && right == l-1 {
			break
		} else if (left == l-1 && right != l-1) || (left != l-1 && right != l-1 && right < left) {
			// 移动右边
			t := l - 2
			for t > 0 {
				if copy1[t-1] < copy1[t]-1 {
					copy1[l-1] = copy1[t] - 1
					break
				}
				t--
			}

		} else if (right == l-1 && left != l-1) || (left != l-1 && right != l-1 && right > left) {
			t := 1
			for t < l-1 {
				if copy1[t+1] > copy1[t]+1 {
					copy1[0] = copy1[t] + 1
					break
				}
				t++
			}

		}

		ret[0]++
	}

	for {
		sort.Ints(copy2)
		left := copy2[l-1] - copy2[1] + 1
		right := copy2[l-2] - copy2[0] + 1
		fmt.Println(l, left, right, copy2)
		if left == l-1 && right == l-1 {
			break
		} else if (left == l-1 && right != l-1) || (left != l-1 && right != l-1 && right > left) {
			// 移动右边
			t := l - 2
			for t > 0 {
				if copy2[t-1] < copy2[t]-1 {
					copy2[l-1] = copy2[t] - 1
					break
				}
				t--
			}

		} else if (right == l-1 && left != l-1) || (left != l-1 && right != l-1 && right < left) {
			t := 1
			for t < l-1 {
				if copy2[t+1] > copy2[t]+1 {
					copy2[0] = copy2[t] + 1
					break
				}
				t++
			}

		}

		ret[1]++
	}

	return ret
}
