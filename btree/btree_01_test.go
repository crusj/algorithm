package btree

import (
	"reflect"
	"testing"
)

func Test_preOrderTraverse(t *testing.T) {
	type args struct {
		tree *BinaryNode
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test1",
			args: args{
				tree: NewBinaryTree([]int64{1, 2, 3, 4, 5}, &BinarySearchTree{}),
			},
			want: []int64{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preOrderTraverse(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preOrderTraverse() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_preOrderTraverseV2(t *testing.T) {
	type args struct {
		tree *BinaryNode
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test1",
			args: args{
				tree: NewBinaryTree([]int64{1, 2, 3, 4, 5}, &BinarySearchTree{}),
			},
			want: []int64{
				1, 2, 3, 4, 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preOrderTraverseV2(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preOrderTraverseV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
