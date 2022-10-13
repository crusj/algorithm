package common

import (
	"fmt"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	tmp := NewBinaryTree([]int64{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}, &BinaryTree{})
	ret := make([]int, 0)
	ft(tmp, &ret)
	fmt.Printf("%+v", ret)
}

func ft(node *TreeNode, ret *[]int) {
	*ret = append(*ret, int(node.Val))
	if node.Left != nil {
		ft(node.Left, ret)
	}
	if node.Right != nil {
		ft(node.Right, ret)
	}
}

func mt(node *TreeNode, ret *[]int) {
	if node.Left != nil {
		ft(node.Left, ret)
	}
	*ret = append(*ret, int(node.Val))
	if node.Right != nil {
		ft(node.Right, ret)
	}
}

func lt(node *TreeNode, ret *[]int) {
	if node.Left != nil {
		ft(node.Left, ret)
	}
	if node.Right != nil {
		ft(node.Right, ret)
	}
	*ret = append(*ret, int(node.Val))
}
