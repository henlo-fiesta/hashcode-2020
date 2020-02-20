package common

import "fmt"

// Book is a type of all books
type Book struct {
	ID    uint32
	Score uint32
}

// type Book map[uint32]uint32

func (b *Book) String() string {
	return fmt.Sprintf("{i:%d s:%d}", b.ID, b.Score)
}
