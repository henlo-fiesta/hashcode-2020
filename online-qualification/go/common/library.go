package common

// Library provide custom type
type Library struct {
	ID     uint32
	Books  []*Book
	Score  int
	SignUp uint32
	Ship   uint32 // Shiping Per Day
}

// CalcScore FUCK
func (l Library) CalcScore(remainingDays uint32) int {
	// TODO: remove duplicate books
	// we gon do that using pointers

	// asume books is already sorted
	if remainingDays > l.SignUp {
		scores := 0

		// num of avail books
		canTake := l.Ship * (remainingDays - l.SignUp)
		if canTake > uint32(len(l.Books)) {
			canTake = uint32(len(l.Books))
		}
		// POINTERS remove scanned books
		/*var books []Book
		for _, book := range l.Books {
		}*/

		availBooks := l.Books[:canTake]
		for _, book := range availBooks {
			scores += int(book.Score)
			// print(scores)
		}
		return scores
	}
	return -1
}

// func (l Library)
