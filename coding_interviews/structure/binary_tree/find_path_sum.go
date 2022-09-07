package binarytree

func FindPathSum(root *TreeNode, sum int) int {
	ret := 0
	// write code here
	if root == nil {
		return ret
	}

	var tf func(node *TreeNode, tmp int, used map[*TreeNode]struct{})
	tf = func(node *TreeNode, tmp int, used map[*TreeNode]struct{}) {
		if tmp == sum {
			ret++
		}

		if node.Left != nil {
			tf(node.Left, int(node.Left.Val)+tmp, used)
			if _, exists := used[node.Left]; !exists {
				used[node.Left] = struct{}{}
				tf(node.Left, int(node.Left.Val), used)
			}

		}

		if node.Right != nil {
			tf(node.Right, int(node.Right.Val)+tmp, used)
			if _, exists := used[node.Right]; !exists {
				used[node.Right] = struct{}{}
				tf(node.Right, int(node.Right.Val), used)
			}
		}
	}

	tf(root, int(root.Val), map[*TreeNode]struct{}{})
	return ret
}

var res int

func FindPathSum2(root *TreeNode, sum int) int {
	// write code here
	if root == nil {
		return 0
	}

	dfs(root, sum)

	dfs(root.Left, sum)
	dfs(root.Right, sum)

	return res
}

func dfs(node *TreeNode, sum int) {
	if node == nil {
		return
	}

	if sum == int(node.Val) {
		res++
	}

	dfs(node.Left, sum-int(node.Val))
	dfs(node.Right, sum-int(node.Val))
}
