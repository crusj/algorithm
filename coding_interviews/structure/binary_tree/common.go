package binarytree

// 构建二叉树
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int64
}

// 二叉树接口
type BinaryTreeBuilder interface {
	Create([]int64)
	GetRoot() *TreeNode
}

// 普通二叉树
func NewBinaryTree(nodes []int64, i BinaryTreeBuilder) *TreeNode {
	i.Create(nodes)
	return i.GetRoot()
}

type BinaryTree struct {
	root *TreeNode
}

func (b *BinaryTree) Create(data []int64) {
	if len(data) == 0 {
		return
	}

	b.root = &TreeNode{
		Val: data[0],
	}

	var tf func(*TreeNode, int, []int64)

	tf = func(node *TreeNode, pos int, data []int64) {
		// 左子树
		if pos*2 <= len(data) && data[pos*2-1] != -1 {
			node.Left = &TreeNode{
				Val: data[pos*2-1],
			}
			tf(node.Left, pos*2, data)
		}

		// 右子树
		if pos*2+1 <= len(data) && data[pos*2] != -1 {
			node.Right = &TreeNode{
				Val: data[pos*2],
			}

			tf(node.Right, pos*2+1, data)
		}
	}

	tf(b.root, 1, data)
}

func (b *BinaryTree) GetRoot() *TreeNode {
	return b.root
}

// 二分查找树
type BinarySearchTree struct {
	root *TreeNode
}

func (b *BinarySearchTree) Create(data []int64) {
	if len(data) == 0 {
		return
	}

	b.root = &TreeNode{
		Val: data[0],
	}

	var tf func(node *TreeNode, data int64)
	tf = func(node *TreeNode, data int64) {
		if data <= node.Val {
			if node.Left != nil {
				tf(node.Left, data)
				return
			}

			node.Left = &TreeNode{
				Val: data,
			}

			return
		}

		if node.Right != nil {
			tf(node.Right, data)
			return
		}

		node.Right = &TreeNode{
			Val: data,
		}
	}

	for _, v := range data[1:] {
		tf(b.root, v)
	}
}

func (b *BinarySearchTree) GetRoot() *TreeNode {
	return b.root
}

func preOrderTraverse(tree *TreeNode) []int64 {
	res := make([]int64, 0)
	if tree == nil {
		return res
	}

	res = append(res, tree.Val)
	res = append(res, preOrderTraverse(tree.Left)...)
	res = append(res, preOrderTraverse(tree.Right)...)

	return res
}
