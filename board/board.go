// Package board represents a square puzzle board
// and its general operations
package board

// TODO: think on placing things in init()

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

  blankRow int
  blankCol int
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

  // validity of the configuration and position of blank tile
  valid, index := scanner.IsLegal(SIZE, values)
  if (valid) {
    blankRow, blankCol := position(index)
    fmt.Println(moves(blankRow, blankCol))
    // LINK IT FOR FURTHER USE
  } else {
    b.arrange()
  }
}

// Returns zero-based row and column position for a tile
func position(index int) (int, int) {
  return index/3, index%3
}

// Returns a list of all the possible moves from a given tile position
// TODO: find any idiomatic thing for this, if any.
func moves(row, column int) ([]int) {
  move := make([]int, 8) // 8 because, it can move in maximum 4 directions

  if (column != 0) {
    move = append(move, row)
    move = append(move, column-1)
  }

  if (column != 2) {
    move = append(move, row)
    move = append(move, column+1)
  }

  if (row != 0) {
    move = append(move, row-1)
    move = append(move, column)
  }

  if (row != 2) {
    move = append(move, row+1)
    move = append(move, column)
  }

  return move
}
