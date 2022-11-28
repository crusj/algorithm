package simulation

import (
	"reflect"
	"testing"
)

func TestFindContinuousSequence(t *testing.T) {
	type args struct {
		sum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				sum: 9,
			},

			want: [][]int{
				{
					2, 3, 4,
				}, {
					4, 5,
				},
			},
		},
		{
			name: "test2",
			args: args{
				sum: 1,
			},
			want: nil,
		},
		{
			name: "test15",
			args: args{
				sum: 15,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindContinuousSequence(tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindContinuousSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
