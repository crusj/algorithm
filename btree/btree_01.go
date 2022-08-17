package btree

// preOrderTraverse function  
// 分治法
func preOrderTraverse(tree *BinaryNode) []int64 {
	res := make([]int64, 0)
	if tree == nil {
		return res
	}

	res = append(res, tree.Data)
	res = append(res, preOrderTraverse(tree.Left)...)
	res = append(res, preOrderTraverse(tree.Right)...)

	return res
}

// preOrderTraversev2 function  
// 遍历法
func preOrderTraverseV2(tree *BinaryNode) []int64 {
	res := make([]int64, 0)

	var fn func(*BinaryNode)
	fn = func(n *BinaryNode) {
		if n == nil {
			return
		}
		res = append(res, n.Data)

		fn(n.Left)
		fn(n.Right)
	}

	fn(tree)

	return res
}
