package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"time"
)

var fileMatch *regexp.Regexp
var pathPrefix = "./../"

func init() {
	fileMatch = regexp.MustCompile("^(.*)\\.in$")
}

func main() {
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

	rand.Seed(time.Now().UnixNano())

	for _, file := range files {
		if fileMatch.Match([]byte(file.Name())) {
			if err = orderPizza(file.Name()); err != nil {
				logError(err)
			}
		}
	}
}

func logError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "error: %+v\n", err)
}

func logErrorExit(err error, code int) {
	logError(err)
	os.Exit(code)
}

type Input struct {
	MaxSlices uint32
	Types     []uint32
}

type TypePair struct {
	Index     uint32
	NumSlices uint32
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

func orderPizza(filename string) error {
	input, err := parseInput(pathPrefix + filename)
	if err != nil {
		return err
	}
	fmt.Printf("input: %+v\n", input)

	types := make([]TypePair, len(input.Types))
	for i, n := range input.Types {
		types[i].Index = uint32(i)
		types[i].NumSlices = n
	}

	maxSum := uint32(0)
	var maxTypes []uint32
	for p := 0; p < 64; p++ {
		// shuffle only from the second time onwards
		if p > 0 {
			rand.Shuffle(len(types), func(i, j int) {
				types[i], types[j] = types[j], types[i]
			})
		}
		sum := uint32(0)
		var toOrder []uint32
		for _, t := range types {
			if t.NumSlices+sum <= input.MaxSlices {
				sum += t.NumSlices
				toOrder = append(toOrder, t.Index)
			}
		}
		fmt.Printf("%d:%d ", p, sum)
		if sum > maxSum {
			maxSum = sum
			maxTypes = toOrder
		}
		if sum == input.MaxSlices {
			break
		}
	}
	sort.Slice(maxTypes, func(i, j int) bool {
		return maxTypes[i] < maxTypes[j]
	})
	fmt.Printf("\nmaxSum=%d\n%+v\n\n", maxSum, maxTypes)
	return writeOutput(string(fileMatch.ReplaceAll([]byte(filename), []byte("$1.out"))), maxTypes)
}

func writeOutput(filename string, output []uint32) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			logError(err)
		}
	}()

	if _, err = fmt.Fprintln(file, len(output)); err != nil {
		return err
	}
	for i, n := range output {
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
