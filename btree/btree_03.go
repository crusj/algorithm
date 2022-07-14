package btree

// leetcode94
// 分治法
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

// leetcode94
// v2
func inorderTraversal_v2(root *TreeNode) []int {
	ret := make([]int, 0)

	var fn func(root *TreeNode)
	fn = func(root *TreeNode) {
		if root == nil {
			return
		}

		fn(root.Left)
		ret = append(ret, root.Val)
		fn(root.Right)
	}

	fn(root)

	return ret
}

// leetcode94
// v3
// 与v2的区别是把要变化的结果值作为参数传给内部函数
func inorderTraversal_v3(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var fn func(root *TreeNode, ret []int)
	fn = func(root *TreeNode, ret []int) {
		if root == nil {
			return
		}

		fn(root.Left, ret)
		ret = append(ret, root.Val)
		fn(root.Right, ret)
	}

	fn(root, res)

	return res
}

// leetcode100
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
	} else {
		return false
	}

	// 右子树相同与左子树相同
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}