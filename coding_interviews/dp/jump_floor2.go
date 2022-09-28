package dp

var dp = make(map[int]int)

func jumpFloorII(number int) int {
	// write code here
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

	ret := 1
	for i := number - 1; i > 0; i-- {
		ret += jumpFloorII(i)
	}

	dp[number] = ret

	return ret
}
