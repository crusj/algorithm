package leetcode

func removeDigit(number string, digit byte) string {
	last := 0

	ns := []byte(number)
	i := 0
	for i < len(ns) {
		if ns[i] == digit+48 {
			last = i
			if i+1 < len(ns) && ns[i] < ns[i+1] {
				return string(append(ns[0:i], ns[i+1:]...))
			}
		}
		i++
	}

	return string(append(ns[0:last], ns[last+1:]...))
}
