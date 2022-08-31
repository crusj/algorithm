package binarytree

func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return nil
	}

	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left

	Mirror(pRoot.Left)
	Mirror(pRoot.Right)

	return pRoot
}
