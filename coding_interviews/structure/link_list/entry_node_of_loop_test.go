package linklist

import (
	"reflect"
	"testing"

	. "github.com/crusj/algorithm/common"
)

func TestEntryNodeOfLoop(t *testing.T) {
	a := NewSingleList([]int{1, 2, 3})
	b := NewSingleList([]int{4, 5})

	a.Next.Next.Next = b
	b.Next.Next = a.Next.Next

	type args struct {
		pHead *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				pHead: a,
			},
			want: a.Next.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EntryNodeOfLoop(tt.args.pHead); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EntryNodeOfLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
