// Package notification implements inter process messaging for the game
package notification

// Notification messages to show in the game
const (
	WelcomeMessage        string = "Welcome to the game Puzzl!"
	RightMoveMessage      string = "Woot! Right Move"
	WrongMoveMessage      string = "Oops! Wrong Move"
	GameCompleteMessage   string = "Woot! You completed the game"
	ImpossibleMoveMessage string = "Unable to move there"
	WaitMessage           string = "Wait! Let bot solve it first"
	ReadyToPlayMessage    string = "OK! You can play now"
	QuitMessage           string = "Press ESC key to quit"
)

// Notification struct
// It consists the channel used for internal notification communication
type Notification struct {
	Tunnel chan string
}

// New returns pointer to a Notification struct
func New() *Notification {
	return &Notification{Tunnel: make(chan string)}
}
