package linklist

import . "github.com/crusj/algorithm/common"

func deleteDuplication(pHead *ListNode) *ListNode {
	var fn func(pHead *ListNode) *ListNode
	fn = func(pHead *ListNode) *ListNode {
		if pHead == nil {
			return nil
		}

		if pHead.Next == nil {
			return pHead
		}
		next := pHead.Next

		if pHead.Val != next.Val {
			pHead.Next = fn(pHead.Next)
			return pHead
		}

		for next != nil && next.Val == pHead.Val {
			next = next.Next
		}

		return fn(next)
	}

	return fn(pHead)
}
