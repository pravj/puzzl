// Package notification implements inter process messaging for the game
package notification

type Notification struct {
  Tunnel chan string
}

func New() *Notification {
  return &Notification{Tunnel: make(chan string)}
}
