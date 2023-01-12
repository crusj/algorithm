package leetcode

import (
	"reflect"
	"testing"
)

func Test_camelMatch(t *testing.T) {
	type args struct {
		queries []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "test1",
			args: args{
				queries: []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
				pattern: "FB",
			},
			want: []bool{true, false, true, true, false},
		},
		{
			name: "test2",
			args: args{
				queries: []string{
					"IXfGawluvnCa", "IsXfGaxwulCa", "IXfGawlqtCva", "IXjfGawlmeCa", "IXfGnaynwlCa", "IXfGcamwelCa",
				},
				pattern: "IXfGawlCa",
			},
			want: []bool{true, true, true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := camelMatch(tt.args.queries, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("camelMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
