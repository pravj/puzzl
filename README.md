puzzl
=====

> An intelligent version of the sliding-puzzle game for your terminal built in golang

[![GoDoc](https://godoc.org/github.com/pravj/puzzl?status.svg)](http://godoc.org/github.com/pravj/puzzl)

puzzl is an intelligent implementation of the classical sliding-puzzle game.

It's built on top of Golang's *concurrency primitives*. It uses *goroutines* and *channels* for inter process communications to provide real time notification experience in unix terminals.

puzzl comes with an in-built solver which can solve any puzzle configuration faster than Iron Man. :zap:

![puzzl](https://github.com/pravj/puzzl/blob/development/puzzl.gif?raw=true)

I have also written a paper\* describing technical details about the game, have a look, in case you want to.

> Implementing an intelligent version of the classical sliding-puzzle game for unix terminals using Golang's concurrency primitives
>
> http://arxiv.org/abs/1503.08345

####Installation
```go
go get github.com/pravj/puzzl
```
* Make sure that the workspace's **bin** directory is added to your **PATH**
```go
export PATH=$PATH:$GOPATH/bin
```

####Controls
* Start the game with the command *puzzl*.
* Use Arrow Keys to move the blank tile wherever you want.
* Press 'h' or 'H' to get any hint for next move.
* Press ESC key to quit the game.

####Features
* puzzl comes with an [in-built solver](#in-built-solver) that powers the automation for the game.
* puzzl gives you some hope by showing the optimal possible moves to solve any board configuration.
* puzzl helps you survive the game by giving [hints](#hints-policy) for next move.
* puzzl tracks all the user moves and accordingly generates [score](#scoring-policy) for the game.
* puzzl shows [notifications](#notification-mechanism) according to the real time game status.
* puzzl notifies that whether your last move was right or wrong.

####In-built Solver
* puzzl uses A-star algorithm to solve the game board.
* puzzl's solver is enough fuel-efficient that it can solve the hardest 3x3 puzzle in 31 moves. Exactly what the [ideal solvability condition](http://en.wikipedia.org/wiki/15_puzzle#Solvability) asks for.

####Hints Policy
* You will get a maximum of 3 hints per game session. No more cheatings. :oncoming_police_car:

####Scoring Policy
* puzzl has its own scoring system. It measures the real time game score using two parameters, one is *total played game moves (T-score)* and another is *accumulated correct score (A-score)* from all the moves.
* Whenever a user moves in a correct direction as the solver would have moved, the *A-score* increases by 1 and decreases by 1 when the user moves in a wrong direction.
* The score of game at any point of time is calculated by this function. [ score = A-score / T-score ]
* This way the maximum score of 1 would be possible in only one situation when the user traverse the game's state space in the right direction all the time.

####Notification Mechanism
* puzzl uses a combination of *goroutines* and *channels* to deliver real time notifications in the game.
* Here you can see all the [available notifications](https://github.com/pravj/puzzl/blob/development/notification/notification.go#L5-L14).

####Dependencies
* [termbox-go](https://github.com/nsf/termbox-go) - Text based graphic user interface for the game.
* [Box-drawing characters](http://en.wikipedia.org/wiki/Box-drawing_character) for drawing different sections.

---

Built with *Muzi* and *Coffee* by [Pravendra Singh](https://twitter.com/hackpravj)
