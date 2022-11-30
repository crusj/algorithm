package simulation

func LeftRotateString(str string, n int) string {
	// write code here
	if str == "" {
		return ""
	}

	md := n
	if n > len(str) {
		md = n % len(str)
	}

	movedStr := make([]byte, len(str))
	for i, v := range []byte(str) {
		if i-md < 0 {
			movedStr[len(str)+i-md] = v
			continue
		}
		movedStr[i-md] = v
	}

	return string(movedStr)
}
