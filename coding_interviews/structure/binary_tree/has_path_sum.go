package binarytree

func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	if root == nil {
		return false
	}

	ret := false

	var tf func(node *TreeNode, total int)
	tf = func(node *TreeNode, total int) {
		if node.Left == nil && node.Right == nil {
			ret = total == sum

			return
		}

		if node.Left != nil {
			tf(node.Left, total+int(node.Left.Val))
		}

		if ret {
			return
		}

		if node.Right != nil {
			tf(node.Right, total+int(node.Right.Val))
		}
	}

	tf(root, int(root.Val))

	return ret
}
