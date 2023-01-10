package leetcode

import "testing"

func Test_minimumCardPickup(t *testing.T) {
	type args struct {
		cards []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				cards: []int{3, 4, 2, 3, 4, 7},
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				cards: []int{1, 0, 5, 3},
			},
			want: -1,
		},
		{
			name: "case3",
			args: args{
				cards: []int{95, 11, 8, 65, 5, 86, 30, 27, 30, 73, 15, 91, 30, 7, 37, 26, 55, 76, 60, 43, 36, 85, 47, 96, 6},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCardPickup(tt.args.cards); got != tt.want {
				t.Errorf("minimumCardPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}
