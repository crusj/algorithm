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
			want: []int{4, 3, 2},
		},
		{
			name: "test3",
			args: args{
				arr: []int{
					2, 3, 1, 5, 4,
				},
			},
			want: []int{5, 4, 1, 3, 2},
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

func Test_heapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{2, 3, 1, 5, 4},
			},
			want: []int{
				1, 2, 3, 4, 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
