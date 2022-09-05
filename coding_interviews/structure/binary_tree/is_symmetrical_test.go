package binarytree

import "testing"

func Test_isSymmetrical(t *testing.T) {
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
				pRoot: NewBinaryTree([]int64{1, 2, 2, 3, 4, 4, 3}, &BinaryTree{}),
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				pRoot: NewBinaryTree([]int64{8, 6, 9, 5, 7, 7, 5}, &BinaryTree{}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetrical(tt.args.pRoot); got != tt.want {
				t.Errorf("isSymmetrical() = %v, want %v", got, tt.want)
			}
		})
	}
}
