package binarytree

import (
	"reflect"
	"testing"
)

func TestPrintFromTopToBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				root: nil,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				root: NewBinaryTree([]int64{8, 6, 10, -1, -1, 2, 1}, &BinaryTree{}),
			},
			want: []int{
				8, 6, 10, 2, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintFromTopToBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintFromTopToBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
