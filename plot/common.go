package plot

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LabeledValue struct {
	Label string
	Value float64
}

func inputScanner(args []string) *bufio.Scanner {
	input := os.Stdin
	var err error
	if len(args) > 0 {
		input, err = os.Open(args[0])
		if err != nil {
			log.Fatalf("plot: %v\n", err)
		}
	}
	return bufio.NewScanner(input)
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

func readLabeledValues(scanner *bufio.Scanner) ([]LabeledValue, error) {
	var values []LabeledValue
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, *separator)
		if len(parts) != 2 {
			return values, fmt.Errorf("input is not in \"label%svalue\" format: %s", *separator, text)
		}

		label := parts[0]
		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return values, fmt.Errorf("input value is not a number: %v", err)
		}

		if value < 0 {
			return values, fmt.Errorf("bar plot with negative number is not supported yet: %v", value)
		}
		values = append(values, LabeledValue{label, value})
	}
	return values, nil
}
