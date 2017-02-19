package goplot

import (
	"bufio"
	"os"
	"log"
)

func inputScanner(args []string) *bufio.Scanner {
	input := os.Stdin
	var err error
	if len(args) > 0 {
		input, err = os.Open(args[0])
		if err != nil {
			log.Fatalf("goplot: %v\n", err)
		}
	}
	return bufio.NewScanner(input)
}

