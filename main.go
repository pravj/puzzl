package main

import (
	"github.com/pravj/puzzl/board"
	"github.com/pravj/puzzl/notification"
	"github.com/pravj/puzzl/solver"
	"github.com/pravj/puzzl/surface"
)

func main() {
	gameBoard := board.New()

	gameNotification := notification.New()

	gameSolver := solver.New(gameBoard)
	go func() {
		gameSolver.Solve()
		gameNotification.Tunnel <- notification.WelcomeMessage
	}()

	surface.New(gameBoard, gameSolver, gameNotification)
}
