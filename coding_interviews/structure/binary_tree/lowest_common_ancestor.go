package binarytree

import "math"

func findVal(node *TreeNode, o1 int, o2 int) int {
	if node == nil {
		return int(math.MaxInt32)
	}

	if int(node.Val) == o1 {
		return o1
	}

	if int(node.Val) == o2 {
		return o2
	}

	t1 := findVal(node.Left, o1, o2)
	t2 := findVal(node.Right, o1, o2)

	if t1 != int(math.MaxInt32) && t2 != int(math.MaxInt32) {
		return int(node.Val)
	}

	if t1 != int(math.MaxInt32) && t2 == int(math.MaxInt32) {
		return t1
	}
	if t2 != int(math.MaxInt32) && t1 == int(math.MaxInt32) {
		return t2
	}

	return int(math.MaxInt32)
}

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here
	t := findVal(root, o1, o2)
	if t != int(math.MaxInt32) {
		return t
	}

	return 0
}

func bLowestCommonAncestor(root *TreeNode, p int, q int) int {
	// write code here
	if int(root.Val) < p && int(root.Val) < q {
		bLowestCommonAncestor(root.Right, p, q)
	}

	if int(root.Val) > p && int(root.Val) > q {
		bLowestCommonAncestor(root.Left, p, q)
	}

	return int(root.Val)
}
