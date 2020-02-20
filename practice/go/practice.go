package main

import (
	"fmt"
	"os"
	"regexp"
)

var fileMatch *regexp.Regexp
var pathPrefix = "./../"
var strategy Strategy

func init() {
	fileMatch = regexp.MustCompile("^(.*)\\.in$")
}

func main() {
	// TODO select a strategy
	var err error
	strategy, err = NewRandomStrategy()
	if err != nil {
		logErrorExit(err, 3)
	}

	dir, err := os.Open(pathPrefix)
	if err != nil {
		logErrorExit(err, 1)
	}
	files, err := dir.Readdir(0)
	if err != nil {
		logErrorExit(err, 2)
	}
	if err := dir.Close(); err != nil {
		logError(err)
	}

	for _, file := range files {
		if fileMatch.Match([]byte(file.Name())) {
			if err = orderPizza(file.Name()); err != nil {
				logError(err)
			}
		}
	}
}

func orderPizza(filename string) error {
	input, err := parseInput(pathPrefix + filename)
	if err != nil {
		return err
	}
	fmt.Printf("input: %+v\n", input)

	output, err := strategy.Solve(input)
	if err != nil {
		return err
	}

	sum := uint32(0)
	for _, i := range *output {
		sum += input.Types[i]
	}
	fmt.Printf("\nmaxSum=%d\n%+v\n\n", sum, *output)
	return writeOutput(string(fileMatch.ReplaceAll([]byte(filename), []byte("$1.out"))), output)
}
