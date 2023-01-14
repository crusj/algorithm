package leetcode

/*
输入：bills = [5,5,10,10,20]
输出：false
解释：
前 2 位顾客那里，我们按顺序收取 2 张 5 美元的钞票。
对于接下来的 2 位顾客，我们收取一张 10 美元的钞票，然后返还 5 美元。
对于最后一位顾客，我们无法退回 15 美元，因为我们现在只有两张 10 美元的钞票。
由于不是每位顾客都得到了正确的找零，所以答案是 false。
*/
func lemonadeChange(bills []int) bool {
	m := map[int]int{
		5:  0,
		10: 0,
	}
	for i := 0; i < len(bills); i++ {
		diff := bills[i] - 5
		if diff == 0 {
			m[5]++
		} else if diff == 5 {
			if m[5] == 0 {
				return false
			}
			m[5]--
			m[10]++
		} else if diff == 15 {
			if m[10] > 0 && m[5] > 0 {
				m[10]--
				m[5]--
				continue
			}

			if m[5] >= 3 {
				m[5] -= 3
				continue
			}

			return false
		}
	}

	return true
}
