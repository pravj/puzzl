// Package surface implements terminal GUI operations for the game
package surface

import (
	"container/list"
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/pravj/puzzl/board"
	"github.com/pravj/puzzl/notification"
	"github.com/pravj/puzzl/score"
	"github.com/pravj/puzzl/solver"
	"strconv"
	"unicode/utf8"
)

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

// Surface represents surface type struct
// It contains methods for general game terminal interface manipulation
type Surface struct {
	gameBoard  *board.Board
	gameSolver *solver.Solver

	currentBoard *list.Element
	solved       bool

	scorer        *score.Score
	solvableMoves int

	Message string

	Notifier          *notification.Notification
	NotificationColor termbox.Attribute

	channelClosed bool

	hintCount int
}

// New returns pointer to a new Surface instance
func New(b *board.Board, s *solver.Solver, n *notification.Notification) *Surface {
	scorer := score.New()
	sf := &Surface{gameBoard: b, gameSolver: s, scorer: scorer, Message: notification.WelcomeMessage, Notifier: n, NotificationColor: termbox.ColorCyan, hintCount: 3}

	sf.initiate()

	return sf
}

// Draws a cell(square) structure on terminal
// That consists the digit from the game board
func (s *Surface) drawCell(x, y int, ch rune) {
	// Red color for blank cell and Blue for others
	var bgColor termbox.Attribute
	if ch == '0' {
		bgColor = termbox.ColorRed
	} else {
		bgColor = termbox.ColorBlue
	}

	termbox.SetCell(x, y, cornerUL, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+1, y, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+2, y, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+3, y, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+4, y, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+5, y, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+6, y, cornerUR, termbox.ColorDefault, termbox.ColorCyan)

	termbox.SetCell(x, y+1, vDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+1, y+1, blank, termbox.ColorDefault, bgColor)
	termbox.SetCell(x+2, y+1, blank, termbox.ColorDefault, bgColor)
	termbox.SetCell(x+3, y+1, ch, termbox.ColorDefault, bgColor)
	termbox.SetCell(x+4, y+1, blank, termbox.ColorDefault, bgColor)
	termbox.SetCell(x+5, y+1, blank, termbox.ColorDefault, bgColor)
	termbox.SetCell(x+6, y+1, vDash, termbox.ColorDefault, termbox.ColorCyan)

	termbox.SetCell(x, y+2, cornerLL, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+1, y+2, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+2, y+2, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+3, y+2, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+4, y+2, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+5, y+2, hDash, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+6, y+2, cornerLR, termbox.ColorDefault, termbox.ColorCyan)
}

// Draws a vertical wall to separate sections
func (s *Surface) drawWall(x, y int, isLeft bool) {
	for i := 1; i < 8; i++ {
		termbox.SetCell(x+21, y+i, vDash, termbox.ColorDefault, termbox.ColorCyan)
	}

	if isLeft {
		termbox.SetCell(x+21, y+8, cornerLL, termbox.ColorDefault, termbox.ColorCyan)
	} else {
		termbox.SetCell(x+21, y+8, cornerLR, termbox.ColorDefault, termbox.ColorCyan)
	}
}

// Draws game's scoring section
func (s *Surface) drawScore(x, y int) {
	// fills the blank space of two ticks
	termbox.SetCell(x+21, y, blank, termbox.ColorDefault, termbox.ColorCyan)
	termbox.SetCell(x+34, y, blank, termbox.ColorDefault, termbox.ColorCyan)

	// score banner
	chars := []rune{'G', 'A', 'M', 'E', ' ', 'S', 'C', 'O', 'R', 'E', ' ', ' '}
	for i := 0; i < 12; i++ {
		termbox.SetCell(x+22+i, y, chars[i], termbox.ColorDefault, termbox.ColorYellow)
	}

	// score value
	moves := fmt.Sprintf("%v", s.scorer.Value())
	length := len(moves)
	if length > 5 {
		length = 5
	}

	var k int
	for i := 0; i < length; i++ {
		r, _ := utf8.DecodeRuneInString(string(moves[i]))
		termbox.SetCell(x+22+i, y+1, r, termbox.ColorDefault, termbox.ColorMagenta)
		k = i
	}

	for j := k + 1; j < 12; j++ {
		termbox.SetCell(x+22+j, y+1, blank, termbox.ColorDefault, termbox.ColorMagenta)
	}
}

