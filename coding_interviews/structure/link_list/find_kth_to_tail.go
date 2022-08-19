package linklist

import . "github.com/crusj/algorithm/common"

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	if pHead == nil || k == 0 {
		return nil
	}

	fast, slow := pHead, pHead
	for k > 1 && fast != nil {
		fast = fast.Next
		k--
	}

	if fast == nil {
		return nil
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
