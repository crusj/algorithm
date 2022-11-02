package dp

import "testing"

func Test_bag(t *testing.T) {
	type args struct {
		m    []*bagv
		size int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				m: []*bagv{
					{
						w: 1,
						v: 1500,
					},
					{
						w: 3,
						v: 2000,
					},
					{
						w: 4,
						v: 3000,
					},
				},
				size: 4,
			},
			want: 3500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bag(tt.args.m, tt.args.size); got != tt.want {
				t.Errorf("bag() = %v, want %v", got, tt.want)
			}
		})
	}
}
