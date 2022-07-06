package backtrack

// makesquare function    leetcode473
func makesquare(matchsticks []int) bool {
	result := false

	sum := 0
	for _, item := range matchsticks {
		sum += item
	}

	if sum%4 != 0 {
		return false
	}

	sideLen := sum / 4

	isValid := func(path []int) bool {
		sum := 0
		times := 0
		for _, v := range path {
			if sum < sideLen {
				sum = sum + matchsticks[v]

				if sum > sideLen {
					return false
				}

				if sum == sideLen {
					sum = 0
					times++
				}
			}
		}

		return times == 4
	}

	var fn func(path []int, pathMap map[int]struct{})
	fn = func(path []int, pathMap map[int]struct{}) {
		if result == true {
			return
		}

		if len(path) == len(matchsticks) {
			if isValid(path) {
				result = true
			}

			return
		}

		newPath := make([]int, len(path))
		copy(newPath, path)
		newPathMap := make(map[int]struct{}, len(path))
		for k, v := range pathMap {
			newPathMap[k] = v
		}

		for i := 0; i < len(matchsticks); i++ {
			if _, exists := newPathMap[i]; exists {
				continue
			}

			newPath = append(newPath, i)
			newPathMap[i] = struct{}{}

			fn(newPath, newPathMap)

			newPath = newPath[0 : len(newPath)-1] // back
			delete(newPathMap, i)                 // back
		}
	}

	for i := 0; i < len(matchsticks); i++ {
		fn([]int{i}, map[int]struct{}{
			i: {},
		})
	}

	return result
}
