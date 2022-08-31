package binarytree

import (
	"reflect"
	"testing"
)

func TestMirror(t *testing.T) {
	type args struct {
		pRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			args: args{
				pRoot: nil,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				pRoot: NewBinaryTree([]int64{8, 6, 10}, &BinaryTree{}),
			},
			want: NewBinaryTree([]int64{8, 10, 6}, &BinaryTree{}),
		},
		{
			name: "test3",
			args: args{
				pRoot: NewBinaryTree([]int64{8, 6, 10, 5, 7, 9, 11}, &BinaryTree{}),
			},
			want: NewBinaryTree([]int64{8, 10, 6, 11, 9, 7, 5}, &BinaryTree{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mirror(tt.args.pRoot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mirror() = %v, want %v", got, tt.want)
			}
		})
	}
}
