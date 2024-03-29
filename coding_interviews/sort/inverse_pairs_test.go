package sort

import "testing"

func TestInversePairs(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				data: []int{1, 2, 3, 4, 0},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				data: []int{1, 2, 3, 4},
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				data: []int{1, 2, 3, 4},
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				data: []int{364, 637, 341, 406, 747, 995, 234, 971, 571, 219, 993, 407, 416, 366, 315, 301, 601, 650, 418, 355, 460, 505, 360, 965, 516, 648, 727, 667, 465, 849, 455, 181, 486, 149, 588, 233, 144, 174, 557, 67, 746, 550, 474, 162, 268, 142, 463, 221, 882, 576, 604, 739, 288, 569, 256, 936, 275, 401, 497, 82, 935, 983, 583, 523, 697, 478, 147, 795, 380, 973, 958, 115, 773, 870, 259, 655, 446, 863, 735, 784, 3, 671, 433, 630, 425, 930, 64, 266, 235, 187, 284, 665, 874, 80, 45, 848, 38, 811, 267, 575},
			},
			want: 2519,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InversePairs(tt.args.data); got != tt.want {
				t.Errorf("InversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
