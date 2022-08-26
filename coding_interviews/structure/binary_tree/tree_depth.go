package binarytree

func TreeDepth(pRoot *TreeNode) int {
	// write code here
	if pRoot == nil {
		return 0
	}

	// 版本一
	// max := 1
	// var tf func(node *TreeNode, depth int)
	// tf = func(node *TreeNode, depth int) {
	// 	if node == nil {
	// 		return
	// 	}
	//
	// 	if depth > max {
	// 		max = depth
	// 	}
	//
	// 	tf(node.Left, depth+1)
	// 	tf(node.Right, depth+1)
	// }
	//
	// tf(pRoot, 1)
	//
	// return max

	// 版本二
	// a, b := 1+TreeDepth(pRoot.Left), 1+TreeDepth(pRoot.Right)
	// if a > b {
	// 	return a
	// }
	//
	// return b

	// 版本三
	a, b := 1+TreeDepth(pRoot.Left), 1+TreeDepth(pRoot.Right)
	if a > b {
		return a + 1
	}

	return b + 1
}
