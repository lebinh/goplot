package plot

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Bound struct {
	left  float64
	right float64
}

// extend histogram flags from bar's flags
var histFlags = barFlags
var nBin int
var leftBound float64
var rightBound float64

func Histogram(args []string) error {
	// create histogram specific flag here because we reused bar plot flags for hist
	histFlags.IntVar(&nBin, "bin", 10, "number of bins in histogram")
	histFlags.Float64Var(&leftBound, "left", math.NaN(), "left bound of the histogram, default is min value")
	histFlags.Float64Var(&rightBound, "right", math.NaN(), "right bound of the histogram, default is max value")
	// replace the Usage method of histFlags since we cannot change its name from "bar"
	histFlags.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of hist:")
		histFlags.PrintDefaults()
	}
	histFlags.Parse(args)

	scanner := inputScanner(barFlags.Args())
	values, err := readValues(scanner)
	if err != nil {
		return err
	}

	// left, right bounds is min, max value by default, which is indicated by NaN
	if math.IsNaN(leftBound) || math.IsNaN(rightBound) {
		bound := getBounds(values)
		if math.IsNaN(leftBound) {
			leftBound = bound.left
		}
		if math.IsNaN(rightBound) {
			rightBound = bound.right
		}
	}

	bins := groupValuesToBins(values, nBin, Bound{leftBound, rightBound})
	drawBars(bins)
	return nil
}

func readValues(scanner *bufio.Scanner) ([]float64, error) {
	var values []float64
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return values, fmt.Errorf("input value is not a number: %v", err)
		}
		values = append(values, value)
	}
	return values, nil
}

func getBounds(values []float64) Bound {
	// let's just return [0, 0] as default bound for an empty slice
	if len(values) == 0 {
		return Bound{0, 0}
	}

	min := values[0]
	max := values[0]
	for _, val := range values[1:] {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return Bound{min, max}
}

func groupValuesToBins(values []float64, nBin int, bound Bound) []LabeledValue {
	binSize := (bound.right - bound.left) / float64(nBin)
	bins := make([]LabeledValue, nBin)

	// label the bin by the upper/right bound
	for bin := 0; bin < nBin; bin++ {
		right_bound := bound.left + float64(bin+1)*binSize
		bins[bin].label = fmt.Sprintf("%.2f", right_bound)
	}
	bins[0].label = fmt.Sprintf("%.2f -> %s", bound.left, bins[0].label)

	for _, val := range values {
		switch {
		case val < bound.left, val > bound.right:
			continue
		case val == bound.right:
			bins[nBin-1].value++
		default:
			bins[int((val-bound.left)/binSize)].value++
		}
	}
	return bins
}