// Draws game's player moves section
func (s *Surface) drawPlayerMoves(x, y int) {
	// player moves banner
	chars := []rune{'P', 'L', 'A', 'Y', 'E', 'R', ' ', 'M', 'O', 'V', 'E', 'S'}
	for i := 0; i < 12; i++ {
		termbox.SetCell(x+22+i, y+3, chars[i], termbox.ColorDefault, termbox.ColorYellow)
	}

	// player moves value
	moves := fmt.Sprintf("%v", s.scorer.TotalMoves)
	length := len(moves)

	var k int
	for i := 0; i < length; i++ {
		r, _ := utf8.DecodeRuneInString(string(moves[i]))
		termbox.SetCell(x+22+i, y+4, r, termbox.ColorDefault, termbox.ColorMagenta)
		k = i
	}

	for j := k + 1; j < 12; j++ {
		termbox.SetCell(x+22+j, y+4, blank, termbox.ColorDefault, termbox.ColorMagenta)
	}
}

// Draws game's solver moves section
func (s *Surface) drawSolverMoves(x, y int) {
	// solver moves banner
	chars := []rune{'S', 'O', 'L', 'V', 'A', 'B', 'L', 'E', ' ', 'I', 'N', ' '}
	for i := 0; i < 12; i++ {
		termbox.SetCell(x+22+i, y+6, chars[i], termbox.ColorDefault, termbox.ColorYellow)
	}

	// solver moves value
	moves := fmt.Sprintf("%v", s.solvableMoves)
	length := len(moves)

	var k int
	for i := 0; i < length; i++ {
		r, _ := utf8.DecodeRuneInString(string(moves[i]))
		termbox.SetCell(x+22+i, y+7, r, termbox.ColorDefault, termbox.ColorMagenta)
		k = i
	}

	for j := k + 1; j < 12; j++ {
		termbox.SetCell(x+22+j, y+7, blank, termbox.ColorDefault, termbox.ColorMagenta)
	}
}

// Draws horizontal partitioning between different sections
func (s *Surface) drawPartition(x, y int) {
	for i := 0; i < 12; i++ {
		termbox.SetCell(x+22+i, y+8, hDash, termbox.ColorDefault, termbox.ColorCyan)
	}
}

// Draws notifications in real time
func (s *Surface) drawNotification(x, y int, message string) {
	// notification widget boundary
	termbox.SetCell(x, y-3, cornerUL, termbox.ColorDefault, termbox.ColorBlue)
	termbox.SetCell(x, y-2, vDash, termbox.ColorDefault, termbox.ColorBlue)
	termbox.SetCell(x, y-1, cornerLL, termbox.ColorDefault, termbox.ColorBlue)

	for i := 0; i < 33; i++ {
		termbox.SetCell(x+1+i, y-3, hDash, termbox.ColorDefault, termbox.ColorBlue)
		termbox.SetCell(x+1+i, y-1, hDash, termbox.ColorDefault, termbox.ColorBlue)
	}

	termbox.SetCell(x+34, y-3, cornerUR, termbox.ColorDefault, termbox.ColorBlue)
	termbox.SetCell(x+34, y-2, vDash, termbox.ColorDefault, termbox.ColorBlue)
	termbox.SetCell(x+34, y-1, cornerLR, termbox.ColorDefault, termbox.ColorBlue)

	// notification message value
	var k int
	for i := 0; i < len(message); i++ {
		r, _ := utf8.DecodeRuneInString(string(message[i]))
		termbox.SetCell(x+1+i, y-2, r, termbox.ColorDefault, s.NotificationColor)
		k = i
	}

	for j := k + 2; j < 34; j++ {
		termbox.SetCell(x+j, y-2, blank, termbox.ColorDefault, s.NotificationColor)
	}
}

// Combines all the sections and draw the entire game board accordingly
func (s *Surface) drawBoard() {
	w, h := termbox.Size()
	const coldef = termbox.ColorDefault

	midy := h/2 - 5
	midx := w/2 - 15

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r, _ := utf8.DecodeRuneInString(strconv.Itoa(s.gameBoard.Rows[i].Tiles[j].Value))
			s.drawCell(midx+7*j, midy+3*i, r)
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

	s.drawNotification(midx, midy, s.Message)

	termbox.Flush()
}

// decide the game actions according to the game state
func (s *Surface) decideAction(dx, dy int) {
	if !s.channelClosed {
		s.moveTile(dx, dy)
	} else {
		// NOTIFICATION COLOR -> RED
		s.NotificationColor = termbox.ColorRed

		s.Message = notification.QuitMessage
		s.drawBoard()
	}
}

