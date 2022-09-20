package search

import "testing"

func Test_minNumberInRotateArray(t *testing.T) {
	type args struct {
		rotateArray []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				rotateArray: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				rotateArray: []int{3, 100, 200, 3},
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				rotateArray: []int{1, 0, 1, 1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberInRotateArray(tt.args.rotateArray); got != tt.want {
				t.Errorf("minNumberInRotateArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
