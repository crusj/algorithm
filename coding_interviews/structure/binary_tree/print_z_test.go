package binarytree

import (
	"reflect"
	"testing"
)

func TestPrint(t *testing.T) {
	type args struct {
		pRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				pRoot: nil,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				pRoot: NewBinaryTree([]int64{
					1, 2, 3, -1, -1, 4, 5,
				}, &BinaryTree{}),
			},
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Print(tt.args.pRoot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
