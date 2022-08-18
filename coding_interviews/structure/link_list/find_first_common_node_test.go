package linklist

import (
	"reflect"
	"testing"

	. "github.com/crusj/algorithm/common"
)

func Test_findFirstCommonNode(t *testing.T) {
	a := NewSingleList([]int{1, 2, 3})
	b := NewSingleList([]int{4, 5})
	c := NewSingleList([]int{6, 7})

	a.Next.Next.Next = c
	b.Next.Next = c

	m := NewSingleList([]int{1, 2})
	n := NewSingleList([]int{3, 4})

	type args struct {
		pHead1 *ListNode
		pHead2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				pHead1: a,
				pHead2: b,
			},
			want: c,
		},
		{
			name: "test2",
			args: args{
				pHead1: m,
				pHead2: n,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstCommonNode(tt.args.pHead1, tt.args.pHead2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFirstCommonNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
