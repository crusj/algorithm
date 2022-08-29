package binarytree

import "testing"

func TestKthNode(t *testing.T) {
	type args struct {
		proot *TreeNode
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				proot: nil,
				k:     0,
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				proot: NewBinaryTree([]int64{1, 2}, &BinarySearchTree{}),
				k:     3,
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				proot: NewBinaryTree([]int64{5, 3, 7, 2, 4, 6, 8}, &BinarySearchTree{}),
				k:     3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KthNode(tt.args.proot, tt.args.k); got != tt.want {
				t.Errorf("KthNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
