package sort

// 幷归排序，时间复杂度nlogn
func mergeSort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return []int{arr[0]}
	}

	pivot := len(arr) / 2

	left := mergeSort(arr[0:pivot]) // logn
	right := mergeSort(arr[pivot:]) // logn

	merged := make([]int, 0, len(arr))
	lp, rp := 0, 0
	for { // n
		if lp == len(left) {
			merged = append(merged, right[rp:]...)
			break
		}

		if rp == len(right) {
			merged = append(merged, left[lp:]...)
			break
		}

		if left[lp] < right[rp] {
			merged = append(merged, left[lp])
			lp++
			continue
		}

		merged = append(merged, right[rp])
		rp++
	}

	return merged
}
