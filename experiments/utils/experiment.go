package utils

import "github.com/lzambarda/gomcts"

type Experiment interface {
	HumanMove(playerIndex int) gomcts.Move
	RenderGame()
}
