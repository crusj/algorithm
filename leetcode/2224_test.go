package leetcode

import "testing"

func Test_convertTime(t *testing.T) {
	type args struct {
		current string
		correct string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				current: "02:30",
				correct: "04:35",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTime(tt.args.current, tt.args.correct); got != tt.want {
				t.Errorf("convertTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTimeStr2Minutes(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTimeStr2Minutes(tt.args.t); got != tt.want {
				t.Errorf("convertTimeStr2Minutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
