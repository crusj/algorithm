package simulation

func FindContinuousSequence(sum int) [][]int {
	// write code here
	ret := make([][]int, 0)
	var fn func(s int, value int, path []int)
	already := make(map[int]struct{}, 0)
	fn = func(s, value int, path []int) {
		if s == sum {
			if len(path) > 1 {
				ret = append(ret, path)
			}
			return
		}

		if value >= sum {
			return
		}

		if s+value <= sum {
			fn(s+value, value+1, append(path, value))
		}

		if _, exists := already[value]; !exists {
			already[value] = struct{}{}
			fn(value, value+1, []int{value})
		}
	}

	fn(1, 2, []int{1})

	return ret
}
