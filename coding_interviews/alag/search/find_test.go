package search

import (
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		target int
		array  [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				target: 7,
				array: [][]int{
					{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15},
				},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				target: 1,
				array: [][]int{
					{2},
				},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				target: 2,
				array: [][]int{
					{1, 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.target, tt.args.array); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind2(t *testing.T) {
	type args struct {
		target int
		array  [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				target: 7,
				array: [][]int{
					{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15},
				},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				target: 1,
				array: [][]int{
					{2},
				},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				target: 2,
				array: [][]int{
					{1, 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find2(tt.args.target, tt.args.array); got != tt.want {
				t.Errorf("Find2() = %v, want %v", got, tt.want)
			}
		})
	}
}
