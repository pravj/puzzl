
package solver

// TODO: init() content

import (
  _ "fmt"
  "container/heap"
  "github.com/pravj/puzzl/board"
)

type Node struct {
  parent *Node
  state board.Board

  gCost int
  hCost int
  fCost int

  index int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int {
  return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
  return pq[i].fCost < pq[j].fCost
}

func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
  pq[i].index = i
  pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
  n := len(*pq)
  item := x.(Node)
  item.index = n
  *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
  old := *pq
  n := len(old)
  item := old[n-1]
  item.index = -1 // for safety
  *pq = old[0 : n-1]
  return item
}

// What about using linked-list for path tracing?
// whatever you do, make sure to think about memory of supporting data types
type Path struct {
  nodes []Node
}

type OpenList struct {
  table map[board.Board]bool
  queue *PriorityQueue
}

type CloseList struct {
  table map[board.Board]bool
  queue *PriorityQueue
}

type Solver struct {
  openlist *OpenList
  closelist *CloseList

  goal board.Board
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

// scoring updates scores for a Node used in the progress
func scoring(node *Node, isRoot bool) {
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

  for i := 0; i < len(moves)/2; i++ {
    bTemp := b
    bTemp.Move(moves[2*i], moves[2*i+1])

    list = append(list, bTemp)
  }

  return list
}

func New(b board.Board) *Solver {
  openlist := &OpenList{}
  closelist := &CloseList{}

  // initiate the solver with open and close lists
  solver := &Solver{openlist: openlist, closelist: closelist}
  solver.openlist.table = make(map[board.Board]bool)
  solver.closelist.table = make(map[board.Board]bool)

  var opq, cpq PriorityQueue
  heap.Init(&opq)
  heap.Init(&cpq)

  solver.openlist.queue = &opq
  solver.closelist.queue = &cpq

  // Node representing the initial configuration of the board
  currentNode := Node{parent: nil, state: b}
  // updates traversal cost values for the node(root)
  scoring(&currentNode, true)

  // add initial configuration(root Node) to open list
  solver.openlist.table[currentNode.state] = true
  heap.Push(solver.openlist.queue, currentNode)

  // generate the default goal state for the process
  solver.goalState()

  return solver
}

// goalState generates the default goal state
// TODO: size as a variable(const)
func (s *Solver) goalState() {
  b := board.New()

  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      b.Rows[i].Tiles[j].Value = (3*i+j+1)%9
    }
  }

  b.BlankRow, b.BlankCol = 2, 2

  s.goal = *b
}

func (s *Solver) Solve() {
  var currentNode Node

  for (s.openlist.queue.Len() > 0) {
    // remove node index from open list
    s.openlist.table[currentNode.state] = false
    currentNode = heap.Pop(s.openlist.queue).(Node)

    // add node index to close list
    s.closelist.table[currentNode.state] = true
    heap.Push(s.closelist.queue, currentNode)

    if (currentNode.state == s.goal) {
      // TODO: termination state
      break
    } else {
      adjacents := neighbours(currentNode.state)
      for i := 0; i < len(adjacents); i++ {
        //
      }
    }
  }
}
