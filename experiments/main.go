package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lzambarda/gomcts"
	"github.com/lzambarda/gomcts/experiments/tictactoe"
	"github.com/lzambarda/gomcts/experiments/utils"
)

func main() {
	rand.Seed(time.Now().Unix())
	var state gomcts.GameState
	state = tictactoe.NewTicTacToe(1)
	// state = nim.NewNim(1)
	// state = connectfour.NewConnectFour(1)
	againstHuman := true
	humanIndex := -1
	var move gomcts.Move
	for !state.IsGameEnded() {
		if againstHuman && state.GetCurrentPlayer() == humanIndex {
			state.(utils.Experiment).RenderGame()
			fmt.Println("Make your move:")
			move = state.(utils.Experiment).HumanMove(state.GetCurrentPlayer())
		} else {
			if !againstHuman {
				state.(utils.Experiment).RenderGame()
			}
			move = gomcts.Search(state, 1000, 1.4, state.GetCurrentPlayer())
		}
		state = move.ApplyTo(state)
	}
	state.(utils.Experiment).RenderGame()
	winner := state.GetWinner()
	if winner == 0 {
		fmt.Println("Draw")
		return
	}
	if !againstHuman {
		fmt.Println("Winner:", winner)
		return
	}
	if winner == humanIndex {
		fmt.Println("You win")
		return
	}
	fmt.Println("CPU wins")
}
