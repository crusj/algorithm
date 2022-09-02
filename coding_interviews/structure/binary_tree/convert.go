package binarytree

func Convert(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	if pRootOfTree == nil {
		return nil
	}

	nodes := make([]*TreeNode, 0)
	var tf func(node *TreeNode)
	tf = func(node *TreeNode) {
		if node == nil {
			return
		}

		tf(node.Left)
		nodes = append(nodes, node)
		tf(node.Right)
	}

	tf(pRootOfTree)

	for i := 1; i <= len(nodes)-1; i++ {
		nodes[i-1].Right = nodes[i]
		nodes[i].Left = nodes[i-1]
	}

	return nodes[0]
}
