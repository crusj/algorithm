package leetcode

import "testing"

func Test_largestSumAfterKNegations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{4, 2, 3},
				k:    7,
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				nums: []int{2, -3, -1, 5, -4},
				k:    2,
			},
			want: 13,
		},
		{
			name: "test3",
			args: args{
				nums: []int{5, 6, 9, -3, 3},
				k:    2,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumAfterKNegations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}
