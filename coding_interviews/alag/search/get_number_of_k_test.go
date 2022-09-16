package search

import "testing"

func TestGetNumberOfK(t *testing.T) {
	type args struct {
		data []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				data: []int{1, 2, 3, 3, 3, 3, 4, 5},
				k:    3,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				data: []int{1, 3, 4, 5},
				k:    6,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumberOfK(tt.args.data, tt.args.k); got != tt.want {
				t.Errorf("GetNumberOfK() = %v, want %v", got, tt.want)
			}
		})
	}
}
