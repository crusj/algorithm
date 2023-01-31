package leetcode

import (
	"reflect"
	"testing"
)

func Test_answerQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums:    []int{4, 5, 2, 1}, // [1,2,3,5]
				queries: []int{3, 10, 21},
			},
			want: []int{2, 3, 4},
		},
		{
			name: "test2",
			args: args{
				nums:    []int{624082},
				queries: []int{972985, 564269, 607119, 693641, 787608, 46517, 500857, 140097},
			},
			want: []int{1, 0, 0, 1, 1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
