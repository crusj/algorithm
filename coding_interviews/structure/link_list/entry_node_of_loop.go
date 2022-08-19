package linklist

import . "github.com/crusj/algorithm/common"

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}

	// 判断是否有环
	fast, slow := pHead, pHead

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}

	}
	if fast == nil || slow == nil {
		return nil
	}

	// 找出环入口
	tmp := pHead
	for tmp != slow {
		tmp = tmp.Next
		slow = slow.Next
	}

	return tmp
}
