package leetcode

import "testing"

func Test_longestNiceSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 3, 8, 48, 10},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 1, 5, 11, 13},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestNiceSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestNiceSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
