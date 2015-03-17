// package surface implements terminal GUI operations for the game
package surface

import "github.com/nsf/termbox-go"
import "github.com/pravj/puzzl/solver"
import "github.com/pravj/puzzl/board"
import "github.com/pravj/puzzl/score"
import "unicode/utf8"
import "strconv"
import "fmt"
import "container/list"
import "os"

// rune type Box-drawing characters
// used to draw the game board interface
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

  currentBoard *list.Element
  solved bool

  scorer *score.Score
  solvableMoves int
}

func New(b *board.Board, s *solver.Solver) *Surface {
  scorer := score.New()
  sf := &Surface{gameBoard: b, gameSolver: s, scorer: scorer}
  sf.initiate()

  return sf
}

func (s *Surface) drawCell(x, y int, ch rune) {
  termbox.SetCell(x, y, cornerUL, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+1, y, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+2, y, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+3, y, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+4, y, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+5, y, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+6, y, cornerUR, termbox.ColorDefault, termbox.ColorYellow)

  termbox.SetCell(x, y+1, vDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+1, y+1, blank, termbox.ColorDefault, termbox.ColorRed)
  termbox.SetCell(x+2, y+1, blank, termbox.ColorDefault, termbox.ColorRed)
  termbox.SetCell(x+3, y+1, ch, termbox.ColorDefault, termbox.ColorDefault)
  termbox.SetCell(x+4, y+1, blank, termbox.ColorDefault, termbox.ColorRed)
  termbox.SetCell(x+5, y+1, blank, termbox.ColorDefault, termbox.ColorRed)
  termbox.SetCell(x+6, y+1, vDash, termbox.ColorDefault, termbox.ColorYellow)

  termbox.SetCell(x, y+2, cornerLL, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+1, y+2, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+2, y+2, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+3, y+2, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+4, y+2, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+5, y+2, hDash, termbox.ColorDefault, termbox.ColorYellow)
  termbox.SetCell(x+6, y+2, cornerLR, termbox.ColorDefault, termbox.ColorYellow)
}

func (s *Surface) drawWall(x, y int, isLeft bool) {
  for i := 1; i < 8; i++ {
    termbox.SetCell(x + 21, y + i, vDash, termbox.ColorDefault, termbox.ColorYellow)
  }

  if (isLeft) {
    termbox.SetCell(x + 21, y + 8, cornerLL, termbox.ColorDefault, termbox.ColorYellow)
  } else {
    termbox.SetCell(x + 21, y + 8, cornerLR, termbox.ColorDefault, termbox.ColorYellow)
  }
}

func (s *Surface) drawScore(x, y int) {
  // score banner
  chars := []rune{'G', 'A', 'M', 'E', ' ', 'S', 'C', 'O', 'R', 'E', ' ', ' '}
  for i := 0; i < 12; i++ {
    termbox.SetCell(x + 22 + i, y, chars[i], termbox.ColorDefault, termbox.ColorYellow)
  }

  // score value
  moves := fmt.Sprintf("%v", s.scorer.Value())
  length := len(moves)
  if (length > 5) {
    length = 5
  }

  var k int
  for i := 0; i < length; i++ {
    r, _ := utf8.DecodeRuneInString(string(moves[i]))
    termbox.SetCell(x + 22 + i, y + 1, r, termbox.ColorDefault, termbox.ColorYellow)
    k = i
  }

  for j := k+1; j < 12; j++ {
    termbox.SetCell(x + 22 + j, y + 1, blank, termbox.ColorDefault, termbox.ColorYellow)
  }
}

func (s *Surface) drawPlayerMoves(x, y int) {
  // player moves banner
  chars := []rune{'P', 'L', 'A', 'Y', 'E' ,'R', ' ', 'M', 'O', 'V', 'E', 'S'}
  for i := 0; i < 12; i++ {
    termbox.SetCell(x + 22 + i, y + 3, chars[i], termbox.ColorDefault, termbox.ColorYellow)
  }

  // player moves value
  moves := fmt.Sprintf("%v", s.scorer.TotalMoves)
  length := len(moves)

  var k int
  for i := 0; i < length; i++ {
    r, _ := utf8.DecodeRuneInString(string(moves[i]))
    termbox.SetCell(x + 22 + i, y + 4, r, termbox.ColorDefault, termbox.ColorYellow)
    k = i
  }

  for j := k+1; j < 12; j++ {
    termbox.SetCell(x + 22 + j, y + 4, blank, termbox.ColorDefault, termbox.ColorYellow)
  }
}

