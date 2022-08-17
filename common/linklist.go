package common

type ListNode struct {
	Val  int
	Next *ListNode
}

// NewSingleList function    return single list
func NewSingleList(nodes []int) *ListNode {
	head := &ListNode{
		Val: nodes[0],
	}

	root := head

	for _, item := range nodes[1:] {
		head.Next = &ListNode{
			Val: item,
		}

		head = head.Next
	}

	return root
}
