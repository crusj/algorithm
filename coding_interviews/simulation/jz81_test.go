package simulation

import (
	"reflect"
	"testing"
)

func Test_reOrderArrayTwo(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				array: []int{1, 2, 3, 4},
			},
			want: []int{
				1, 3, 2, 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reOrderArrayTwo(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reOrderArrayTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
