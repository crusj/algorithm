package sort

// 创建最大堆
func maxHeap(arr []int) []int {
	var heapify func(pivot int)
	heapify = func(pivot int) {
		max := pivot
		l := pivot*2 + 1
		r := pivot*2 + 2

		if l < len(arr) && arr[l] > arr[max] {
			max = l
		}

		if r < len(arr) && arr[r] > arr[max] {
			max = r
		}

		if max != pivot {
			arr[max], arr[pivot] = arr[pivot], arr[max]
			heapify(max)
		}
	}

	pivot := len(arr) / 2
	for pivot >= 0 {
		heapify(pivot)
		pivot--
	}

	return arr
}

// 堆排序
func heapSort(arr []int) []int {
	l := len(arr)
	for l > 0 {
		maxHeap(arr[0:l])
		arr[l-1], arr[0] = arr[0], arr[l-1]
		l--
	}

	return arr
}
