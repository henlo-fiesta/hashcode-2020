package main

type Input struct {
	MaxSlices uint32
	Types     []uint32
}

type TypePair struct {
	Index     uint32
	NumSlices uint32
}

type Output []uint32

type Strategy interface {
	Solve(input *Input) (*Output, error)
}
