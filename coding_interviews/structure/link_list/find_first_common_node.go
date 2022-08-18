package linklist

import . "github.com/crusj/algorithm/common"

func findFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	tmp := pHead1
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = pHead2

	fast := pHead1
	slow := pHead1

	for {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}

		if slow.Next != nil {
			slow = slow.Next
		} else {
			return nil
		}

		if fast.Val == slow.Val {
			break
		}
	}

	fast = pHead1
	for {
		fast = fast.Next
		slow = slow.Next

		if fast.Val == slow.Val {
			return fast
		}
	}
}
