package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// leetcode 543
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDiameter := 0

	var fn func(*TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}

		maxDiameter = max(maxDepth(node.Left)+maxDepth(node.Right), maxDiameter)

		fn(node.Left)
		fn(node.Right)
	}

	fn(root)

	return maxDiameter
}

// 后序
func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftMax := maxDepth(node.Left)
	rightMax := maxDepth(node.Right)

	return 1 + max(leftMax, rightMax)
}
