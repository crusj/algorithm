package leetcode

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for {
		if i >= j {
			return true
		}
		if s[i] == s[j] {
			i++
			j--
			continue
		}

		a, b := i+1, j
		for {
			if a >= b {
				return true
			}
			if s[a] != s[b] {
				break
			}
			a++
			b--
		}

		m, n := i, j-1
		for {
			if m >= n {
				return true
			}
			if s[m] != s[n] {
				return false
			}
			m++
			n--
		}
	}
}
