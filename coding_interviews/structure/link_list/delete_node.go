package linklist

import . "github.com/crusj/algorithm/common"

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	prev := head
	next := head.Next

	for next != nil {
		if next.Val == val {
			prev.Next = next.Next
			break
		}

		next, prev = next.Next, prev.Next
	}

	return head
}
