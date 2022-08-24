package linklist

import (
	"reflect"
	"testing"

	. "github.com/crusj/algorithm/common"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test-nil",
			args: args{
				head: nil,
				val:  0,
			},
			want: nil,
		},
		{
			name: "test-head",
			args: args{
				head: NewSingleList([]int{1, 2, 3}),
				val:  1,
			},
			want: NewSingleList([]int{2, 3}),
		},
		{
			name: "test-first",
			args: args{
				head: NewSingleList([]int{1}),
				val:  1,
			},
			want: nil,
		},
		{
			name: "test-middle",
			args: args{
				head: NewSingleList([]int{1, 2, 3}),
				val:  2,
			},
			want: NewSingleList([]int{1, 3}),
		},
		{
			name: "test-last",
			args: args{
				head: NewSingleList([]int{1, 2, 3}),
				val:  3,
			},
			want: NewSingleList([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
