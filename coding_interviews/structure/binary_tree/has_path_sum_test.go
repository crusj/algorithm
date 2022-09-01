package binarytree

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				root: nil,
				sum:  0,
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				root: NewBinaryTree([]int64{5, 4, 8, 1, 11, -1, 9, -1, -1, 2, 7}, &BinaryTree{}),
				sum:  22,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				root: NewBinaryTree([]int64{1, 2}, &BinaryTree{}),
				sum:  1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
