package btree

import (
	"log"
	"reflect"
	"testing"
)

func TestNewBinaryTreeByPre(t *testing.T) {
	// 二分查找树
	root := NewBinaryTree([]int64{10, 5, 20, 3, 6, 19, 30}, &BinarySearchTree{})
	if !reflect.DeepEqual(preOrderTraverse(root), []int64{10, 5, 3, 6, 20, 19, 30}) {
		log.Fatalf("Binary Search Tree failed: %+v", preOrderTraverse(root))
	}

	// 前序遍历二叉树
}
