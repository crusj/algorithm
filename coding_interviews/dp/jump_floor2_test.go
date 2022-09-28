package dp

import "testing"

func Test_jumpFloorII(t *testing.T) {
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
				number: 3,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				number: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpFloorII(tt.args.number); got != tt.want {
				t.Errorf("jumpFloorII() = %v, want %v", got, tt.want)
			}
		})
	}
}
