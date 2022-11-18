package sort

// max heap
func maxHeap(arr []int) []int {
	maxHeap := make([]int, len(arr)+1)
	for i := 1; i <= len(arr); i++ { // nlogn
		maxHeap[i] = arr[i-1]
		j := i
		for j != 1 && maxHeap[j/2] < maxHeap[j] { // logn
			maxHeap[j] = maxHeap[j/2]
			j /= 2
		}

		maxHeap[j] = arr[i-1]
	}

	return maxHeap[1:]
}
