package binarytree

// 层序遍历可以用递归或者用for循环实现
// Z字形遍历，要么遍历形式改变(从左到右遍历节点，还是从右向左遍历节点)，要么改变节点位置需要从左向右遍历的时候节点顺序不变，需要从右向左遍历时候节点顺序逆转,卡在是否要将节点进行顺序调节
func Print(pRoot *TreeNode) [][]int {
	// write code here
	if pRoot == nil {
		return nil
	}

	ret := make([][]int, 0)
	var tf func(nodes []*TreeNode, left bool)
	tf = func(nodes []*TreeNode, left bool) {
		if len(nodes) == 0 {
			return
		}

		tmp := make([]int, 0, len(nodes))
		if left {
			for _, item := range nodes {
				tmp = append(tmp, int(item.Val))
			}
		} else {
			for i := len(nodes) - 1; i >= 0; i-- {
				tmp = append(tmp, int(nodes[i].Val))
			}
		}

		childs := make([]*TreeNode, 0)
		for _, node := range nodes {
			if node.Left != nil {
				childs = append(childs, node.Left)
			}
			if node.Right != nil {
				childs = append(childs, node.Right)
			}
		}

		ret = append(ret, tmp)
		tf(childs, !left)
	}

	tf([]*TreeNode{pRoot}, true)

	return ret
}

func PrintOne(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	ret := make([][]int, 0)

	for len(stack) > 0 {
		tmp := make([]int, 0)
		tStack := make([]*TreeNode, 0)

		for _, node := range stack {
			if node.Left != nil {
				tmp = append(tmp, int(node.Left.Val))
				tStack = append(tStack, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, int(node.Right.Val))
				tStack = append(tStack, node.Right)
			}
		}

		stack = tStack
		ret = append(ret, tmp)
	}

	return ret
}
