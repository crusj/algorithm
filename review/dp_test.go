package review

import (
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
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				array: []int{1, -2, 3, 10, -4, 7, 2, -5},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindGreatestSumOfSubArray(tt.args.array); got != tt.want {
				t.Errorf("FindGreatestSumOfSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinSumPath(t *testing.T) {
	type args struct {
		paths [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				paths: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 12,
		},
		{
			name: "test2",
			args: args{
				paths: [][]int{
					{2, 1, 3},
					{6, 5, 4},
					{7, 8, 9},
				},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSumPath(tt.args.paths); got != tt.want {
				t.Errorf("MinSumPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFallingPathSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 12,
		},
		{
			name: "test2",
			args: args{
				matrix: [][]int{
					{2, 1, 3},
					{6, 5, 4},
					{7, 8, 9},
				},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinFallingPathSum(tt.args.matrix); got != tt.want {
				t.Errorf("MinFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxValue(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				grid: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 12,
		},
		{
			name: "test2",
			args: args{
				grid: [][]int{
					{9, 1, 4, 8},
				},
			},
			want: 22,
		},
		{
			name: "test3",
			args: args{
				grid: [][]int{
					{1, 2, 5},
					{3, 2, 1},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.grid); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lis(t *testing.T) {
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
				array: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				array: []int{2, 1, 4, 3, 5, 7, 6},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lis(tt.args.array); got != tt.want {
				t.Errorf("lis() = %v, want %v", got, tt.want)
			}
		})
	}
}
