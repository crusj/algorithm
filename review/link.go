package review

import . "github.com/crusj/algorithm/common"

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	if head == nil {
		return nil
	}

	return append(printListFromTailToHead(head.Next), head.Val)
}

func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}

	// write code here
	var newHead *ListNode
	var tf func(pHead *ListNode)
	tf = func(pHead *ListNode) {
		if pHead.Next == nil {
			newHead = pHead
			return
		}

		tf(pHead.Next)
		pHead.Next.Next = pHead
	}

	tf(pHead)
	pHead.Next = nil

	return newHead
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead2 == nil && pHead1 == nil {
		return nil
	}
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}

	var first *ListNode

	if pHead1.Val < pHead2.Val {
		first = pHead1
		pHead1 = pHead1.Next
	} else {
		first = pHead2
		pHead2 = pHead2.Next
	}

	tmp := first
	for {
		if pHead1 == nil {
			tmp.Next = pHead2
			break
		}

		if pHead2 == nil {
			tmp.Next = pHead1
			break
		}

		if pHead1.Val < pHead2.Val {
			tmp.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			tmp.Next = pHead2
			pHead2 = pHead2.Next
		}

		tmp = tmp.Next
	}

	return first
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	step1s := 0
	step2s := 0
	tmp1, tmp2 := pHead1, pHead2
	for pHead2 != nil || pHead1 != nil {
		if pHead2 != nil {
			step2s++
			pHead2 = pHead2.Next
		}

		if pHead1 != nil {
			step1s++
			pHead1 = pHead1.Next
		}
	}

	if step1s > step2s {
		for i := 0; i < step1s-step2s; i++ {
			tmp1 = tmp1.Next
		}
	} else {
		for i := 0; i < step2s-step1s; i++ {
			tmp2 = tmp2.Next
		}
	}

	for tmp1.Val != tmp2.Val {
		if tmp1.Next == nil || tmp2.Next == nil {
			return nil
		}

		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}

	return tmp1
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}

	fast, slow := pHead, pHead
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			break
		}
	}

	fast = pHead

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	if pHead == nil || k == 0 {
		return nil
	}

	// write code here
	fast, slow := pHead, pHead
	for k > 0 {
		if fast == nil {
			return nil
		}

		fast = fast.Next
		k--
	}

	for {
		if fast == nil {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	prev, p := head, head.Next
	for p != nil {
		if p.Val == val {
			prev.Next = p.Next
			return head
		}
		prev = prev.Next
		p = p.Next
	}

	return nil
}

func deleteDuplication(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}

	if pHead.Next == nil {
		return pHead
	}

	p := pHead.Next
	if p.Val != pHead.Val {
		pHead.Next = deleteDuplication(p)
		return pHead
	}

	for p != nil && p.Val == pHead.Val {
		p = p.Next
	}

	if p == nil {
		return nil
	}

	return deleteDuplication(p)
}
