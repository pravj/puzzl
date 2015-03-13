// Package solver implemets auto-solving of the game
package solver

// TODO: init() content

import (
  "fmt"
  "github.com/pravj/puzzl/board"
)

type Node struct {
  parent *Node
  state board.Board

  gCost int
  hCost int
  fCost int
}

// What about using linked-list for path tracing?
// whatever you do, make sure to think about memory of supporting data types
type Path struct {
  nodes []Node
}

type OpenList struct {
  nodes map[Node]bool
}

type CloseList struct {
  nodes map[*Node]bool
}

type Solver struct {
  openlist *OpenList
  closelist *CloseList
}

// implements misplaced tile count as a heuristic scoring function
func heuristicScore(b board.Board) int {
  var score int

  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if (b.Rows[i].Tiles[j].Value != ((3*i+j+1)%9)) {
        score += 1
      }
    }
  }

  return score
}

// scoring returns scores for a Node used in the progress
func scoring(node Node, isRoot bool) {
  var g int
  if (!isRoot) {
    g = node.parent.gCost + 1
  }

  h := heuristicScore(node.state)
  f := g+h

  node.gCost, node.hCost, node.fCost = g, h, f
}

// Neighbours returns a list of board configurations
// adjacent to a given configuration
func neighbours(b board.Board) ([]board.Board) {
  var list []board.Board
  moves := b.Moves(b.BlankRow, b.BlankCol)
  //fmt.Println(moves)

  for i := 0; i < len(moves)/2; i++ {
    var bTemp board.Board
    bTemp = b
    //fmt.Println(bTemp)
    bTemp.Move(moves[2*i], moves[2*i+1])
    //fmt.Println(bTemp)
    list = append(list, bTemp)
  }

  return list
}

func New(b board.Board) *Solver {
  openlist := &OpenList{}
  closelist := &CloseList{}

  // initiate the solver with open and close lists
  solver := &Solver{openlist: openlist, closelist: closelist}
  solver.openlist.nodes = make(map[Node]bool)
  solver.closelist.nodes = make(map[*Node]bool)

  // Node representing the initial configuration of the board
  currentNode := Node{parent: nil, state: b}
  // updates traversal cost values for the node(root)
  scoring(currentNode, true)

  // add initial configuration(root Node) to open list
  solver.openlist.nodes[currentNode] = true

  fmt.Println("start", currentNode.state)
  fmt.Println(neighbours(currentNode.state))

  return solver
}
