package simulation

var (
	bytes []byte
	mb    = map[byte]int{}
)

func Insert(ch byte) {
	bytes = append(bytes, ch)
	if _, exists := mb[ch]; exists {
		mb[ch]++
	} else {
		mb[ch] = 1
	}
}

func FirstAppearingOnce() byte {
	for _, v := range bytes {
		if mb[v] == 1 {
			return v
		}
	}

	return '#'
}
