package review

import (
	"reflect"
	"testing"

	. "github.com/crusj/algorithm/common"
)

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

func TestMirror(t *testing.T) {
	type args struct {
		pRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
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
				pRoot: NewBinaryTree([]int64{8, 6, 10}, &BinaryTree{}),
			},
			want: NewBinaryTree([]int64{8, 10, 6}, &BinaryTree{}),
		},
		{
			name: "test3",
			args: args{
				pRoot: NewBinaryTree([]int64{8, 6, 10, 5, 7, 9, 11}, &BinaryTree{}),
			},
			want: NewBinaryTree([]int64{8, 10, 6, 11, 9, 7, 5}, &BinaryTree{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mirror(tt.args.pRoot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mirror() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			name: "test3",
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

func TestIsBalanced_Solution(t *testing.T) {
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
				pRoot: nil,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				pRoot: NewBinaryTree([]int64{1, 2, 3, 4, 5, 6, 7}, &BinaryTree{}),
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				pRoot: NewBinaryTree([]int64{1, 2, 3, 4, 5, -1, 6, -1, -1, 7}, &BinaryTree{}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalanced_Solution(tt.args.pRoot); got != tt.want {
				t.Errorf("IsBalanced_Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			want: 3,
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

func TestHasSubtree(t *testing.T) {
	type args struct {
		pRoot1 *TreeNode
		pRoot2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				pRoot1: nil,
				pRoot2: nil,
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				pRoot1: NewBinaryTree([]int64{8, 8, 7, 9, 2, -1, -1, -1, 4, 7}, &BinaryTree{}),
				pRoot2: NewBinaryTree([]int64{8, 9, 2}, &BinaryTree{}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasSubtree(tt.args.pRoot1, tt.args.pRoot2); got != tt.want {
				t.Errorf("HasSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
