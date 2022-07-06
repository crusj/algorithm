package btree

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

var tree *Node

func init() {
	tree = &Node{
		Left: &Node{
			Left: &Node{
				Value: 5,
			},
			Right: &Node{
				Left: &Node{
					Value: 6,
				},
				Right: &Node{
					Value: 7,
				},
				Value: 4,
			},
			Value: 2,
		},
		Right: &Node{
			Left: &Node{
				Value: 8,
			},
			Right: &Node{
				Value: 9,
			},
			Value: 3,
		},
		Value: 1,
	}
}

// preOrderTraverse function  
// 分治法
func preOrderTraverse(tree *Node) []int {
	res := make([]int, 0)
	if tree == nil {
		return res
	}

	res = append(res, tree.Value)
	res = append(res, preOrderTraverse(tree.Left)...)
	res = append(res, preOrderTraverse(tree.Right)...)

	return res
}

// preOrderTraversev2 function  
// 遍历法
func preOrderTraverseV2(tree *Node) []int {
	res := make([]int, 0)

	var fn func(*Node)
	fn = func(n *Node) {
		if n == nil {
			return
		}
		res = append(res, n.Value)

		fn(n.Left)
		fn(n.Right)
	}

	fn(tree)

	return res
}
