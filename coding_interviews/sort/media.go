package sort

import "fmt"

var minHeapS, maxHeapS = make([]int, 0), make([]int, 0)

func Insert(num int) {
	if len(minHeapS) == 0 {
		minHeapS = append(minHeapS, num)
		CreateMinHeap(minHeapS)
	} else if minHeapS[0] < num {
		minHeapS = append(minHeapS, num)
		CreateMinHeap(minHeapS)
		if len(minHeapS) > len(maxHeapS)+1 {
			maxHeapS = append(maxHeapS, minHeapS[0])
			minHeapS = minHeapS[1:]
			CreateMaxHeap(maxHeapS)
			CreateMinHeap(minHeapS)
		}
	} else {
		maxHeapS = append(maxHeapS, num)
		CreateMaxHeap(maxHeapS)
		if len(maxHeapS) > len(minHeapS)+1 {
			minHeapS = append(minHeapS, maxHeapS[0])
			maxHeapS = maxHeapS[1:]
			CreateMinHeap(minHeapS)
			CreateMaxHeap(maxHeapS)
		}
	}
	fmt.Printf("min heap%v\n", minHeapS)
	fmt.Printf("max heap%v\n", maxHeapS)
}

func GetMedian() float64 {
	if len(minHeapS) == len(maxHeapS) {
		return float64(minHeapS[0]+maxHeapS[0]) / 2.0
	}

	if len(minHeapS) < len(maxHeapS) {
		return float64(maxHeapS[0])
	}

	return float64(minHeapS[0])
}

// AdjustMaxHeap function    调整最大堆
func AdjustMaxHeap(nums []int, index int) {
	leftChild := index * 2
	rightChild := index*2 + 1

	max := index
	if leftChild <= len(nums) && nums[leftChild-1] > nums[index-1] {
		max = leftChild
	}

	if rightChild <= len(nums) && nums[rightChild-1] > nums[index-1] {
		max = leftChild
	}

	if max != index {
		nums[index-1], nums[max-1] = nums[max-1], nums[index-1]
		AdjustMaxHeap(nums, max)
	}
}

// AdjustMinHeap function    调整最小堆
func AdjustMinHeap(nums []int, index int) {
	leftChild, rightChild := index*1, index*2+1
	min := index
	if leftChild <= len(nums) && nums[leftChild-1] < nums[index-1] {
		min = leftChild
	}

	if rightChild <= len(nums) && nums[rightChild-1] < nums[index-1] {
		min = rightChild
	}

	if min != index {
		nums[index-1], nums[min-1] = nums[min-1], nums[index-1]
		AdjustMinHeap(nums, min)
	}
}

// CreateMaxHeap function    创建最大堆
func CreateMaxHeap(nums []int) {
	last := len(nums) / 2
	for last >= 0 {
		max := last
		left, right := last*2+1, last*2+2
		if left < len(nums) && nums[left] > nums[max] {
			max = left
		}
		if right < len(nums) && nums[right] > nums[max] {
			max = right
		}

		if max != last {
			nums[last], nums[max] = nums[max], nums[last]
		}
		AdjustMaxHeap(nums, max+1)

		last--
	}
}

// CreateMinHeap function    创建最小堆
func CreateMinHeap(nums []int) {
	last := len(nums) / 2
	for last >= 0 {
		min := last
		left, right := last*2+1, last*2+2
		if left < len(nums) && nums[left] < nums[min] {
			min = left
		}
		if right < len(nums) && nums[right] < nums[min] {
			min = right
		}

		if min != last {
			nums[last], nums[min] = nums[min], nums[last]
		}
		AdjustMinHeap(nums, min+1)

		last--
	}
}
