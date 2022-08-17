package linklist

import . "github.com/crusj/algorithm/common"

// ReverseList function    思路和倒序打印单链表相同，都是利用类似树的后续遍历
func ReverseList(pHead *ListNode) *ListNode {
	var ret *ListNode
	var tf func(prev, pHead *ListNode)
	tf = func(prev, pHead *ListNode) {
		if pHead != nil {
			tf(pHead, pHead.Next)
			pHead.Next = prev
		} else {
			ret = prev
		}
	}

	tf(nil, pHead)

	return ret
}
