package simulation

import "testing"

func TestIsContinuous(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				numbers: []int{
					1, 2, 3,
				},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				numbers: []int{
					3, 2, 1,
				},
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				numbers: []int{
					3, 2, 1, 5, 0,
				},
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				numbers: []int{
					3, 1, 1, 5, 0,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsContinuous(tt.args.numbers); got != tt.want {
				t.Errorf("IsContinuous() = %v, want %v", got, tt.want)
			}
		})
	}
}
