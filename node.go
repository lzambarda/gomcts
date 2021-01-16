package gomcts

import (
	"math"
	"math/rand"
)

type node struct {
	parent *node
	// A fully expanded node will have chldren length equal to the initial
	// length of unexpandedMoves, and the latter will be empty
	children []*node
	// The game move which created this node
	move Move
	// The game state associated with this node
	state GameState
	// The game moves which are left to be explored
	unexpandedMoves []Move
	// The total accrued score for this node
	score float64
	// The total simulations which backpropagated through this node
	visits int
}

func newNode(parent *node, state GameState, causingMove Move) node {
	return node{
		parent:          parent,
		state:           state,
		move:            causingMove,
		children:        []*node{},
		unexpandedMoves: state.GetLegalActions(),
	}
}

// Get the Upper Confidence Bound applied to trees of the current node.
func (n *node) getUtc(c float64) float64 {
	return (float64(n.score) / float64(n.visits)) + c*math.Sqrt(math.Log(float64(n.parent.visits))/float64(n.visits))
}

func (n *node) isFullyExpanded() bool {
	return len(n.unexpandedMoves) == 0
}

func (n *node) isLeaf() bool {
	return n.state.IsGameEnded()
}

func (n *node) selectNode(c float64) *node {
	for !n.isLeaf() {
		if !n.isFullyExpanded() {
			return n.expand()
		}
		n = n.getBestUtcChild(c)
	}
	return n
}

func (n *node) getBestUtcChild(c float64) *node {
	var best *node
	max := -math.MaxFloat64
	for _, child := range n.children {
		current := child.getUtc(c)
		if current > max {
			max = current
			best = child
		}
	}
	return best
}

func (n *node) expand() *node {
	// move := n.unexpandedMoves[0]
	// n.unexpandedMoves = n.unexpandedMoves[1:]
	i := rand.Intn(len(n.unexpandedMoves))
	move := n.unexpandedMoves[i]
	n.unexpandedMoves = append(n.unexpandedMoves[:i], n.unexpandedMoves[i+1:]...)
	expandedChild := newNode(n, move.ApplyTo(n.state), move)
	n.children = append(n.children, &expandedChild)
	return &expandedChild
}

func (n *node) simulate(playerIndex int) float64 {
	currentState := n.state
	for !currentState.IsGameEnded() {
		moves := currentState.GetLegalActions()
		currentState = moves[rand.Intn(len(moves))].ApplyTo(currentState)
	}
	return currentState.GetScore(playerIndex)
}

func (n *node) backpropagate(score float64) {
	for n.parent != nil {
		n.score += score
		n.visits++
		n = n.parent
	}
	n.visits++
}
