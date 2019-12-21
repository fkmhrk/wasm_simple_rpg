package model

type GameState struct {
	State    int
	Party    *Party
	TmpParty *Party
	Enemy    *Enemy
	Floor    int
	Gold     int
	Index    int
}

func NewGameState() *GameState {
	return &GameState{
		State: 1,
		Party: newParty(),
		Enemy: NewNormalEnemy(1),
		Floor: 1,
		Gold:  10,
		Index: 0,
	}
}

func (s *GameState) PrepareForBattle() {
	s.TmpParty = s.Party.Copy()
	s.Index = -1
}

func (s *GameState) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"state": s.State,
		"party": s.Party.ToJSON(),
		"floor": s.Floor,
		"gold":  s.Gold,
		"enemy": s.Enemy.ToJSON(),
	}
}

const (
	StateMoveMain                 = 1
	StateMoveStateSelectCharacter = 2
	StateMoveStatusShow           = 3

	StateBattleStart          = 100
	StateBattleSelectCommand  = 101
	StateBattlePlayerAttacked = 102
	StateBattleGotXP          = 103
	StateBattleEnemyAttacked  = 110
	StateBattleLevelUp        = 120
)
