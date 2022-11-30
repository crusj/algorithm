package simulation

import "testing"

func TestLastRemaining_Solution(t *testing.T) {
	type args struct {
		n int
		m int
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
				n: 6,
				m: 7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastRemaining_Solution(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("LastRemaining_Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
