package dp

import "testing"

func Test_rectCover(t *testing.T) {
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
				number: 4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rectCover(tt.args.number); got != tt.want {
				t.Errorf("rectCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
