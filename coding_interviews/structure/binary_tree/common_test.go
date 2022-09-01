package binarytree

import (
	"log"
	"reflect"
	"testing"
)

func TestNewBinarySearchTreeByPre(t *testing.T) {
	// 二分查找树
	root := NewBinaryTree([]int64{10, 5, 20, 3, 6, 19, 30}, &BinarySearchTree{})
	if !reflect.DeepEqual(preOrderTraverse(root), []int64{10, 5, 3, 6, 20, 19, 30}) {
		log.Fatalf("Binary Search Tree failed: %+v", preOrderTraverse(root))
	}
}

func TestNewBinaryTreeByPre(t *testing.T) {
	// 普通树
	root := NewBinaryTree([]int64{1, 2, 3, -1, 4}, &BinaryTree{})
	if !reflect.DeepEqual(preOrderTraverse(root), []int64{1, 2, 4, 3}) {
		log.Fatalf("Binary Search Tree failed: %+v", preOrderTraverse(root))
	}
}

func TestNewBinaryTreeByPre2(t *testing.T) {
	// 普通树
	root := NewBinaryTree([]int64{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}, &BinaryTree{})
	if !reflect.DeepEqual(preOrderTraverse(root), []int64{5, 4, 11, 7, 2, 8, 13, 4, 5, 1}) {
		log.Fatalf("Binary Search Tree failed: %+v", preOrderTraverse(root))
	}
}
