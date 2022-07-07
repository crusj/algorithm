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

// leetcode 111
// 二叉树最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	nodes := []*TreeNode{root}

	for {
		if len(nodes) == 0 {
			return depth
		}

		nexts := make([]*TreeNode, 0)
		depth++

		isMinNode := false
		for _, next := range nodes {
			if next.Left == nil && next.Right == nil {
				isMinNode = true
				break
			}

			if next.Left != nil {
				nexts = append(nexts, next.Left)
			}

			if next.Right != nil {
				nexts = append(nexts, next.Right)
			}
		}

		if isMinNode {
			return depth
		}

		nodes = make([]*TreeNode, len(nexts))
		copy(nodes, nexts)
	}
}
