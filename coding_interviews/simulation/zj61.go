package simulation

// IsContinuous function    是否是顺子
func IsContinuous(numbers []int) bool {
	// write code here
	CreateMinHeap(numbers)
	l := len(numbers)
	for l > 1 {
		numbers[l-1], numbers[0] = numbers[0], numbers[l-1]
		CreateMinHeap(numbers[0 : l-1])
		l--
	}
	p := len(numbers) - 1
	i := 0
	for {
		if numbers[i] == 0 {
			return true
		}
		if i >= p {
			return true
		}

		if numbers[i] == numbers[i+1] {
			return false
		}

		if numbers[i]-1 == numbers[i+1] {
			i++
			continue
		}

		if numbers[p] == 0 {
			numbers[i]--
			p--
			continue
		}

		return false
	}
}

func CreateMinHeap(numbers []int) {
	i := len(numbers) / 2
	for {
		if i < 0 {
			return
		}

		AdjustMinHeap(numbers, i)
		i--
	}
}

// AdjustMinHeap function    调整堆
func AdjustMinHeap(numbers []int, index int) {
	i, j := index*2+1, index*2+2
	min := index

	if i < len(numbers) && numbers[i] < numbers[min] {
		min = i
	}

	if j < len(numbers) && numbers[j] < numbers[min] {
		min = j
	}

	if min != index {
		numbers[min], numbers[index] = numbers[index], numbers[min]
		AdjustMinHeap(numbers, min)
	}
}
