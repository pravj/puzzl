package main

import (
	"github.com/pravj/puzzl/board"
	"github.com/pravj/puzzl/solver"
	"github.com/pravj/puzzl/surface"
	"github.com/pravj/puzzl/notification"
)

func main() {
	gameBoard := board.New()

        gameNotification := notification.New()

	gameSolver := solver.New(gameBoard)
        go func() {
          gameSolver.Solve()
          gameNotification.Tunnel <- "Try Now"
        }()

	surface.New(gameBoard, gameSolver, gameNotification)
}
