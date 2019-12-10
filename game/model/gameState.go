package model

type GameState struct {
	State int
}

func NewGameState() *GameState {
	return &GameState{
		State: 1,
	}
}
