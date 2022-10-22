package search

import "testing"

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				n: 10,
			},
			want: 1,
		},
		{
			name: "test4",
			args: args{
				n: 13,
			},
			want: 1,
		},
		{
			name: "test5",
			args: args{
				n: 500000000,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
