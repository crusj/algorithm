package binarytree

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	type args struct {
		pRootOfTree *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			args: args{
				pRootOfTree: nil,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				pRootOfTree: NewBinaryTree([]int64{}, &BinarySearchTree{}),
			},
			want: &TreeNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.pRootOfTree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
