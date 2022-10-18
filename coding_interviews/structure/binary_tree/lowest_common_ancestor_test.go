package binarytree

import "testing"

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		o1   int
		o2   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				root: NewBinaryTree([]int64{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, &BinaryTree{}),
				o1:   5,
				o2:   1,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				root: NewBinaryTree([]int64{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, &BinaryTree{}),
				o1:   2,
				o2:   7,
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				root: NewBinaryTree([]int64{1, 2, 3, 4, 5}, &BinaryTree{}),
				o1:   4,
				o2:   5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.o1, tt.args.o2); got != tt.want {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
