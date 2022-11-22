package sort

import "testing"

func TestGetMedian(t *testing.T) {
	asserts := map[int]float64{
		1: 1.0,
		2: 1.5,
		3: 2.0,
		4: 2.5,
	}

	for _, v := range []int{1, 2, 3, 4} {
		Insert(v)
		tmp := GetMedian()
		if tmp != asserts[v] {
			t.Errorf("%d want %f get %f", v, asserts[v], tmp)
		}
	}
}
