package leetcode

func maxDistance(colors []int) int {
	i, j := 0, len(colors)-1
	l, r := colors[i], colors[j]
	if l != r {
		return len(colors) - 1
	}

	for i < len(colors)-1 && colors[i] == r {
		i++
	}
	if i == len(colors)-1 {
		return 0
	}

	for j >= 0 && colors[j] == l {
		j--
	}
	if j == 0 {
		return 0
	}

	if len(colors)-i > j+1 {
		return len(colors) - i - 1
	}
	return j
}
