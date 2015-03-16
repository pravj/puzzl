// package surface implements terminal GUI operations for the game
package surface

import "github.com/nsf/termbox-go"
import "github.com/pravj/puzzl/solver"
import "github.com/pravj/puzzl/board"
import "unicode/utf8"
import "strconv"
//import "fmt"
//import "os"

const (
  cornerUL rune = '┌'
  cornerUR rune = '┐'
  cornerLL rune = '└'
  cornerLR rune = '┘'

  hDash rune = '─'
  vDash rune = '│'

  blank rune = ' '
)

type Surface struct {
  gameBoard *board.Board
  gameSolver *solver.Solver

  currentBoard board.Board
}

func New(b *board.Board, s *solver.Solver) *Surface {
  sf := &Surface{gameBoard: b, gameSolver: s, currentBoard: *b}
  sf.initiate()

  return sf
}

func (s *Surface) drawCell(x, y int, ch rune) {
  termbox.SetCell(x, y, cornerUL, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+1, y, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+2, y, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+3, y, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+4, y, cornerUR, termbox.ColorDefault, termbox.ColorYellow)

  termbox.SetCell(x, y+1, vDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+1, y+1, blank, termbox.ColorDefault, termbox.ColorRed)
  termbox.SetCell(x+2, y+1, ch, termbox.ColorDefault, termbox.ColorDefault)
  termbox.SetCell(x+3, y+1, blank, termbox.ColorDefault, termbox.ColorRed)
  termbox.SetCell(x+4, y+1, vDash, termbox.ColorDefault, termbox.ColorYellow)

  termbox.SetCell(x, y+2, cornerLL, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+1, y+2, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+2, y+2, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+3, y+2, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+4, y+2, cornerLR, termbox.ColorDefault, termbox.ColorYellow)
}

func (s *Surface) drawBoard() {
  w, h := termbox.Size()
  const coldef = termbox.ColorDefault

  midy := h/2
  midx := w/2

  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      r, _ := utf8.DecodeRuneInString(strconv.Itoa(s.gameBoard.Rows[i].Tiles[j].Value))
      s.drawCell(midx + 5*j, midy + 3*i, r)
    }
  }

  termbox.Flush()
}

func (s *Surface) moveTile(dx, dy int) {
  newX, newY := s.gameBoard.BlankRow + dx, s.gameBoard.BlankCol + dy

  if (((newX >= 0) && (newX <= 2)) && ((newY >= 0) && (newY <= 2))) {
    // a possible move
    if (s.gameSolver.Solved) {
      // solved by solver
      s.gameBoard.Move(newX, newY)
      s.drawBoard()

      // right move by player
      if (s.gameSolver.Path[s.currentBoard] == *s.gameBoard) {
        s.currentBoard = *s.gameBoard
        // increase the score
      } else {
        // wrong move by player
      }

      // solved by player too
      if (*s.gameBoard == s.gameSolver.Goal) {
      }
    } else {
      // not yet solved by solver : NOTIFICATION -> WAIT
    }
  } else {
    // an impossible move : NOTIFICATION -> CAN NOT MOVE THERE
  }
}

func (s *Surface) initiate() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()

  termbox.SetInputMode(termbox.InputEsc)
  termbox.HideCursor()

  s.drawBoard()

  GameLoop:
    for {
      switch ev := termbox.PollEvent(); ev.Type {

      case termbox.EventKey:
        switch ev.Key {
        case termbox.KeyEsc:
          break GameLoop
        case termbox.KeyArrowUp:
          s.moveTile(-1, 0)
        case termbox.KeyArrowDown:
          s.moveTile(1, 0)
        case termbox.KeyArrowLeft:
          s.moveTile(0, -1)
        case termbox.KeyArrowRight:
          s.moveTile(0, 1)
        }

      case termbox.EventError:
        panic(ev.Err)
      }

      s.drawBoard()
    }
}
