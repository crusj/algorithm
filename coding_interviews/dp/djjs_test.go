package dp

import "testing"

func Test_djjs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := djjs(tt.args.nums); got != tt.want {
				t.Errorf("djjs() = %v, want %v", got, tt.want)
			}
		})
	}
}
