package leetcode

import "strconv"

func convertTime(current string, correct string) int {
	ans := 0
	a, b := convertTimeStr2Minutes(current), convertTimeStr2Minutes(correct)
	if a > b {
		return -1
	}

	p := []int{60, 15, 5, 1}
	i := 0

	for {
		if a+p[i] == b {
			ans++
			return ans
		}

		if a+p[i] < b {
			ans++
			a += p[i]
			continue
		}

		i++
	}
}

func convertTimeStr2Minutes(t string) int {
	i := 0
	for t[i] == '0' {
		i++
	}

	j := i + 1
	for t[j] != ':' {
		j++
	}

	hours, _ := strconv.Atoi(t[i:j])

	i = j + 1
	for t[i] == '0' {
		i++
	}
	mins, _ := strconv.Atoi(t[i:])
	return hours*60 + mins
}
