package binarytree

import (
	"testing"
)

func TestFindPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				root: nil,
				sum:  0,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				root: NewBinaryTree([]int64{2, 2, 3, 4, 5, 4, 4, -1, -1, -2}, &BinaryTree{}),
				sum:  7,
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				root: NewBinaryTree([]int64{1, -1, 2, -1, -1, -1, 3, -1, -1, -1, -1, -1, 4}, &BinaryTree{}),
				sum:  3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("FindPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPathSum2(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				root: nil,
				sum:  0,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				root: NewBinaryTree([]int64{2, 2, 3, 4, 5, 4, 4, -1, -1, -2}, &BinaryTree{}),
				sum:  7,
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				root: NewBinaryTree([]int64{1, -1, 2, -1, -1, -1, 3, -1, -1, -1, -1, -1, 4}, &BinaryTree{}),
				sum:  3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPathSum2(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("FindPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
