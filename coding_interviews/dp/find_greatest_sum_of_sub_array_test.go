package dp

import (
	"reflect"
	"testing"
)

func TestFindGreatestSumOfSubArray(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				array: []int{1, -2, 3, 10, -4, 7, 2, -5},
			},
			want: 18,
		},
		{
			name: "test2",
			args: args{
				array: []int{2},
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				array: []int{-10},
			},
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindGreatestSumOfSubArrayDPReview(tt.args.array); got != tt.want {
				t.Errorf("FindGreatestSumOfSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindGreatestSumOfSubArrayMaxLen(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				array: []int{1, -2, 3, 10, -4, 7, 2, -5},
			},
			want: []int{3, 10, -4, 7, 2},
		},
		{
			name: "test2",
			args: args{
				array: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindGreatestSumOfSubArrayMaxLen(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindGreatestSumOfSubArrayMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
