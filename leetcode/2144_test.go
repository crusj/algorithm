package leetcode

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				cost: []int{6, 5, 7, 9, 2, 2},
			},
			want: 23,
		},
		{
			name: "test2",
			args: args{
				cost: []int{5, 5},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.cost); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
