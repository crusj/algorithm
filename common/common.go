package common

func CA(src []int, elem int) []int {
	dst := make([]int, len(src))
	copy(dst, src)

	return append(dst, elem)
}

func Min(args ...int) int {
	if len(args) == 1 {
		return args[0]
	}

	a := args[0]
	b := Min(args[1:]...)
	if a > b {
		return b
	}

	return a
}

func Max(args ...int) int {
	if len(args) == 1 {
		return args[0]
	}

	tmp := Max(args[1:]...)
	if tmp > args[0] {
		return tmp
	}

	return args[0]
}
