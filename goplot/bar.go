package goplot

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"flag"
	"bufio"
)

const block = "▇"
const tinyBlock = "▏"

type LabeledValue struct {
	label string
	value float64
}

var barFlags = flag.NewFlagSet("bar", flag.ExitOnError)
var separator = barFlags.String("sep", " ", "string used to separate values")
var maxWidth = barFlags.Int("width", 60, "maximum width of the plotted bar")

func BarPlot(args []string) error {
	barFlags.Parse(args)

	scanner := inputScanner(barFlags.Args())
	values, err := readLabeledValues(scanner)
	if err != nil {
		return err
	}
	drawBars(values)
	return nil
}

func readLabeledValues(scanner *bufio.Scanner) ([]LabeledValue, error) {
	var values []LabeledValue
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, *separator)
		if len(parts) != 2 {
			return values, fmt.Errorf("bar goplot: input is not in \"label%svalue\" format: %s", *separator, text)
		}

		label := parts[0]
		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return values, fmt.Errorf("bar goplot: input value is not a number: %v", err)
		}

		if value < 0 {
			return values, fmt.Errorf("bar goplot: goplot with negative number is not supported yet: %v", value)
		}
		values = append(values, LabeledValue{label, value})
	}
	return values, nil
}

func drawBars(values []LabeledValue) {
	// iterate over the values once to get maximum values and label width
	max := -math.MaxFloat64
	maxLabelWidth := 0
	for _, val := range values {
		if val.value > max {
			max = val.value
		}
		if len(val.label) > maxLabelWidth {
			maxLabelWidth = len(val.label)
		}
	}

	// normalize float64 values into int [0, maxWidth] range
	normalized := make([]int, len(values))
	for index, bar := range values {
		normalized[index] = int(bar.value / max * float64(*maxWidth))
	}

	labelWidth := strconv.Itoa(maxLabelWidth)
	for index, bar := range values {
		rect := strings.Repeat(block, normalized[index])
		if normalized[index] == 0 {
			// use a small line for bar with 0 width
			rect = tinyBlock
		}
		fmt.Printf("%" + labelWidth + "s: %s  %g\n", bar.label, rect, bar.value)
	}
}