func (s *Surface) drawSolverMoves(x, y int) {
  // solver moves banner
  chars := []rune{'S', 'O', 'L', 'V', 'A', 'B', 'L', 'E', ' ', 'I', 'N', ' '}
  for i := 0; i < 12; i++ {
    termbox.SetCell(x + 22 + i, y + 6, chars[i], termbox.ColorDefault, termbox.ColorYellow)
  }

  // solver moves value
  moves := fmt.Sprintf("%v", s.solvableMoves)
  length := len(moves)

  var k int
  for i := 0; i < length; i++ {
    r, _ := utf8.DecodeRuneInString(string(moves[i]))
    termbox.SetCell(x + 22 + i, y + 7, r, termbox.ColorDefault, termbox.ColorYellow)
    k = i
  }

  for j := k+1; j < 12; j++ {
    termbox.SetCell(x + 22 + j, y + 7, blank, termbox.ColorDefault, termbox.ColorYellow)
  }
}

func (s *Surface) drawPartition(x, y int) {
  for i := 0; i < 12; i++ {
    termbox.SetCell(x + 22 + i, y + 8, hDash, termbox.ColorDefault, termbox.ColorYellow)
  }
}

func (s *Surface) drawBoard() {
  w, h := termbox.Size()
  const coldef = termbox.ColorDefault

  midy := h/2 - 5
  midx := w/2 - 15

  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      r, _ := utf8.DecodeRuneInString(strconv.Itoa(s.gameBoard.Rows[i].Tiles[j].Value))
      s.drawCell(midx + 7*j, midy + 3*i, r)
    }
  }

  s.drawWall(midx, midy, true)
  s.drawWall(midx+13, midy, false)

  s.drawScore(midx, midy)
  s.drawPlayerMoves(midx, midy)
  s.drawSolverMoves(midx, midy)

  s.drawPartition(midx, midy-3)
  s.drawPartition(midx, midy-6)
  s.drawPartition(midx, midy)

  termbox.Flush()
}

func (s *Surface) moveTile(dx, dy int) {
  newX, newY := s.gameBoard.BlankRow + dx, s.gameBoard.BlankCol + dy

  if (((newX >= 0) && (newX <= 2)) && ((newY >= 0) && (newY <= 2))) {
    // a possible move
    if (s.gameSolver.Solved) {
      if (!s.solved) {
        s.solved = true
        s.currentBoard = s.gameSolver.Path.Front()
        s.solvableMoves = s.gameSolver.Path.Len()
      }

      // solved by solver
      s.gameBoard.Move(newX, newY)

      // updates the total game moves played yet
      s.scorer.TotalMoves += 1

      // right move by player
      // NOTIFICATION -> RIGHT MOVE
      if (s.currentBoard.Value.(board.Board) == *s.gameBoard) {
        if (s.currentBoard.Next() != nil) {
          s.currentBoard = s.currentBoard.Next()

        }

        // updates the solvable moves count
        s.solvableMoves -= 1

        // increase the player's total
        s.scorer.PlayerTotal += 1
      } else {
        // wrong move by player
        // NOTIFICATION -> WRONG MOVE
        s.gameSolver = solver.New(s.gameBoard)
        s.gameSolver.Solve()
        s.solved = false
        s.solvableMoves = s.gameSolver.Path.Len()

        // decrease the player's total
        s.scorer.PlayerTotal -= 1
      }

      // draws the game board, will have the updated total moves also
      s.drawBoard()

      // solved by player too. Bingo.
      // NOTIFICATION -> GAME COMPLETE
      // WAIT FOR A WHILE AND EXIT THE PROCESS
      if (*s.gameBoard == s.gameSolver.Goal) {
        os.Exit(0)
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
