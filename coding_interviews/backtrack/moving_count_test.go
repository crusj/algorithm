package backtrack

import "testing"

func Test_movingCount(t *testing.T) {
	type args struct {
		threshold int
		rows      int
		cols      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				threshold: 1,
				rows:      2,
				cols:      3,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				threshold: 0,
				rows:      1,
				cols:      3,
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				threshold: 10,
				rows:      1,
				cols:      100,
			},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.threshold, tt.args.rows, tt.args.cols); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
