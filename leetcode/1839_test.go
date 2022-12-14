package leetcode

import "testing"

func Test_longestBeautifulSubstring(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				word: "aeiaaioaaaaeiiiiouuuooaauuaeiu",
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestBeautifulSubstring(tt.args.word); got != tt.want {
				t.Errorf("longestBeautifulSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
