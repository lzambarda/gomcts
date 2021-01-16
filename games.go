package gomcts

type GameState interface {
	GetCurrentPlayer() int
	GetLegalActions() []Move
	IsGameEnded() bool
	GetScore(playerIndex int) float64
	GetWinner() int
}

type Move interface {
	ApplyTo(GameState) GameState
}
