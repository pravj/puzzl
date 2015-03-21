puzzl
=====

> An intelligent version of the sliding-puzzle game for your terminal built in golang

[![GoDoc](https://godoc.org/github.com/pravj/puzzl?status.svg)](http://godoc.org/github.com/pravj/puzzl)

puzzl is an intelligent implementation of the classical sliding-puzzle game.

It's built on top of Golang's *concurrency primitives*. It uses *goroutines* and *channels* for inter process communications to provide real time notification experience in unix terminals.

puzzl comes with an in-built solver which can solve any puzzle configuration faster than Iron Man. :zap:

![puzzl](https://github.com/pravj/puzzl/blob/development/puzzl.gif?raw=true)

####Installtion
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

---

Built with *Muzi* and *Coffee* by [Pravendra Singh](https://twitter.com/hackpravj)
