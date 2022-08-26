package binarytree

import "testing"

func TestTreeDepth(t *testing.T) {
	type args struct {
		pRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-nil",
			args: args{
				pRoot: nil,
			},
			want: 0,
		},
		{
			name: "test-1",
			args: args{
				pRoot: NewBinaryTree([]int64{1, 2, 3, 4, 5, -1, 6, -1, -1, 7}, &BinaryTree{}),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeDepth(tt.args.pRoot); got != tt.want {
				t.Errorf("TreeDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
