package dp

import "testing"

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				word1: "abc",
				word2: "abb",
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				word1: "abbbbbbb",
				word2: "abb",
			},
			want: 5,
		},
		{
			name: "test3",
			args: args{
				word1: "abcd",
				word2: "asbc",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
