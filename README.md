# Yet another Monte Carlo Tree Search algorithm

Yup, another one!

## Installation

```
go get github.com/lzambarda/gomcts
```

## Usage

MCTS is applied in the context of a game. The interfaces `GameState` and `Move`
allow you to define any context you wish.

### Demo

I have implemented a couple
experiments with simple and naive implementations of:

- [Tic-tac-toe](https://en.wikipedia.org/wiki/Tic-tac-toe)
- [Connect Four](https://en.wikipedia.org/wiki/Connect_Four)
- [Nim](https://en.wikipedia.org/wiki/Nim)

Try them out with:

```
go run experiments/main.go
```

Currently the game used is hardcoded.
All but Nim work fine, I suspect this is because recognising Nim endgame
requires predicting many steps ahead. Not sure.

### Search

The move the CPU picks is defined by:

```
gomcts.Search(gameState, 1000, 1.4, gameState.GetCurrentPlayer())
```

This will return the best next move from the point of view of the provided
player index (1 or -1 for now).

## References

[Bandit Based Monte-Carlo Planning by Levente Kocsis and Csaba Szepesvári (2006)](http://ggp.stanford.edu/readings/uct.pdf)
