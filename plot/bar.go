package plot

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

const block = "▇"
const tinyBlock = "▏"

var barFlags = flag.NewFlagSet("bar", flag.ExitOnError)
var separator = barFlags.String("sep", " ", "string used to separate values")
var maxWidth = barFlags.Int("width", 60, "maximum width of the plotted bar")

func Bar(args []string) error {
	barFlags.Parse(args)

	scanner := inputScanner(barFlags.Args())
	values, err := readLabeledValues(scanner)
	if err != nil {
		return err
	}
	drawBars(values)
	return nil
}

func drawBars(values []LabeledValue) {
	if len(values) == 0 {
		return
	}

	// iterate over the values once to get maximum values and label width
	max := values[0].Value
	maxLabelWidth := len(values[0].Label)
	for _, val := range values[1:] {
		if val.Value > max {
			max = val.Value
		}
		if len(val.Label) > maxLabelWidth {
			maxLabelWidth = len(val.Label)
		}
	}

	// normalize float64 values into int [0, maxWidth] range
	normalized := make([]int, len(values))
	for index, val := range values {
		normalized[index] = int(val.Value / max * float64(*maxWidth))
	}

	labelWidth := strconv.Itoa(maxLabelWidth)
	for index, bar := range values {
		rect := strings.Repeat(block, normalized[index])
		if normalized[index] == 0 {
			// use a small line for bar with 0 width
			rect = tinyBlock
		}
		fmt.Printf("%"+labelWidth+"s: %s  %g\n", bar.Label, rect, bar.Value)
	}
}
