package nim

import (
	"fmt"
	"strings"

	"github.com/lzambarda/gomcts"
	"github.com/lzambarda/gomcts/experiments/utils"
)

type Nim struct {
	currentPlayer    int
	board            [4][]int
	remainingMatches int
	winner           int
	alone            int
	notAlone         int
}

func NewNim(initialPlayer int) Nim {
	n := Nim{
		currentPlayer:    initialPlayer,
		remainingMatches: 16,
	}
	n.board[0] = make([]int, 1)
	n.board[1] = make([]int, 3)
	n.board[2] = make([]int, 5)
	n.board[3] = make([]int, 7)
	return n
}

func (s Nim) IsGameEnded() bool {
	return s.winner != 0 || s.remainingMatches == 0
}

func (s Nim) GetScore(playerIndex int) float64 {
	if s.winner == 0 {
		return 0.5
	}
	if playerIndex == s.winner {
		return 1
	}
	return 0
}

func (s Nim) GetLegalActions() []gomcts.Move {
	moves := make([]gomcts.Move, 0, 4*4)
	if !s.IsGameEnded() {
		for i := 0; i < len(s.board); i++ {
			for start := 0; start < len(s.board[i]); start++ {
			loop:
				for end := start; end < len(s.board[i]); end++ {
					for check := start; check < end+1; check++ {
						if s.board[i][check] != 0 {
							continue loop
						}
					}
					moves = append(moves, NimMove{row: i, start: start, end: end, player: s.currentPlayer})
				}
			}

		}
	}
	return moves
}

func (s Nim) GetCurrentPlayer() int {
	return s.currentPlayer
}

func (s Nim) GetWinner() int {
	return s.winner
}

type NimMove struct {
	row    int
	start  int
	end    int
	player int
}

func (a NimMove) ApplyTo(s gomcts.GameState) gomcts.GameState {
	state := s.(Nim)
	for i := 0; i < len(state.board); i++ {
		x := state.board[i]
		state.board[i] = make([]int, len(x))
		copy(state.board[i], x)
	}
	if state.currentPlayer != a.player {
		panic("wrong player turn")
	}
	for i := a.start; i < a.end+1; i++ {
		if state.board[a.row][i] != 0 {
			panic("illegal move")
		}
		state.board[a.row][i] = 1
		state.remainingMatches--
	}
	state.currentPlayer = -state.currentPlayer

	// Now count how many matches are alone, if the number is odd the player has
	// lost.
	alone := 0
	notAlone := 0
	for i := 0; i < len(state.board); i++ {
		for j := 0; j < len(state.board[i]); j++ {
			if state.board[i][j] != 0 {
				continue
			}
			if state.board[i][j] == 0 && (j == 0 || state.board[i][j-1] != 0) && (j == len(state.board[i])-1 || state.board[i][j+1] != 0) {
				alone++
			} else {
				notAlone++
			}
		}
	}

	// Last one to pick has lost
	if alone%2 == 1 && notAlone == 0 {
		state.winner = -state.currentPlayer
	}
	state.alone = alone
	state.notAlone = notAlone

	return state
}

func (t Nim) HumanMove(playerIndex int) gomcts.Move {
	row, err := utils.ReadNumber("row [0-3]:")
	if err != nil {
		panic(err)
	}
	start, err := utils.ReadNumber("start:")
	if err != nil {
		panic(err)
	}
	end, err := utils.ReadNumber("end:")
	if err != nil {
		panic(err)
	}
	return NimMove{row: row, start: start, end: end, player: playerIndex}
}

func (t Nim) RenderGame() {
	fmt.Println("board:")
	tk := []string{"I", " "}
	for i := 0; i < 4; i++ {
		pad := strings.Repeat(" ", 3-i)
		matches := ""
		for j := 0; j < len(t.board[i]); j++ {
			matches += tk[t.board[i][j]]
		}
		fmt.Printf("%s%s%s\n", pad, matches, pad)
	}
	fmt.Println("remaining matches:", t.remainingMatches)
}
