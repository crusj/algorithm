package simulation

import "testing"

func Test_cutRope(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				number: 8,
			},
			want: 18,
		},
		{
			name: "test2",
			args: args{
				number: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutRope(tt.args.number); got != tt.want {
				t.Errorf("cutRope() = %v, want %v", got, tt.want)
			}
		})
	}
}
