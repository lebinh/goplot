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

type Bin struct {
	bound Bound
	count int
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
	drawBins(bins)
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

func groupValuesToBins(values []float64, nBin int, bound Bound) []Bin {
	binSize := (bound.right - bound.left) / float64(nBin)
	bins := make([]Bin, nBin)

	// label the bin by the upper/right bound
	for bin := 0; bin < nBin; bin++ {
		leftBound := bound.left + float64(bin)*binSize
		rightBound := leftBound + binSize
		bins[bin].bound = Bound{leftBound, rightBound}
	}

	for _, val := range values {
		switch {
		case val < bound.left, val > bound.right:
			continue
		case val == bound.right:
			bins[nBin-1].count++
		default:
			bins[int((val-bound.left)/binSize)].count++
		}
	}
	return bins
}

func drawBins(bins []Bin) {
	bars := make([]LabeledValue, len(bins))
	for i, bin := range bins {
		bars[i].label = fmt.Sprintf("%.2f", bin.bound.right)
	}
	if len(bins) > 1 {
		bars[0].label = fmt.Sprintf("%.2f -> %.2f", bins[0].bound.left, bins[0].bound.right)
	}
	drawBars(bars)
}
