package model

type GameState struct {
	State int
	Party Party
}

func NewGameState() *GameState {
	return &GameState{
		State: 1,
	}
}

func (s *GameState) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"party": s.Party.ToJSON(),
	}
}
