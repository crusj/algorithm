package common

import (
	"reflect"
	"testing"
)

func TestNewSingleList(t *testing.T) {
	type args struct {
		nodes []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				nodes: []int{1, 2, 3},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingleList(tt.args.nodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingleList() = %v, want %v", got, tt.want)
			}
		})
	}
}
