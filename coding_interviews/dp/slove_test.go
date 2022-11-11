package dp

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		nums string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: "12",
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				nums: "31717126241541717",
			},
			want: 192,
		},
		{
			name: "test3",
			args: args{
				nums: "0",
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				nums: "10",
			},
			want: 1,
		},
		{
			name: "test5",
			args: args{
				nums: "7256224241262422162912798142114",
			},
			want: 7200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.nums); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
