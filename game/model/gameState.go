package model

type GameState struct {
	State int
	Party *Party
}

func NewGameState() *GameState {
	return &GameState{
		State: 1,
		Party: newParty(),
	}
}

func (s *GameState) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"state": s.State,
		"party": s.Party.ToJSON(),
	}
}

const (
	StateMoveMain                 = 1
	StateMoveStateSelectCharacter = 2
	StateMoveStatusShow           = 3
)
