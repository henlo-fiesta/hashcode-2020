package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type randomStrategy struct {
}

func NewRandomStrategy() (Strategy, error) {
	rand.Seed(time.Now().UnixNano())
	return &randomStrategy{}, nil
}

func (*randomStrategy) Solve(input *Input) (*Output, error) {
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

	return (*Output)(&maxTypes), nil
}
