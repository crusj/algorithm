package leetcode

import (
	"testing"
)

func Test_getSumPrefix(t *testing.T) {
	type args struct {
		arr []int
		i   int
		j   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6},
				i:   1,
				j:   3,
			},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getSumPrefix(tt.args.arr, tt.args.i, tt.args.j); gotAns != tt.wantAns {
				t.Errorf("getSumPrefix() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			wantAns: 2,
		},
		{
			name: "test21",
			args: args{
				nums: []int{1},
				k:    0,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := subarraySum(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("subarraySum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_longestWPI(t *testing.T) {
	type args struct {
		hours []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				hours: []int{
					9, 9, 6, 0, 1, 0,
				},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				hours: []int{
					9, 1, 3, 9, 7, 9, 3, 0,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWPI(tt.args.hours); got != tt.want {
				t.Errorf("longestWPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
