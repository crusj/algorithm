package quenuestack

var (
	stack1 []int
	stack2 []int
)

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		length := len(stack1)
		for length > 0 {
			stack2 = append(stack2, stack1[length-1])
			stack1 = stack1[0 : length-1]
			length = len(stack1)
		}
	}

	t := stack2[len(stack2)-1]
	stack2 = stack2[0 : len(stack2)-1]

	return t
}
