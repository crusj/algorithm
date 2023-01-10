package leetcode

func minimumCardPickup(cards []int) int {
	if len(cards) == 0 {
		return -1
	}

	i, j, w, s := 0, 0, make(map[int]int), 0
	min := 999999

	for j < len(cards) {
		s++
		if _, exists := w[cards[j]]; !exists {
			w[cards[j]] = 1
		} else {
			w[cards[j]] += 1
		}

		for len(w)+1 == s {
			if j-i+1 < min {
				min = j - i + 1
			}

			s--
			w[cards[i]] -= 1
			if w[cards[i]] == 0 {
				delete(w, cards[i])
			}
			i++
		}

		j++
	}

	if min == 999999 {
		return -1
	}

	return min
}
