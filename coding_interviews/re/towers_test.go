package re

import "testing"

func Test_tower(t *testing.T) {
	type args struct {
		n int
		a string
		b string
		c string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				n: 8,
				a: "a",
				b: "b",
				c: "c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tower(tt.args.n, tt.args.a, tt.args.b, tt.args.c)
		})
	}
}
