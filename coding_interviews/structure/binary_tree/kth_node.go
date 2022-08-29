package binarytree

func KthNode(proot *TreeNode, k int) int {
	// write code here
	if proot == nil || k <= 0 {
		return -1
	}

	ret := make([]int64, 0)

	var tf func(proot *TreeNode)
	tf = func(proot *TreeNode) {
		if proot == nil || len(ret) >= k {
			return
		}

		tf(proot.Left)
		ret = append(ret, proot.Val)
		tf(proot.Right)
	}

	tf(proot)

	if len(ret) < k {
		return -1
	}

	return int(ret[k-1])
}
