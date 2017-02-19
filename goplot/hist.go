package goplot

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

// extend histogram flags from bar's flags
var histFlags = barFlags
var nBin int
var leftBound float64
var rightBound float64

func HistogramPlot(args []string) error {
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
		min, max := getBounds(values)
		if math.IsNaN(leftBound) {
			leftBound = min
		}
		if math.IsNaN(rightBound) {
			rightBound = max
		}
	}
	bins := groupValuesToBins(values, nBin, leftBound, rightBound)
	fmt.Printf("%d bins histogram of %d values from %g to %g\n", nBin, len(values), leftBound, rightBound)
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

func getBounds(values []float64) (min float64, max float64) {
	min = math.MaxFloat64
	max = -min
	for _, val := range values {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return
}

func groupValuesToBins(values []float64, nBin int, left float64, right float64) []LabeledValue {
	binSize := (right - left) / float64(nBin)
	bins := make([]LabeledValue, nBin)

	// label the bin by the upper/right bound
	for bin := 0; bin < nBin; bin++ {
		left_bound := left + float64(bin) * binSize
		right_bound := left_bound + binSize
		bins[bin].label = fmt.Sprintf("%.2f", right_bound)
	}

	for _, val := range values {
		switch {
		case val < left, val > right:
			continue
		case val == right:
			bins[nBin - 1].value++
		default:
			bins[int((val - left) / binSize)].value++
		}
	}
	return bins
}
