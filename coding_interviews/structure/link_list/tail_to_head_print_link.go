package linklist

import . "github.com/crusj/algorithm/common"

// 利用递归，类似树的后续遍历
func printListFromTailToHead(head *ListNode) []int {
	// write code here
	ret := make([]int, 0)
	var tf func(head *ListNode)
	tf = func(head *ListNode) {
		if head != nil {
			tf(head.Next)
			ret = append(ret, head.Val)
		}
	}

	tf(head)

	return ret
}
