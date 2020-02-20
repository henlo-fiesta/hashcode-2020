package main

import (
	"fmt"
	"os"
)

func logError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "error: %+v\n", err)
}

func logErrorExit(err error, code int) {
	logError(err)
	os.Exit(code)
}

func parseInput(filename string) (*Input, error) {
	fmt.Printf("parsing input '%s'...\n", filename)
	var input Input
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			logError(err)
		}
	}()

	var numTypes int
	if _, err = fmt.Fscanf(file, "%d %d\n", &input.MaxSlices, &numTypes); err != nil {
		return nil, err
	}
	input.Types = make([]uint32, numTypes)
	for i := range input.Types {
		if _, err = fmt.Fscanf(file, "%d", &input.Types[i]); err != nil {
			return nil, err
		}
	}
	return &input, nil
}

func writeOutput(filename string, output *Output) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			logError(err)
		}
	}()

	if _, err = fmt.Fprintln(file, len(*output)); err != nil {
		return err
	}
	for i, n := range *output {
		if i == 0 {
			_, err = fmt.Fprintf(file, "%d", n)
		} else {
			_, err = fmt.Fprintf(file, " %d", n)
		}
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintln(file)
	return err
}
