package simulation

import (
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	node1 := &RandomListNode{
		Label: 1,
	}
	node2 := &RandomListNode{
		Label: 2,
	}
	node3 := &RandomListNode{
		Label: 3,
	}
	node4 := &RandomListNode{
		Label: 4,
	}
	node5 := &RandomListNode{
		Label: 5,
	}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node1.Random = node3
	node2.Random = node5
	node4.Random = node2

	type args struct {
		head *RandomListNode
	}
	tests := []struct {
		name string
		args args
		want *RandomListNode
	}{
		{
			name: "test1",
			args: args{
				head: node1,
			},
			want: &RandomListNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clone(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				if got.Next.Random == nil {
					t.Error()
				}
			}
		})
	}
}
