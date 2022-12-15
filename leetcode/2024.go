package leetcode

func maxConsecutiveAnswers(answerKey string, k int) int {
	// check answerKey is empty.
	if len(answerKey) == 0 {
		return 0
	}

	i, j, w, max := 0, 0, make(map[byte]int), 0

	for j < len(answerKey) {
		if _, exists := w[answerKey[j]]; !exists {
			w[answerKey[j]] = 1
		} else {
			w[answerKey[j]]++
		}

		for w['F'] > k && w['T'] > k {
			w[answerKey[i]]--
			i++
		}

		if max < j-i+1 {
			max = j - i + 1
		}

		j++
	}

	return max
}
