package leetcode

import "testing"

func Test_halfQuestions(t *testing.T) {
	type args struct {
		questions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				questions: []int{1, 2, 3, 1},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				questions: []int{1, 3, 14, 11, 1, 1, 11, 9, 3, 12},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halfQuestions(tt.args.questions); got != tt.want {
				t.Errorf("halfQuestions() = %v, want %v", got, tt.want)
			}
		})
	}
}
