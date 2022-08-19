package linklist

import (
	"reflect"
	"testing"

	. "github.com/crusj/algorithm/common"
)

func TestFindKthToTail(t *testing.T) {
	type args struct {
		pHead *ListNode
		k     int
	}

	testLink := NewSingleList([]int{1, 2, 3, 4})
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				pHead: NewSingleList([]int{1, 2}),
				k:     3,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				pHead: testLink,
				k:     3,
			},
			want: testLink.Next,
		},
		{
			name: "test2",
			args: args{
				pHead: testLink,
				k:     4,
			},
			want: testLink,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthToTail(tt.args.pHead, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindKthToTail() = %v, want %v", got, tt.want)
			}
		})
	}
}
