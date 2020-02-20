package common

// Library provide custom type
type Library struct {
	ID     uint32
	Books  []*Book
	Score  float32
	SignUp uint32
	Ship   uint32 // Shiping Per Day
}

// CalcScore FUCK
func (l Library) CalcScore(remainingDays uint32, signUpCost float32) float32 {
	scores := float32(0)

	// TODO: remove duplicate books
	// we gon do that using pointers

	// asume books is already sorted
	if remainingDays > l.SignUp {

		// num of avail books
		canTake := l.Ship * (remainingDays - l.SignUp)
		if canTake > uint32(len(l.Books)) {
			canTake = uint32(len(l.Books))
		}

		// POINTERS remove scanned books
		/*var books []Book
		for _, book := range l.Books {
		}*/

		completeionCost := (float32(l.SignUp) + float32(canTake/l.Ship))

		availBooks := l.Books[:canTake]
		for _, book := range availBooks {
			scores += float32(book.Score)
			// print(scores)
		}
		return scores / completeionCost
	}
	return -1
}
