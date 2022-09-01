package binarytree

import (
	"reflect"
	"testing"
)

func TestFindPath(t *testing.T) {
	type args struct {
		root         *TreeNode
		expectNumber int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		root:         NewBinaryTree([]int64{10, 5, 12, 4, 7}, &BinaryTree{}),
		// 		expectNumber: 22,
		// 	},
		// 	want: [][]int{
		// 		{10, 5, 7},
		// 		{10, 12},
		// 	},
		// },
		{
			name: "test2",
			args: args{
				root:         NewBinaryTree([]int64{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}, &BinaryTree{}),
				expectNumber: 22,
			},
			want: [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPath(tt.args.root, tt.args.expectNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
