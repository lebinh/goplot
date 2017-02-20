package plot

import (
	"testing"
	"reflect"
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
		want  []LabeledValue
	}{
		{List{}, 0, []LabeledValue{}},
		{List{}, 1, []LabeledValue{{"0.00", 0}}},
		{List{}, 2, []LabeledValue{{"0.00", 0}, {"0.00", 0}}},
		{List{1}, 1, []LabeledValue{{"1.00", 1}}},
		{List{1, 2}, 1, []LabeledValue{{"2.00", 2}}},
		{List{1, 2}, 2, []LabeledValue{{"1.50", 1}, {"2.00", 1}}},
		{List{0, 1, 1, 2, 3}, 3, []LabeledValue{{"1.00", 1}, {"2.00", 2}, {"3.00", 2}}},
	}
	for _, test := range tests {
		bounds := getBounds(test.input)
		if got := groupValuesToBins(test.input, test.nBin, bounds); !reflect.DeepEqual(got, test.want) {
			t.Errorf("groupsValuesToBins(%v, %v, %v) = %v, expected: %v",
				test.input, test.nBin, bounds, got, test.want)
		}
	}
}
