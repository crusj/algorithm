package binarytree

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot2 == nil {
		return false
	}

	var tf func(node1 *TreeNode, node2 *TreeNode) bool
	tf = func(node1, node2 *TreeNode) bool {
		if node2 == nil {
			return true
		}

		if node1 == nil {
			return false
		}

		return (node1.Val == node2.Val && tf(node1.Left, node2.Left) && tf(node1.Right, node2.Right)) || tf(node1.Left, pRoot2) || tf(node1.Right, pRoot2)
	}

	return tf(pRoot1, pRoot2)
}
