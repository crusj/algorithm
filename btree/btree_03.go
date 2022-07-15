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

// leetcode101
// BFS
func isSymmetric(root *TreeNode) bool {
	nodes := make([]*TreeNode, 0)
	// 第一层
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		nextNodes := make([]*TreeNode, 0)
		for _, node := range nodes {
			if node != nil {
				nextNodes = append(nextNodes, node.Left, node.Right)
			}
		}

		flag := true
		for i, j := 0, len(nextNodes)-1; i <= j; i, j = i+1, j-1 {
			if nextNodes[i] == nil && nextNodes[j] == nil {
				continue
			}

			if nextNodes[i] == nil || nextNodes[j] == nil {
				flag = false
				break
			}

			if nextNodes[i].Val != nextNodes[j].Val {
				flag = false
				break
			}
		}

		if !flag {
			return false
		}

		nodes = nextNodes
	}

	return true
}

// leetcode101_v2
func isSymmetric_v2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var fn func(left *TreeNode, right *TreeNode) bool
	fn = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left != nil && right != nil {
			if left.Val != right.Val {
				return false
			}
		} else {
			return false
		}

		return fn(left.Right, right.Left) && fn(left.Left, right.Right)
	}

	return fn(root.Left, root.Right)
}

// leetcode 110
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var fn func(root *TreeNode) bool
	fn = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		leftMaxDepth := maxDepth(root.Left)
		rightMaxDepth := maxDepth(root.Right)

		if tmp := leftMaxDepth - rightMaxDepth; tmp > 1 || tmp < -1 {
			return false
		}

		return fn(root.Left) && fn(root.Right)
	}

	return fn(root)
}

// leetcode 112
// 前序遍历路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	var fn func(node *TreeNode, sum int) bool
	fn = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}

		if sum += node.Val; sum == targetSum && node.Left == nil && node.Right == nil {
			return true
		}

		return fn(node.Left, sum) || fn(node.Right, sum)
	}

	return fn(root, 0)
}

// 前序遍历判断是否存在某个节点
func findNode(root *TreeNode, val1, val2 int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val1 || root.Val == val2 {
		return root
	}

	leftFind := findNode(root.Left, val1, val2)
	if leftFind != nil {
		return leftFind
	}

	rightFind := findNode(root.Right, val1, val2)
	if rightFind.Right != nil {
		return rightFind
	}

	return nil
}

// leetcode235
// 二叉搜索树的两个节点的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 情况一
	if root == p || root == q {
		return root
	}

	// 情况2
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	return nil
}

// leetcode235
// 二叉搜索树的两个节点的最近公共祖先
// 利用搜索二叉树的性质
func lowestCommonAncestor_v2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 右子树去找
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor_v2(root.Right, p, q)
	}

	// 左子树去找
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor_v2(root.Left, p, q)
	}

    return root
}
