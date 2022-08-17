package linklist

import (
	"reflect"
	"testing"

	. "github.com/crusj/algorithm/common"
)

func TestMerge(t *testing.T) {
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
				pHead1: NewSingleList([]int{1, 2, 3}),
				pHead2: NewSingleList([]int{4, 5, 6}),
			},
			want: NewSingleList([]int{1, 2, 3, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.pHead1, tt.args.pHead2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %+v", got, tt.want)
			}
		})
	}
}
