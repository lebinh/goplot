package plot

import (
	"reflect"
	"testing"
)

type List []float64

func TestGetBounds(t *testing.T) {
	tests := []struct {
		input List
		want  Bound
	}{
		{List{}, Bound{0, 0}},
		{List{1}, Bound{1, 1}},
		{List{1, 2, 3, -1}, Bound{-1, 3}},
		{List{1, 0.1, -1.0, -1.1, 1.0}, Bound{-1.1, 1}},
	}
	for _, test := range tests {
		if got := getBounds(test.input); got != test.want {
			t.Errorf("getBounds(%v) = %v, expected: %v", test.input, got, test.want)
		}
	}
}

func TestGroupValuesToBins(t *testing.T) {
	tests := []struct {
		input List
		nBin  int
		want  []Bin
	}{
		{List{}, 0, []Bin{}},
		{List{}, 1, []Bin{{Bound{0, 0}, 0}}},
		{List{1}, 1, []Bin{{Bound{1, 1}, 1}}},
		{List{1, 2}, 1, []Bin{{Bound{1, 2}, 2}}},
		{List{1, 2}, 2, []Bin{{Bound{1, 1.5}, 1}, {Bound{1.5, 2}, 1}}},
		{List{0, 1, 1, 2, 3}, 3, []Bin{{Bound{0, 1}, 1}, {Bound{1, 2}, 2}, {Bound{2, 3}, 2}}},
	}
	for _, test := range tests {
		bounds := getBounds(test.input)
		if got := groupValuesToBins(test.input, test.nBin, bounds); !reflect.DeepEqual(got, test.want) {
			t.Errorf("groupsValuesToBins(%v, %v, %v) = %v, expected: %v",
				test.input, test.nBin, bounds, got, test.want)
		}
	}
}
