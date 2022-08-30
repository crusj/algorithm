package binarytree

import "testing"

func TestHasSubtree(t *testing.T) {
	type args struct {
		pRoot1 *TreeNode
		pRoot2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				pRoot1: nil,
				pRoot2: nil,
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				pRoot1: NewBinaryTree([]int64{8, 8, 7, 9, 2, -1, -1, -1, 4, 7}, &BinaryTree{}),
				pRoot2: NewBinaryTree([]int64{8, 9, 2}, &BinaryTree{}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasSubtree(tt.args.pRoot1, tt.args.pRoot2); got != tt.want {
				t.Errorf("HasSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
