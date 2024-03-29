package review

import (
	"math"
	"reflect"
	"strconv"
	"strings"

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

func Min2() int { // write code here
	return minStack[len(minStack)-1]
}

func ReverseSentence(str string) string {
	ret := ""
	var tf func(string, int)
	tf = func(s string, i int) {
		tmp := []byte{}
		for i < len(str) && str[i] != ' ' {
			tmp = append(tmp, str[i])
			i++
		}

		if i < len(str) {
			tf(s, i+1)
		}

		ret += " " + string(tmp)
	}

	tf(str, 0)
	return strings.TrimLeft(ret, " ")
}

// [1, 2, 3, 4, 5] [4, 5, 3, 2, 1]
func IsPopOrder(pushV []int, popV []int) bool {
	is := false
	var tf func(index int, stack []int, path []int)
	tf = func(index int, stack, path []int) {
		if is {
			return
		}

		if len(path) == len(popV) {
			if reflect.DeepEqual(path, popV) {
				is = true
			}
			return
		}

		// push
		if index < len(pushV) {
			stack = append(stack, pushV[index])
			tf(index+1, stack, path)
			stack = stack[0 : len(stack)-1]
		}

		// pop
		if len(stack) > 0 {
			if popV[len(path)] != stack[len(stack)-1] {
				return
			}
			tf(index, stack[0:len(stack)-1], append(path, stack[len(stack)-1]))
		}
	}

	tf(0, []int{}, []int{})
	return is
}

func GetNumberOfK(data []int, k int) int {
	// write code here
	times := 0
	var tf func(a, b int)
	tf = func(a, b int) {
		if a > b {
			return
		}

		pivot := (a + b) / 2
		if data[pivot] == k {
			for a <= pivot {
				if data[a] == k {
					times++
				}
				a++
			}

			for b > pivot {
				if data[b] == k {
					times++
				}
				b--
			}

			return
		}

		if data[pivot] > k {
			tf(a, pivot-1)
		} else {
			tf(pivot+1, b)
		}
	}

	tf(0, len(data)-1)
	return times
}

func minNumberInRotateArray(rotateArray []int) int {
	var tf func(b, e int) int
	tf = func(b, e int) int {
		if b > e {
			return math.MinInt64
		}

		pivot := (b + e) / 2
		if pivot-1 >= b && rotateArray[pivot-1] > rotateArray[pivot] {
			return rotateArray[pivot]
		}

		if pivot+1 <= e && rotateArray[pivot] > rotateArray[pivot+1] {
			return rotateArray[pivot+1]
		}
		if a := tf(b, pivot-1); a != math.MinInt64 {
			return a
		}

		return tf(pivot+1, e)
	}

	if v := tf(0, len(rotateArray)-1); v == math.MinInt64 {
		return rotateArray[0]
	} else {
		return v
	}
}

//
// [0    -   10)     10
// [10   -   100)    90
// [100  -   1000)   900
// [1000 -   10000)  9000

// 找出范围
// 找出是哪个数
// 找出第几位
func findNthDigit(n int) int {
	// write code here

	a, b := 0, 10
	i := 1
	for n > (b-a)*i {
		n -= (b - a) * i
		i++
		a, b = b, b*10
	}

	return int(strconv.Itoa(a + n/i)[a%i] - '0')
}

func Find(target int, array [][]int) bool {
	if array == nil {
		return false
	}

	i, j := len(array)-1, 0
	for i >= 0 && j < len(array[0]) {
		if target == array[i][j] {
			return true
		}
		if target < array[i][j] {
			i--
		} else {
			j++
		}
	}

	return false
}

func Permutation(str string) []string {
	ret := make([]string, 0)
	if str == "" {
		return nil
	}

	var tf func(slices []byte, index int, path []byte)
	tf = func(slices []byte, index int, path []byte) {
		if len(path) == len(str) {
			ret = append(ret, string(path))
			return
		}

		choosed := make(map[byte]struct{})
		for i := 0; i < len(slices); i++ {
			if _, exists := choosed[slices[i]]; exists {
				continue
			}

			path = append(path, slices[i])
			choosed[slices[i]] = struct{}{}

			tmp := make([]byte, len(slices))
			copy(tmp, slices)
			tf(append(tmp[0:i], tmp[i+1:]...), i, path)

			// back
			path = path[0 : len(path)-1]
		}
	}

	tf([]byte(str), -1, nil)

	return ret
}

func Permutation2(str string) []string {
	if str == "" {
		return nil
	}
	ret := make([]string, 0)
	var tf func(selected map[int]struct{}, path *[]byte)
	tf = func(selected map[int]struct{}, path *[]byte) {
		if len(*path) == len(str) {
			ret = append(ret, string(*path))
			return
		}

		choosed := make(map[byte]struct{})
		for i := 0; i < len(str); i++ {
			if _, exists := selected[i]; exists {
				continue
			}
			if _, exists := choosed[str[i]]; exists {
				continue
			}

			*path = append(*path, str[i])
			selected[i] = struct{}{}
			choosed[str[i]] = struct{}{}
			tf(selected, path)
			// back
			delete(selected, i)
			*path = (*path)[0 : len(*path)-1]
		}
	}

	tf(make(map[int]struct{}), &[]byte{})
	return ret
}
