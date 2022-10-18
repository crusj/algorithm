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

func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}

	var p, ret *TreeNode

	var tf func(node *TreeNode)
	tf = func(node *TreeNode) {
		if node.Left != nil {
			tf(node.Left)
		}
		if p == nil {
			p = node
			ret = p
		} else {
			p.Right = node
			node.Left = p
			p = node
		}
		if node.Right != nil {
			tf(node.Right)
		}
	}
	tf(pRootOfTree)

	return ret
}

func Print2(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}

	// write code here
	ret := make([][]int, 0)

	var tf func(data []*TreeNode)
	tf = func(data []*TreeNode) {
		if len(data) == 0 {
			return
		}
		tmp := make([]*TreeNode, 0)

		val := make([]int, len(data))
		for i, item := range data {
			val[i] = int(item.Val)
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
		}

		ret = append(ret, val)
		tf(tmp)
	}

	tf([]*TreeNode{pRoot})
	return ret
}

func FindPath2(root *TreeNode, sum int) int {
	// write code here
	if root == nil {
		return 0
	}

	ret := 0

	var tf func(node *TreeNode, left int)
	tf = func(node *TreeNode, left int) {
		if left == int(node.Val) {
			ret++
		}

		if node.Left != nil {
			tf(node.Left, left-int(node.Val))
		}

		if node.Right != nil {
			tf(node.Right, left-int(node.Val))
		}
	}

	var frontTraverse func(root *TreeNode)
	frontTraverse = func(root *TreeNode) {
		tf(root, sum)
		if root.Left != nil {
			frontTraverse(root.Left)
		}
		if root.Right != nil {
			frontTraverse(root.Right)
		}
	}

	frontTraverse(root)

	return ret
}

func lowestCommonAncestor2(root *TreeNode, o1 int, o2 int) int {
	// write code here
	repeated := map[*TreeNode]int{}
	var frontTraverse func(root *TreeNode) int
	frontTraverse = func(root *TreeNode) int {
		if _, exists := repeated[root]; exists {
			return repeated[root]
		}

		if root == nil {
			return math.MinInt64
		}

		if int(root.Val) == o1 || int(root.Val) == o2 {
			repeated[root] = int(root.Val)
			return int(root.Val)
		}

		repeated[root] = frontTraverse(root.Left)
		if repeated[root] != math.MinInt64 {
			return repeated[root]
		}

		repeated[root] = frontTraverse(root.Right)
		if repeated[root] != math.MinInt64 {
			return repeated[root]
		}

		return math.MinInt64
	}

	var tf func(root *TreeNode) int
	tf = func(root *TreeNode) int {
		if int(root.Val) == o1 || int(root.Val) == o2 {
			return int(root.Val)
		}

		lv := frontTraverse(root.Left)
		rv := frontTraverse(root.Right)
		if lv != math.MinInt64 && rv != math.MinInt64 {
			return int(root.Val)
		}
		if lv != math.MinInt64 && rv == math.MinInt64 {
			return tf(root.Left)
		}

		if lv == math.MinInt64 && rv != math.MinInt64 {
			return tf(root.Right)
		}

		return math.MinInt64
	}

	return tf(root)
}

var (
	stack1 []int
	stack2 []int
)

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	tmp := stack1[len(stack1)-1]
	stack1 = stack1[1:]

	return tmp
}

var (
	stack3   []int
	minStack []int
)

func Push2(node int) {
	stack3 = append(stack3, node)

	if len(minStack) == 0 || node < minStack[len(minStack)-1] {
		minStack = append(minStack, node)
	} else {
		minStack = append(minStack, minStack[len(minStack)-1])
	}
}

func Pop2() {
	stack3 = stack3[0 : len(stack3)-1]
	minStack = minStack[0 : len(minStack)-1]
}

func Top2() int {
	return stack3[len(stack3)-1]
}

func Min2() int {
	// write code here
	return minStack[len(minStack)-1]
}

func ReverseSentence( str string ) string {
    // write code here
    var 
    func ReverseSentence( str string ) string {
            // write code here
        }
}
