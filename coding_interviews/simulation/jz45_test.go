package simulation

import "testing"

func TestPrintMinNumber(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				numbers: []int{
					11, 3,
				},
			},
			want: "113",
		},
		{
			name: "test2",
			args: args{
				numbers: []int{},
			},
			want: "",
		},
		{
			name: "test3",
			args: args{
				numbers: []int{
					3, 32, 321,
				},
			},
			want: "321323",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintMinNumber(tt.args.numbers); got != tt.want {
				t.Errorf("PrintMinNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
