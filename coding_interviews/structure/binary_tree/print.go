package binarytree

func PrintLine(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}

	ret := make([][]int, 0)
	next := []*TreeNode{pRoot}
	for len(next) > 0 {
		tmp := make([]*TreeNode, 0)
		items := make([]int, 0)
		for _, item := range next {
			items = append(items, int(item.Val))
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
		}

		ret = append(ret, items)
		next = tmp
	}

	return ret
}
