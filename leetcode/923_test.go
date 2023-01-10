package leetcode

import "testing"

func Test_threeSumMulti(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				arr:    []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
				target: 8,
			},
			want: 20,
		},
		{
			name: "test2",
			args: args{
				arr:    []int{1, 1, 2, 2, 2, 2},
				target: 5,
			},
			want: 12,
		},
		{
			name: "test3",
			args: args{
				// 0, 0, 0, 2
				arr:    []int{0, 2, 0, 0},
				target: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumMulti(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("threeSumMulti() = %v, want %v", got, tt.want)
			}
		})
	}
}
