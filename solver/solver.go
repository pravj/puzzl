// Package solver implements auto-solving for any configuration
package solver

// TODO: init() content
// I smell it using some extra memory and that's not legit my friend.

import (
  //"fmt"
  //"reflect"
  "container/heap"
  "container/list"
  "github.com/pravj/puzzl/board"
  //"github.com/pravj/puzzl/notification"
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

type OpenList struct {
  nodeTable map[board.Board]Node
  table map[board.Board]bool

  queue *PriorityQueue
}

type CloseList struct {
  table map[board.Board]bool
}

type Solver struct {
  openlist *OpenList
  closelist *CloseList

  relation map[board.Board]board.Board
  Path *list.List
  Moves int

  Goal board.Board

  Solved bool
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

func New(b *board.Board) *Solver {
  openlist := &OpenList{}
  closelist := &CloseList{}

  // initiate the solver with open and close lists
  solver := &Solver{openlist: openlist, closelist: closelist}

  // initiate traversal lists
  solver.openlist.nodeTable = make(map[board.Board]Node)
  solver.openlist.table = make(map[board.Board]bool)
  solver.closelist.table = make(map[board.Board]bool)

  // initiate parent-child relationship and path
  solver.relation = make(map[board.Board]board.Board)
  solver.Path = list.New()

  var opq PriorityQueue
  heap.Init(&opq)

  solver.openlist.queue = &opq

  // Node representing the initial configuration of the board
  currentNode := &Node{parent: nil, state: *b}
  // updates traversal cost values for the node(root)
  scoring(currentNode, true)

  // add initial configuration(root Node) to open list
  solver.openlist.nodeTable[currentNode.state] = *currentNode
  solver.openlist.table[currentNode.state] = true
  heap.Push(solver.openlist.queue, *currentNode)

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

  s.Goal = *b
}

func (s *Solver) Solve() {
  var currentNode Node
  var count int
  var start board.Board

  for (s.openlist.queue.Len() > 0) {
    // returns the Node having lowest f-cost value(uses min-priority queue)
    currentNode = heap.Pop(s.openlist.queue).(Node)

    if (count == 0) {
      start = currentNode.state
    }

    // goal found, generating path from start to goal state
    if (currentNode.state == s.Goal) {
      state := s.Goal
      for (s.relation[state] != start) {
        state = s.relation[state]
        s.Path.PushFront(state)
      }
      s.Path.PushBack(s.Goal)

      s.Moves = s.Path.Len()
      s.Solved = true

      break
    }

    // shifts low-cost node from open list to close list
    delete(s.openlist.table, currentNode.state)
    delete(s.openlist.nodeTable, currentNode.state)
    // add low-cost node to close list
    s.closelist.table[currentNode.state] = true

    // nodes adjacent to the current node
    adjacents := neighbours(currentNode.state)

    for i := 0; i < len(adjacents); i++ {
      // adjacent node is in close list
      if (s.closelist.table[adjacents[i]]) {
        continue
      }

      // adjacent node either unavailable in open list or can be improved
      adjacentNode := s.openlist.nodeTable[adjacents[i]]
      if ((!s.openlist.table[adjacents[i]]) || (currentNode.gCost + 1 < adjacentNode.gCost)) {
        adjacentNode.gCost = currentNode.gCost + 1
        adjacentNode.hCost = heuristicScore(adjacentNode.state)
        adjacentNode.fCost = adjacentNode.gCost + adjacentNode.hCost
        adjacentNode.state = adjacents[i]

        // adjacent node is not in open list
        if (!s.openlist.table[adjacentNode.state]) {
          node := &Node{parent: &currentNode, state: adjacentNode.state}
          scoring(node, false)

          s.openlist.table[adjacents[i]] = true
          s.openlist.nodeTable[adjacents[i]] = *node

          s.relation[adjacents[i]] = currentNode.state

          heap.Push(s.openlist.queue, *node)
        }
      }
    }

    count += 1
  }
}
