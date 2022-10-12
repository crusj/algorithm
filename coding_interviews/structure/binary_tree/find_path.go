package binarytree

// 查找路径
func FindPath(root *TreeNode, expectNumber int) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)

	var tf func(root *TreeNode, path []int, sum int)
	tf = func(root *TreeNode, path []int, sum int) {
		if root.Left == nil && root.Right == nil {
			if sum == expectNumber {
				ret = append(ret, path)
			}

			return
		}

		newPath := make([]int, len(path))
		copy(newPath, path)

		if root.Left != nil {
			tf(root.Left, append(newPath, int(root.Left.Val)), sum+int(root.Left.Val))
		}

		if root.Right != nil {
			tf(root.Right, append(newPath, int(root.Right.Val)), sum+int(root.Right.Val))
		}
	}

	tf(root, []int{int(root.Val)}, int(root.Val))

	return ret
}
