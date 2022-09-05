package binarytree

func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}

	var tf func(left, right *TreeNode) bool
	tf = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		return tf(left.Left, right.Right) && tf(left.Right, right.Left)
	}

	return tf(pRoot.Left, pRoot.Right)
}
