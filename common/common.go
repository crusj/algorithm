package common

func CA(src []int, elem int) []int {
	dst := make([]int, len(src))
	copy(dst, src)

	return append(dst, elem)
}
