package dp

import "testing"

func Test_jumpFloor(t *testing.T) {
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
				number: 2,
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				number: 7,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpFloor(tt.args.number); got != tt.want {
				t.Errorf("jumpFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}
