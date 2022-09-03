package binarytree

import "testing"

func TestIsBalanced_Solution(t *testing.T) {
	type args struct {
		pRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				pRoot: nil,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				pRoot: NewBinaryTree([]int64{1, 2, 3, 4, 5, 6, 7}, &BinaryTree{}),
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				pRoot: NewBinaryTree([]int64{1, 2, 3, 4, 5, -1, 6, -1, -1, 7}, &BinaryTree{}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalanced_Solution(tt.args.pRoot); got != tt.want {
				t.Errorf("IsBalanced_Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
