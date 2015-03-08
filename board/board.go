// Package board represents a square puzzle board
// and its general operations
package board

import (
  "fmt"
  "time"
  "math/rand"
  "github.com/pravj/puzzl/scanner"
)

const (
  SIZE int = 3 // size of the puzzle board(fixed though)
)

// Tile represents a tile in the puzzle board
type tile struct {
  value int
}

// Row represents a row(list) of tiles in the puzzle board
type row struct {
  size int
  tiles [SIZE]tile
}

// N*N size square puzzle board
type Board struct {
  size int
  rows [SIZE]row
}

// Returns a newly created board for the puzzle
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
      t := tile{value: 0}
      tiles[j] = t
    }

    r.tiles = tiles
    rows[i] = r
  }

  b.rows = rows
}

// Arrange all the tiles in a given order
func (b *Board) arrange() {
  // returns a new Rand that uses unix timestamp as source for random value generation
  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  // slice of integer values randomly distributed
  values := r.Perm(SIZE*SIZE)

  for i := 0; i < SIZE; i++ {
    for j := 0; j < SIZE; j++ {
      b.rows[i].tiles[j].value = values[3*i + j]
    }
  }

  fmt.Println(scanner.IsLegal(SIZE, values))
}
