package binarytree

func PrintFromTopToBottom(root *TreeNode) []int {
	// write code here
	if root == nil {
		return nil
	}

	nodes := []*TreeNode{root}
	ret := make([]int, 0)

	for len(nodes) > 0 {
		tmp := make([]*TreeNode, 0)
		for _, item := range nodes {
			ret = append(ret, int(item.Val))
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}

			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
		}

		nodes = tmp
	}

	return ret
}
