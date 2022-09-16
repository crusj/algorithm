package search

func GetNumberOfK(data []int, k int) int {
	// write code here
	var tf func(data []int, i, j int) int
	tf = func(data []int, i, j int) int {
		if i > j {
			return 0
		}

		if i == j {
			if data[i] == k {
				return 1
			}

			return 0
		}

		pivot := (j + i) / 2
		if data[pivot] < k {
			return tf(data, pivot+1, j)
		}
		if data[pivot] > k {
			return tf(data, i, pivot-1)
		}

		begin := i
		for data[begin] < data[pivot] && begin <= pivot {
			begin++
		}

		end := j
		for data[end] > data[pivot] && end >= pivot {
			end--
		}

		if begin > pivot && end < pivot {
			return 0
		}

		return end - begin + 1
	}

	return tf(data, 0, len(data)-1)
}
