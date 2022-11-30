package simulation

func reOrderArrayTwo(array []int) []int {
	// write code here
	i, j := 0, len(array)-1

	for {
		for array[i]%2 != 0 && i < j {
			i++
		}

		for array[j]%2 != 1 && j > i {
			j--
		}

		// switch
		if i < j {
			array[i], array[j] = array[j], array[i]

			i++
			j--
			continue
		}

		break
	}

	return array
}
