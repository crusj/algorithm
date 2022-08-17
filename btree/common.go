package btree

// 构建二叉树
type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Data  int64
}

// 二叉树接口
type BinaryTreeBuilder interface {
	Insert(int64)
	GetRoot() *BinaryNode
}

// 创建二叉树
func NewBinaryTree(nodes []int64, i BinaryTreeBuilder) *BinaryNode {
	for _, v := range nodes {
		i.Insert(v)
	}

	return i.GetRoot()
}

// 二分查找树
type BinarySearchTree struct {
	root *BinaryNode
}

func (b *BinarySearchTree) Insert(data int64) {
	if b.root == nil {
		b.root = &BinaryNode{
			Data: data,
		}

		return
	}

	var tf func(node *BinaryNode, data int64)
	tf = func(node *BinaryNode, data int64) {
		if data <= node.Data {
			if node.Left != nil {
				tf(node.Left, data)
				return
			}

			node.Left = &BinaryNode{
				Data: data,
			}

			return
		}

		if node.Right != nil {
			tf(node.Right, data)
			return
		}

		node.Right = &BinaryNode{
			Data: data,
		}
	}

	tf(b.root, data)
}

func (b *BinarySearchTree) GetRoot() *BinaryNode {
	return b.root
}
