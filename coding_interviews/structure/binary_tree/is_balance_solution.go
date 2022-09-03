package binarytree

import "math"

// 树高
func getTreeH(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := 1 + getTreeH(node.Left)
	right := 1 + getTreeH(node.Right)
	if left > right {
		return left
	}

	return right
}

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}

	if math.Abs(float64(getTreeH(pRoot.Left)-getTreeH(pRoot.Right))) > 1 {
		return false
	}

	return IsBalanced_Solution(pRoot.Right) && IsBalanced_Solution(pRoot.Left)
}
