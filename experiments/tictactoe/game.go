package tictactoe

import (
	"fmt"

	"github.com/lzambarda/gomcts"
	"github.com/lzambarda/gomcts/experiments/utils"
)

type TicTacToe struct {
	currentPlayer int
	board         [3][3]int
	usedCells     int
	winner        int
}

func NewTicTacToe(initialPlayer int) TicTacToe {
	return TicTacToe{
		currentPlayer: initialPlayer,
	}
}

func (s TicTacToe) IsGameEnded() bool {
	return s.winner != 0 || s.usedCells == 9
}

func (s TicTacToe) GetScore(playerIndex int) float64 {
	if s.winner == 0 {
		return 0.5
	}
	if playerIndex == s.winner {
		return 1
	}
	return 0
}

func (s TicTacToe) GetLegalActions() []gomcts.Move {
	moves := make([]gomcts.Move, 0, 9-s.usedCells)
	if !s.IsGameEnded() {
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				if s.board[y][x] == 0 {
					moves = append(moves, TicTacToeMove{x: x, y: y, player: s.currentPlayer})
				}
			}
		}
	}
	return moves
}

func (s TicTacToe) GetCurrentPlayer() int {
	return s.currentPlayer
}

func (s TicTacToe) GetWinner() int {
	return s.winner
}

type TicTacToeMove struct {
	x      int
	y      int
	player int
}

func (a TicTacToeMove) ApplyTo(s gomcts.GameState) gomcts.GameState {
	state := s.(TicTacToe)
	if state.currentPlayer != a.player {
		panic("wrong player turn")
	}
	if state.board[a.y][a.x] != 0 {
		panic("illegal move, cell already used")
	}

	state.board[a.y][a.x] = a.player
	state.currentPlayer *= -1
	state.usedCells++

	// Check filled rows and columns
	for i := 0; i < 3; i++ {
		if state.board[i][0] != 0 && state.board[i][0] == state.board[i][1] && state.board[i][0] == state.board[i][2] {
			state.winner = state.board[i][0]
			return state
		}
		if state.board[0][i] != 0 && state.board[0][i] == state.board[1][i] && state.board[0][i] == state.board[2][i] {
			state.winner = state.board[0][i]
			return state
		}
	}
	// Check diagonals
	if state.board[0][0] != 0 && state.board[0][0] == state.board[1][1] && state.board[0][0] == state.board[2][2] {
		state.winner = state.board[0][0]
		return state
	}
	if state.board[0][2] != 0 && state.board[0][2] == state.board[1][1] && state.board[0][2] == state.board[2][0] {
		state.winner = state.board[0][2]
		return state
	}

	return state
}

func (t TicTacToe) HumanMove(playerIndex int) gomcts.Move {
	y, err := utils.ReadNumber("row [0-2]:")
	if err != nil {
		panic(err)
	}
	x, err := utils.ReadNumber("column [0-2]:")
	if err != nil {
		panic(err)
	}
	return TicTacToeMove{x: x, y: y, player: playerIndex}
}

func (t TicTacToe) RenderGame() {
	tk := []string{"O", " ", "X"}
	fmt.Printf("\nboard:\n╔═══╗\n║%s%s%s║\n║%s%s%s║\n║%s%s%s║\n╚═══╝\n",
		tk[1+t.board[0][0]], tk[1+t.board[0][1]], tk[1+t.board[0][2]],
		tk[1+t.board[1][0]], tk[1+t.board[1][1]], tk[1+t.board[1][2]],
		tk[1+t.board[2][0]], tk[1+t.board[2][1]], tk[1+t.board[2][2]])
}
