package binarytree

func FindPathSum(root *TreeNode, sum int) int {
	ret := 0
	// write code here
	if root == nil {
		return ret
	}

	var tf func(node *TreeNode, tmp int, repeated map[*TreeNode]struct{})
	tf = func(node *TreeNode, tmp int, repeated map[*TreeNode]struct{}) {
		if tmp == sum {
			ret++
		}

		if node.Left != nil {
			tf(node.Left, int(node.Left.Val)+tmp, repeated)
			if _, exists := repeated[node.Left]; !exists {
				repeated[node.Left] = struct{}{}
				tf(node.Left, int(node.Left.Val), repeated)
			}

		}

		if node.Right != nil {
			tf(node.Right, int(node.Right.Val)+tmp, repeated)
			if _, exists := repeated[node.Right]; !exists {
				repeated[node.Right] = struct{}{}
				tf(node.Right, int(node.Right.Val), repeated)
			}
		}
	}

	tf(root, int(root.Val), map[*TreeNode]struct{}{})
	return ret
}