// move a tile in a given direction
func (s *Surface) moveTile(dx, dy int) {
	newX, newY := s.gameBoard.BlankRow+dx, s.gameBoard.BlankCol+dy

	// a possible move
	if ((newX >= 0) && (newX <= 2)) && ((newY >= 0) && (newY <= 2)) {
		// game has been solved by the solver
		if s.gameSolver.Solved {
			if !s.solved {
				s.solved = true
				s.currentBoard = s.gameSolver.Path.Front()
				s.solvableMoves = s.gameSolver.Path.Len()
			}

			// solved by solver
			s.gameBoard.Move(newX, newY)

			// updates the total game moves played till now
			s.scorer.TotalMoves++

			// right move by player
			if s.currentBoard.Value.(board.Board) == *s.gameBoard {
				if s.currentBoard.Next() != nil {
					s.currentBoard = s.currentBoard.Next()
				}

				// NOTIFICATION -> RIGHT MOVE
				s.Message = notification.RightMoveMessage

				// NOTIFICATION COLOR -> GREEN
				s.NotificationColor = termbox.ColorGreen

				// updates the solvable moves count
				s.solvableMoves--

				// increase the player's total
				s.scorer.PlayerTotal++
			} else {
				// wrong move by player
				s.gameSolver = solver.New(s.gameBoard)

				go func() {
					s.gameSolver.Solve()
					s.solved = false
					s.solvableMoves = s.gameSolver.Path.Len()

					// NOTIFICATION COLOR -> GREEN
					s.NotificationColor = termbox.ColorGreen

					s.Notifier.Tunnel <- notification.ReadyToPlayMessage
				}()

				// NOTIFICATION -> WRONG MOVE
				s.Message = notification.WrongMoveMessage

				// NOTIFICATION COLOR -> RED
				s.NotificationColor = termbox.ColorRed

				// decrease the player's total
				s.scorer.PlayerTotal--
			}

			// solved by player too. Bingo.
			// NOTIFICATION -> GAME COMPLETE
			if *s.gameBoard == s.gameSolver.Goal {
				s.Message = notification.GameCompleteMessage

				// NOTIFICATION COLOR -> CYAN
				s.NotificationColor = termbox.ColorCyan

				// update game status, close notification channel
				close(s.Notifier.Tunnel)
				s.channelClosed = true
			}

			// draws the game board, will have the updated total moves also
			s.drawBoard()

		} else {
			// not yet solved by solver : NOTIFICATION -> WAIT
			s.Message = notification.WaitMessage
			s.drawBoard()

			// NOTIFICATION COLOR -> Yellow
			s.NotificationColor = termbox.ColorYellow
		}
	} else {
		// an impossible move : NOTIFICATION -> CAN NOT MOVE THERE
		s.Message = notification.ImpossibleMoveMessage

		// NOTIFICATION COLOR -> RED
		s.NotificationColor = termbox.ColorRed

		s.drawBoard()
	}
}

// shows hints when asked, there is a limit for hints though
func (s *Surface) showHint() {
	if s.gameSolver.Solved && s.hintCount > 0 {
		present := s.gameBoard
		presentRow, presentCol := present.BlankRow, present.BlankCol

		future := s.currentBoard.Value.(board.Board)
		futureRow, futureCol := future.BlankRow, future.BlankCol

		var direction string

		if futureRow > presentRow {
			direction = "down"
		} else if futureRow < presentRow {
			direction = "up"
		} else if futureCol > presentCol {
			direction = "right"
		} else if futureCol < presentCol {
			direction = "left"
		}

		s.NotificationColor = termbox.ColorCyan
		s.hintCount--
		s.Message = fmt.Sprintf("Hint #%v - move %v side", 3-s.hintCount, direction)
	} else if s.gameSolver.Solved && s.hintCount <= 0 {
		s.NotificationColor = termbox.ColorRed
		s.Message = "No more hints my friend."
	} else {
		s.NotificationColor = termbox.ColorYellow
		s.Message = notification.WaitMessage
	}

	s.drawBoard()
}

// Initialize all the concurrent processes
// To monitor user input events and notification communication
func (s *Surface) initiate() {

	go func() {
		for e := range s.Notifier.Tunnel {
			// updates the solvable moves count for the game
			// it fixes the issue where game wasn't showing it in the starting
			s.solvableMoves = s.gameSolver.Path.Len()

			s.currentBoard = s.gameSolver.Path.Front()

			s.Message = e
			s.drawBoard()
		}
	}()

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
				if !s.channelClosed {
					close(s.Notifier.Tunnel)
				}
				break GameLoop
			case termbox.KeyArrowUp:
				s.decideAction(-1, 0)
			case termbox.KeyArrowDown:
				s.decideAction(1, 0)
			case termbox.KeyArrowLeft:
				s.decideAction(0, -1)
			case termbox.KeyArrowRight:
				s.decideAction(0, 1)
			}

			switch ev.Ch {
			case 'h':
				s.showHint()
			case 'H':
				s.showHint()
			}

		case termbox.EventError:
			panic(ev.Err)
		}

		s.drawBoard()
	}
}
