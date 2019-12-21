package model

type GameState struct {
	State int
	Party *Party
	Floor int
	Gold  int
}

func NewGameState() *GameState {
	return &GameState{
		State: 1,
		Party: newParty(),
		Floor: 1,
		Gold:  10,
	}
}

func (s *GameState) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"state": s.State,
		"party": s.Party.ToJSON(),
		"floor": s.Floor,
		"gold":  s.Gold,
	}
}

const (
	StateMoveMain                 = 1
	StateMoveStateSelectCharacter = 2
	StateMoveStatusShow           = 3
)
