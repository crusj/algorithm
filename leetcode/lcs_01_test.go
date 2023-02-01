package leetcode

import "testing"

func Test_leastMinutes2(t *testing.T) {
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
				n: 2,
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				n: 100000,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastMinutes(tt.args.n); got != tt.want {
				t.Errorf("leastMinutes2() = %v, want %v", got, tt.want)
			}
		})
	}
}
