package plot

import (
	"math"
	"testing"
)

const delta = 1e-6

//  ¯\_(ツ)_/¯
var box1 = `
╥
║
╨
`[1:]

var box2 = `
   ┌───┰──┐
├╶╶┤   ┃  ├╶╶╶┤
   └───┸──┘
`[1:]

var box3 = `
    ╓───┐
  ├╶╢   ├╶┤
    ╙───┘
`[1:]

var box4 = `
   ┌──╖
├╶╶┤  ╟╶╶╶┤
   └──╜
`[1:]

func TestDrawBox(t *testing.T) {
	tests := []struct {
		box  NormalizedBoxAndWhisker
		want string
	}{
		{NormalizedBoxAndWhisker{0, 0, 0, 0, 0}, box1},
		{NormalizedBoxAndWhisker{0, 3, 7, 10, 14}, box2},
		{NormalizedBoxAndWhisker{2, 4, 4, 8, 10}, box3},
		{NormalizedBoxAndWhisker{0, 3, 6, 6, 10}, box4},
	}
	for _, test := range tests {
		if got := drawBox(test.box); got != test.want {
			t.Errorf("drawBox(%v):\n%v, expected:\n%v", test.box, got, test.want)
		}
	}
}

func TestCalculateBoxAndWhisker(t *testing.T) {
	tests := []struct {
		values []float64
		want   BoxAndWhisker
	}{
		{[]float64{}, BoxAndWhisker{}},
		{[]float64{1}, BoxAndWhisker{1, 1, 1, 1, 1}},
		{[]float64{1, 2}, BoxAndWhisker{1, 1, 1.5, 2, 2}},
		{[]float64{1, 2, 3}, BoxAndWhisker{1, 1.5, 2, 2.5, 3}},
		{[]float64{1, 2, 3, 4}, BoxAndWhisker{1, 1.5, 2.5, 3.5, 4}},
		{[]float64{1, 2, 3, 4, 5}, BoxAndWhisker{1, 2, 3, 4, 5}},
		{[]float64{1, 2, 3, 4, 5}, BoxAndWhisker{1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		if got := calculateBoxAndWhisker(test.values); !equalBoxAndWhisker(got, test.want) {
			t.Errorf("calculateBoxAndWhisker(%v) = %v, expected: %v", test.values, got, test.want)
		}
	}
}

func equalBoxAndWhisker(b1, b2 BoxAndWhisker) bool {
	return closeEnough(b1.LeftWhisker, b2.LeftWhisker) &&
		closeEnough(b1.Left, b2.Left) &&
		closeEnough(b1.Mid, b2.Mid) &&
		closeEnough(b1.Right, b2.Right) &&
		closeEnough(b1.RightWhisker, b2.RightWhisker)
}

func closeEnough(f1, f2 float64) bool {
	return math.Abs(f1 - f2) < delta
}

func TestCalculateQuartiles(t *testing.T) {
	tests := []struct {
		values []float64
		want   Quartiles
	}{
		{[]float64{}, Quartiles{}},
		{[]float64{1}, Quartiles{1, 1, 1}},
		{[]float64{1, 2}, Quartiles{1, 1.5, 2}},
		{[]float64{1, 2, 3}, Quartiles{1.5, 2, 2.5}},
		{[]float64{1, 2, 3, 4}, Quartiles{1.5, 2.5, 3.5}},
		{[]float64{1, 2, 3, 4, 5}, Quartiles{2, 3, 4}},
	}
	for _, test := range tests {
		if got := quartilesOfSorted(test.values); !equalQuartiles(got, test.want) {
			t.Errorf("quartilesOfSorted(%v) = %v, expected: %v", test.values, got, test.want)
		}
	}
}

func TestMedianOfSorted(t *testing.T) {
	tests := []struct {
		values []float64
		want   float64
	}{
		{[]float64{}, 0},
		{[]float64{1}, 1},
		{[]float64{1, 2}, 1.5},
		{[]float64{1, 2, 3}, 2},
	}
	for _, test := range tests {
		if got := medianOfSorted(test.values); !closeEnough(got, test.want) {
			t.Errorf("medianOfSorted(%v) = %v, expected: %v", test.values, got, test.want)
		}
	}
}

func equalQuartiles(q1, q2 Quartiles) bool {
	return closeEnough(q1.First, q2.First) && closeEnough(q1.Second, q2.Second) && closeEnough(q1.Third, q2.Third)
}
