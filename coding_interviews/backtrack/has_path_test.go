package backtrack

import "testing"

func Test_hasPath(t *testing.T) {
	type args struct {
		matrix [][]byte
		word   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				matrix: [][]byte{
					{'a', 'b', 'c', 'e'}, {'s', 'f', 'c', 's'}, {'a', 'd', 'e', 'e'},
				},
				word: "abcced",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPath(tt.args.matrix, tt.args.word); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
