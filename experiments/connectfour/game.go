package connectfour

import (
	"fmt"

	"github.com/lzambarda/gomcts"
	"github.com/lzambarda/gomcts/experiments/utils"
)

type ConnectFour struct {
	currentPlayer int
	board         [6][7]int
	usedCells     int
	winner        int
}

func NewConnectFour(initialPlayer int) ConnectFour {
	return ConnectFour{
		currentPlayer: initialPlayer,
	}
}

func (s ConnectFour) IsGameEnded() bool {
	return s.winner != 0 || s.usedCells == 42
}

func (s ConnectFour) GetScore(playerIndex int) float64 {
	if s.winner == 0 {
		return 0.5
	}
	if playerIndex == s.winner {
		return 1
	}
	return 0
}

func (s ConnectFour) GetLegalActions() []gomcts.Move {
	moves := make([]gomcts.Move, 0, 42-s.usedCells)
	if !s.IsGameEnded() {
		for x := 0; x < 7; x++ {
			for y := 0; y < 6; y++ {
				if s.board[y][x] == 0 {
					moves = append(moves, ConnectFourMove{x: x, y: y, player: s.currentPlayer})
					break
				}
			}
		}
	}
	return moves
}

func (s ConnectFour) GetCurrentPlayer() int {
	return s.currentPlayer
}

func (s ConnectFour) GetWinner() int {
	return s.winner
}

type ConnectFourMove struct {
	x      int
	y      int
	player int
}

func (a ConnectFourMove) ApplyTo(s gomcts.GameState) gomcts.GameState {
	state := s.(ConnectFour)
	if state.currentPlayer != a.player {
		panic("wrong player turn")
	}
	if state.board[a.y][a.x] != 0 {
		panic("illegal move, cell already used")
	}

	state.board[a.y][a.x] = a.player
	state.currentPlayer *= -1
	state.usedCells++

	// Check horizontal, vertical and diagonals starting from last play
	// Horizontal
	xmin := a.x - 3
	if xmin < 0 {
		xmin = 0
	}
	xmax := a.x + 3
	if xmax > 6 {
		xmax = 6
	}
	count := 0
	for x := xmin; x <= xmax; x++ {
		if state.board[a.y][x] != a.player {
			count = 0
		} else {
			count++
			if count == 4 {
				state.winner = a.player
				return state
			}
		}
	}
	// Vertical
	ymin := a.y - 3
	if ymin < 0 {
		ymin = 0
	}
	ymax := a.y + 3
	if ymax > 5 {
		ymax = 5
	}
	count = 0
	for y := ymin; y <= ymax; y++ {
		if state.board[y][a.x] != a.player {
			count = 0
		} else {
			count++
			if count == 4 {
				state.winner = a.player
				return state
			}
		}
	}
	// Diagonal 1
	y := ymin
	x := xmin
	for {
		if y > ymax || x > xmax {
			break
		}
		if state.board[y][x] != a.player {
			count = 0
		} else {
			count++
			if count == 4 {
				state.winner = a.player
				return state
			}
		}
		x++
		y++
	}
	// Diagonal 2
	y = ymin
	x = xmax
	for {
		if y > ymax || x < xmin {
			break
		}
		if state.board[y][x] != a.player {
			count = 0
		} else {
			count++
			if count == 4 {
				state.winner = a.player
				return state
			}
		}
		x--
		y++
	}

	return state
}

func (e ConnectFour) HumanMove(playerIndex int) gomcts.Move {
	x, err := utils.ReadNumber("column [0-7]:")
	if err != nil {
		panic(err)
	}
	y := 0
	for y < 6 {
		if e.board[y][x] == 0 {
			break
		}
		y++
	}
	if y == 6 {
		panic("column is full")
	}
	return ConnectFourMove{x: x, y: y, player: playerIndex}
}

func (e ConnectFour) RenderGame() {
	tk := []string{"O", " ", "X"}

	fmt.Println("board:")
	for y := 5; y >= 0; y-- {
		if y == 5 {
			fmt.Println("╔═══════╗")
		}
		fmt.Printf("║%s%s%s%s%s%s%s║\n", tk[1+e.board[y][0]], tk[1+e.board[y][1]], tk[1+e.board[y][2]], tk[1+e.board[y][3]], tk[1+e.board[y][4]], tk[1+e.board[y][5]], tk[1+e.board[y][6]])
		if y == 0 {
			fmt.Println("╚═══════╝")
		}
	}
}
