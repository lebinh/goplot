package main

import (
	"flag"
	"fmt"
	"github.com/lebinh/goplot/plot"
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
		err = plot.Bar(plotArgs)
	case "hist":
		err = plot.Histogram(plotArgs)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", flag.Arg(0))
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func Usage() {
	fmt.Fprintln(os.Stderr, "plot - terminal based stream plotting")
	fmt.Fprintln(os.Stderr, "Usage: \n\t plot (bar|hist) [options]")
}
