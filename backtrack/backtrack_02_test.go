package backtrack

import "testing"

func Test_makesquare(t *testing.T) {
	type args struct {
		matchsticks []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				matchsticks: []int{
					1, 1, 2, 2, 2,
				},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				matchsticks: []int{
					3, 3, 3, 3, 4,
				},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				matchsticks: []int{
					5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare(tt.args.matchsticks); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
