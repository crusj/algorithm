package sort

import (
	"reflect"
	"testing"
)

func Test_maxHeap(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test2",
			args: args{
				arr: []int{
					2, 3, 4,
				},
			},
			want: []int{4, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxHeap(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
