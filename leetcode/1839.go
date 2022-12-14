package leetcode

/*
输入：word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
输出：13
解释：最长子字符串是 "aaaaeiiiiouuu" ，长度为 13 。
*/
func longestBeautifulSubstring(word string) int {
	i, j, w, max, k := 0, 0, make(map[byte]int), 0, 0
	f := map[byte][]byte{
		'a': {'e', 'i', 'o', 'u'},
		'e': {'i', 'o', 'u'},
		'i': {'o', 'u'},
		'o': {'u'},
		'u': {},
	}

	for j < len(word) {
		ch := word[j]
		if _, exist := w[ch]; !exist {
			w[ch] = 1
		} else {
			w[ch] += 1
		}

		if w[ch] == 1 {
			k++
		}

		for {
			flag := false
			for _, ic := range f[ch] {
				if w[ic] > 0 {
					w[word[i]]--
					if w[word[i]] == 0 {
						k--
					}
					flag = true

					break
				}
			}

			if !flag {
				break
			}

			i++
		}

		if k == 5 {
			if j-i+1 > max {
				max = j - i + 1
			}
		}

		j++
	}

	return max
}
