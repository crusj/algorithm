package linklist

import (
	"reflect"
	"testing"

	. "github.com/crusj/algorithm/common"
)

func Test_deleteDuplication(t *testing.T) {
	type args struct {
		pHead *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				pHead: nil,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				pHead: NewSingleList([]int{1, 1}),
			},
			want: nil,
		},
		{
			name: "test3",
			args: args{
				pHead: NewSingleList([]int{1, 2, 2}),
			},
			want: NewSingleList([]int{1}),
		},
		{
			name: "test4",
			args: args{
				pHead: NewSingleList([]int{1, 2, 2, 3}),
			},
			want: NewSingleList([]int{1, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplication(tt.args.pHead); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
