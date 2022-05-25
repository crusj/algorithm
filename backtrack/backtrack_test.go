package backtrack

import (
	"reflect"
	"testing"
)

func TestWholeArrangement(t *testing.T) {
	type args struct {
		nums      []int
		backtrack []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				nums:      []int{1, 2, 3},
				backtrack: []int{},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name: "test2",
			args: args{
				nums:      []int{1, 2},
				backtrack: []int{},
			},
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WholeArrangement(tt.args.nums, tt.args.backtrack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WholeArrangement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueenN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test8",
			args: args{
				n: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueenN(tt.args.n)
		})
	}
}
