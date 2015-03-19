// Package board represents a square zupple board
// and its general operations
package board

// TODO: think on placing things in init()
// Just make sure what all things to add there

import (
	"github.com/pravj/puzzl/scanner"
	"math/rand"
	"time"
)

const (
        // SIZE of the puzzle board(fixed though)
	SIZE int = 3
)

// tile represents a tile in the zupple board
type tile struct {
	Value int
}

// row represents a row(list) of tiles in the zupple board
type row struct {
	size  int
	Tiles [SIZE]tile
}

// Board is a N*N size square puzzle board
type Board struct {
	size int
	Rows [SIZE]row

	BlankRow int
	BlankCol int
}

// New returns pointer to a newly created board instance
func New() *Board {
	board := &Board{size: SIZE}

	board.initiate()
	board.arrange()

	return board
}

// Initialize all the tile values to zero
func (b *Board) initiate() {
	var rows [SIZE]row

	for i := 0; i < b.size; i++ {
		r := row{size: SIZE}
		var tiles [SIZE]tile

		for j := 0; j < b.size; j++ {
			t := tile{Value: 0}
			tiles[j] = t
		}

		r.Tiles = tiles
		rows[i] = r
	}

	b.Rows = rows
}

// Arrange all the tiles in a given order
func (b *Board) arrange() {
	// returns a new Rand that uses unix timestamp as source for random value generation
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// slice of integer values randomly distributed
	values := r.Perm(SIZE * SIZE)
	//values := []int{8,6,7,2,5,4,3,0,1} // hard
	//values := []int{2,5,3,1,0,6,4,7,8} // easy
	//values := []int{1,2,3,7,8,4,0,5,6} // other

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			b.Rows[i].Tiles[j].Value = values[3*i+j]
		}
	}

	// validity of the configuration and position of blank tile
	valid, index := scanner.IsLegal(SIZE, values)
	if valid {
		b.BlankRow, b.BlankCol = position(index)
	} else {
		b.arrange()
	}
}

// Move shifts the blank tile from a tile configuration to another given
func (b *Board) Move(row, column int) {
	b.Rows[b.BlankRow].Tiles[b.BlankCol].Value = b.Rows[row].Tiles[column].Value

	b.BlankRow = row
	b.BlankCol = column

	b.Rows[b.BlankRow].Tiles[b.BlankCol].Value = 0
}

// position returns zero-based row and column position for a tile
func position(index int) (int, int) {
	return index / 3, index % 3
}

// Moves returns a list of all the possible moves from a given tile position
// TODO: find any idiomatic thing for this, if any.
func (b *Board) Moves(row, column int) []int {
	var move []int

	if column != 0 {
		move = append(move, row)
		move = append(move, column-1)
	}

	if column != 2 {
		move = append(move, row)
		move = append(move, column+1)
	}

	if row != 0 {
		move = append(move, row-1)
		move = append(move, column)
	}

	if row != 2 {
		move = append(move, row+1)
		move = append(move, column)
	}

	return move
}
