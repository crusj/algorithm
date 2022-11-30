package simulation

import "testing"

func TestLeftRotateString(t *testing.T) {
	type args struct {
		str string
		n   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				str: "abc",
				n:   10,
			},
			want: "bca",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftRotateString(tt.args.str, tt.args.n); got != tt.want {
				t.Errorf("LeftRotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
