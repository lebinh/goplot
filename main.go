package main

import (
	"flag"
	"fmt"
	"github.com/lebinh/goplot/goplot"
	"os"
)

func main() {
	flag.Usage = Usage
	flag.Parse()
	if flag.NArg() == 0 {
		Usage()
		os.Exit(2)
	}

	plotArgs := flag.Args()[1:]
	var err error
	switch flag.Arg(0) {
	case "bar":
		err = goplot.BarPlot(plotArgs)
	case "hist":
		err = goplot.HistogramPlot(plotArgs)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", flag.Arg(0))
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func Usage() {
	fmt.Fprintln(os.Stderr, "goplot - terminal based stream plotting")
	fmt.Fprintln(os.Stderr, "Usage: \n\t goplot (bar|hist) [options]")
}
