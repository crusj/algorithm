package leetcode

func camelMatch(queries []string, pattern string) []bool {
	ret := make([]bool, len(queries))
loop:
	for i, query := range queries {
		ret[i] = false
		j := 0
		for _, b := range []byte(pattern) {
			for j < len(query) {
				if query[j] >= 'A' && query[j] <= 'Z' && query[j] != b {
					continue loop
				}

				if query[j] == b {
					break
				}

				j++
			}

			if j == len(query) {
				continue loop
			}
			j++
		}

		for j < len(query) && query[j] >= 97 && query[j] <= 122 {
			j++
		}
		if j < len(query) {
			continue
		}

		ret[i] = true
	}

	return ret
}
