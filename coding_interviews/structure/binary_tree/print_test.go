package binarytree

import (
	"reflect"
	"testing"
)

func TestPrintLine(t *testing.T) {
	type args struct {
		pRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "testnil",
			args: args{
				pRoot: nil,
			},
			want: nil,
		},

		{
			name: "test1",
			args: args{
				pRoot: NewBinaryTree([]int64{1, 2, 3, -1, -1, 4, 5}, &BinaryTree{}),
			},
			want: [][]int{
				{1},
				{2, 3},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintLine(tt.args.pRoot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
