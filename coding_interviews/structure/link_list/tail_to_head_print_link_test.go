package linklist

import (
	"reflect"
	"testing"

	. "github.com/crusj/algorithm/common"
)

func Test_printListFromTailToHead(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
        {
        	name: "test1",
        	args: args{
        		head: NewSingleList([]int{1,2,3}),
        	},
        	want: []int{3,2,1},
        }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printListFromTailToHead(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printListFromTailToHead() = %v, want %v", got, tt.want)
			}
		})
	}
}
