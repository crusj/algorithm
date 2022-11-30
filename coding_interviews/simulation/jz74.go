package simulation

func FindContinuousSequence(sum int) [][]int {
	if sum == 1 {
		return nil
	}

	ret := make([][]int, 0)

	for i := 1; i <= sum; i++ {
		path := []int{i}
		s := i
		for j := i + 1; j <= sum; j++ {
			s = s + j
			if s < sum {
				path = append(path, j)
				continue
			} else if s == sum {
				ret = append(ret, append(path, j))
				break
			}

			break
		}
	}

	return ret
}
