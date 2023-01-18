package leetcode

import (
	"sort"
	"strconv"
)

func minimumSum(num int) int {
	numStr := []byte(strconv.Itoa(num))
	sort.Slice(numStr, func(i, j int) bool {
		return numStr[i] < numStr[j]
	})

	a, b := 0, 0
	if numStr[0] == '0' {
		a, _ = strconv.Atoi(string(numStr[3]))
	} else {
		a, _ = strconv.Atoi(string(numStr[0]) + string(numStr[3]))
	}
	if numStr[1] == '0' {
		b, _ = strconv.Atoi(string(numStr[2]))
	} else {
		b, _ = strconv.Atoi(string(numStr[1]) + string(numStr[2]))
	}

	return a + b
}
