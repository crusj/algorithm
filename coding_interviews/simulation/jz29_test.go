package simulation

import (
	"reflect"
	"testing"
)

func Test_jz29(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				arr: [][]int{
					{
						1, 2, 3,
					},
					{
						4, 5, 6,
					},
					{
						7, 8, 9,
					},
				},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "test2",
			args: args{
				arr: [][]int{
					{1}, {2}, {3},
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jz29(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jz29() = %v, want %v", got, tt.want)
			}
		})
	}
}
