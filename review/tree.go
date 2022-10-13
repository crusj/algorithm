package review

import (
	"math"

	. "github.com/crusj/algorithm/common"
)

func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}

	leftDepth := TreeDepth(pRoot.Left)
	rightDepth := TreeDepth(pRoot.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return nil
	}
	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)

	return pRoot
}

func PrintFromTopToBottom(root *TreeNode) []int {
	// write code here
	if root == nil {
		return nil
	}

	ret := []int{}
	basket := []*TreeNode{root}

	for len(basket) > 0 {
		tmp := []*TreeNode{}
		for _, item := range basket {
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
			ret = append(ret, int(item.Val))
		}

		basket = tmp
	}

	return ret
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && sum == int(root.Val) {
		return true
	}

	if root.Left != nil {
		if hasPathSum(root.Left, sum-int(root.Val)) {
			return true
		}
	}

	if root.Right != nil {
		if hasPathSum(root.Right, sum-int(root.Val)) {
			return true
		}
	}

	return false
}

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}

	var tf func(root *TreeNode) int
	tf = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		ld := tf(root.Left) + 1
		rd := tf(root.Right) + 1
		if ld > rd {
			return ld
		}
		return rd
	}

	if math.Abs(float64(tf(pRoot.Left))-float64(tf(pRoot.Right))) > 1 {
		return false
	}

	return IsBalanced_Solution(pRoot.Left) && IsBalanced_Solution(pRoot.Right)
}

func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}

	var tf func(node1, node2 *TreeNode) bool
	tf = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		return tf(node1.Left, node2.Right) && tf(node1.Right, node2.Left)
	}

	return tf(pRoot.Left, pRoot.Right)
}

func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	if root == nil {
		return math.MinInt64
	}

	if (p < int(root.Val) && q > int(root.Val)) || (q < int(root.Val) && p > int(root.Val)) {
		return int(root.Val)
	}

	if p == int(root.Val) || q == int(root.Val) {
		return int(root.Val)
	}

	if p < int(root.Val) && q < int(root.Val) {
		ret := lowestCommonAncestor(root.Left, p, q)
		if ret != math.MinInt64 {
			return ret
		}
	}

	return lowestCommonAncestor(root.Right, p, q)
}

// 层序遍历加后序遍历
func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}

	// write code here
	var tf func(data []int) []int
	tf = func(data []int) []int {
		if len(data) == 0 {
			return nil
		}

		return append(tf(data[1:]), data[0])
	}

	ret := make([][]int, 0)
	brakets := []*TreeNode{pRoot}
	d := true
	for len(brakets) > 0 {
		tmp := make([]*TreeNode, 0)

		vals := make([]int, 0)
		for _, v := range brakets {
			vals = append(vals, int(v.Val))
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		if d {
			ret = append(ret, vals)
		} else {
			ret = append(ret, tf(vals))
		}

		d = !d
		brakets = tmp
	}

	return ret
}

// 中序遍历
func KthNode(proot *TreeNode, k int) int {
	if proot == nil || k == 0 {
		return -1
	}

	vals := make([]int, 0)

	var tf func(proot *TreeNode)
	tf = func(proot *TreeNode) {
		if proot == nil {
			return
		}

		tf(proot.Left)
		vals = append(vals, int(proot.Val))
		tf(proot.Right)
	}

	tf(proot)
	if len(vals) < k {
		return -1
	}

	return vals[k-1]
}

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// write code here
	if pRoot2 == nil {
		return false
	}
	var is func(node1, node2 *TreeNode) bool
	is = func(node1, node2 *TreeNode) bool {
		if node1 != nil && node2 != nil {
			if node1.Val != node2.Val {
				return false
			}

			return is(node1.Left, node2.Left) && is(node1.Right, node2.Right)
		}

		if node1 == nil && node2 != nil {
			return false
		}

		return true
	}

	var ft func(node1 *TreeNode, node2 *TreeNode) bool
	ft = func(node1, node2 *TreeNode) bool {
		if node1 == nil {
			return false
		}

		if is(node1, node2) {
			return true
		}

		if ft(node1.Left, node2) {
			return true
		}

		return ft(node1.Right, node2)
	}

	return ft(pRoot1, pRoot2)
}

func FindPath(root *TreeNode, expectNumber int) [][]int {
	if root == nil {
		return nil
	}

	// write code here
	ret := make([][]int, 0)

	var tf func(root *TreeNode, expectNumber int, path []int)
	tf = func(root *TreeNode, expectNumber int, path []int) {
		// 叶子节点
		if root.Left == nil && root.Right == nil {
			if expectNumber == 0 {
				tmp := make([]int, len(path))
				copy(tmp, path)
				ret = append(ret, tmp)
			}

			return
		}

		if root.Left != nil {
			path := append(path, int(root.Left.Val))
			tf(root.Left, expectNumber-int(root.Left.Val), path)
			path = path[0 : len(path)-1] // back trace
		}

		if root.Right != nil {
			tf(root.Right, expectNumber-int(root.Right.Val), append(path, int(root.Right.Val)))
		}
	}

	tf(root, expectNumber-int(root.Val), []int{int(root.Val)})
	return ret
}
