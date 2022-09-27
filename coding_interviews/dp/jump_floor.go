package dp

func jumpFloor(number int) int {
	var tf func(number int) int
	dp := make(map[int]int)
	tf = func(number int) int {
		switch number {
		case 0:
			return 0
		case 1:
			return 1
		case 2:
			return 2
		}

		if v, exists := dp[number]; exists {
			return v
		}

		ret := tf(number-1) + tf(number-2)
		dp[number] = ret

		return ret
	}

	return tf(number)
}
