package goplot

import (
	"bufio"
	"log"
	"os"
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
