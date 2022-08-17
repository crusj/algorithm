package linklist

import . "github.com/crusj/algorithm/common"

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}

	var ret *ListNode
	if pHead1.Val > pHead2.Val {
		ret = pHead2
		pHead2 = pHead2.Next
	} else {
		ret = pHead1
		pHead1 = pHead1.Next
	}

	tmp := ret

	for pHead2 != nil || pHead1 != nil {
		if pHead1 == nil {
			tmp.Next = pHead2
			break
		}

		if pHead2 == nil {
			tmp.Next = pHead1
			break
		}

		if pHead1.Val <= pHead2.Val {
			tmp.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			tmp.Next = pHead2
			pHead2 = pHead2.Next
		}

		tmp = tmp.Next
	}

	return ret
}
