package sort

func InversePairs(data []int) int {
	moded := 1000000007
	pairs := 0

	var tf func(data []int) []int
	tf = func(data []int) []int {
		if len(data) == 0 {
			return nil
		}

		if len(data) == 1 {
			return data
		}

		left := tf(data[0 : len(data)/2])
		right := tf(data[len(data)/2:])

		ret := make([]int, 0, len(left)+len(right))
		i, j := 0, 0
		for {
			if len(left) == i {
				ret = append(ret, right[j:]...)
				break
			}
			if len(right) == j {
				ret = append(ret, left[i:]...)
				break
			}

			if left[i] > right[j] {
				pairs = pairs + len(left) - i
				ret = append(ret, right[j])
				j++
			} else {
				ret = append(ret, left[i])
				i++
			}
		}

		return ret
	}
	tf(data)

	return pairs % moded
}
