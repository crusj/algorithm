package leetcode

import (
	"reflect"
	"testing"
)

func Test_numMovesStonesII(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				stones: []int{7, 4, 9},
			},
			want: []int{1, 2},
		}, {
			name: "test2",
			args: args{
				stones: []int{6, 5, 4, 3, 10},
			},
			want: []int{
				2, 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMovesStonesII(tt.args.stones); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numMovesStonesII() = %v, want %v", got, tt.want)
			}
		})
	}
}
