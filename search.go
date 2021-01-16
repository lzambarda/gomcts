package gomcts

// MCTS on a tree starting with a given state.
// This will return the best next move found from the point of view of
// playerIndex.
// Simulations is the number of game simulations to run across the tree.
// c is the exploration bias factor for UCT.
func Search(state GameState, simulations int, c float64, playerIndex int) Move {
	root := newNode(nil, state, nil)
	for i := 0; i < simulations; i++ {
		child := root.selectNode(c)
		result := child.simulate(playerIndex)
		child.backpropagate(result)
	}
	return root.getBestUtcChild(0.0).move
}
