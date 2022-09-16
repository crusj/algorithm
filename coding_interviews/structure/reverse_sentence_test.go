package structure

import "testing"

func TestReverseSentence(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				str: "nowcoder. a am I",
			},
			want: "I am a nowcoder.",
		},
		{
			name: "test_null",
			args: args{
				str: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSentence(tt.args.str); got != tt.want {
				t.Errorf("ReverseSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
