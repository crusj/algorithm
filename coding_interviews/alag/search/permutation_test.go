package search

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				str: "",
			},
			want: nil,
		},
		{
			name: "test1",
			args: args{
				str: "abc",
			},
			want: []string{
				"abc", "acb", "bac", "bca", "cab", "cba",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutation(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
