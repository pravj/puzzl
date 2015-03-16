// package surface implements terminal GUI operations for the game
package surface

import "github.com/nsf/termbox-go"

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
  id int
}

func New() *Surface {
  sf := &Surface{}
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
  termbox.SetCell(x+2, y+1, ch, termbox.ColorDefault, termbox.ColorRed)
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

  arr := []rune{'1','2','3','4','5','6','7','8','0'}

  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      s.drawCell(midx + 5*j, midy + 3*i, arr[3*i+j])
    }
  }

  termbox.Flush()
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

Loop:
  for {
    switch ev := termbox.PollEvent(); ev.Type {

    case termbox.EventKey:
      switch ev.Key {
      case termbox.KeyEsc:
        break Loop
      }

    case termbox.EventError:
      panic(ev.Err)
    }

    s.drawBoard()
  }
}
