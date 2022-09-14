package minstack

var (
	stack    []int
	minStack []int
)

func Push(node int) {
	// write code here
	stack = append(stack, node)

	if len(minStack) == 0 || node < minStack[len(minStack)-1] {
		minStack = append(minStack, node)
	} else {
		minStack = append(minStack, minStack[len(minStack)-1])
	}
}

func Pop() {
	stack = stack[0 : len(stack)-1]
	minStack = minStack[0 : len(minStack)-1]
}

func Top() int {
	return stack[len(stack)-1]
}

func Min() int {
	// write code here
	return minStack[len(minStack)-1]
}
