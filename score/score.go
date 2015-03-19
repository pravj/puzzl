// Package score implements the scoring functionality for the game.
//
// It keeps the track of user move correctness
// with the help of the game's in-built package solver.
package score

// Score represents struct that is meant to be scoring controller
type Score struct {
	PlayerTotal float64
	TotalMoves  float64
}

// New returns a pointer to a new Score struct instance
func New() *Score {
	return &Score{}
}

// Value returns the score value for a game
// using total number of moves and player's correct moves
func (s *Score) Value() float64 {
	if s.TotalMoves == 0 {
		return 0
	}

	return s.PlayerTotal / s.TotalMoves
}
