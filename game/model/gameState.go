package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type GameState struct {
	State    int
	Party    *Party
	TmpParty *Party
	Enemy    *Enemy
	Floor    int
	MaxFloor int
	Gold     int
	Index    int
}

func NewGameState() *GameState {
	return &GameState{
		State:    1,
		Party:    newParty(),
		Enemy:    NewNormalEnemy(1),
		Floor:    1,
		MaxFloor: 1,
		Gold:     10,
		Index:    0,
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

func (s *GameState) MakeSaveString() (string, string) {
	data := map[string]interface{}{
		"party":     s.Party.ToJSON(),
		"floor":     s.Floor,
		"max_floor": s.MaxFloor,
		"gold":      s.Gold,
	}
	out, err := json.Marshal(data)
	if err != nil {
		return "", ""
	}
	encrypted, iv, err := Encrypt(out)
	if err != nil {
		fmt.Printf("Failed to encrypt: %s", err)
		return "", ""
	}
	return base64.StdEncoding.EncodeToString(encrypted), iv
}

func (s *GameState) Restore(data, iv string) {
	encrypted, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Printf("Failed to decode\n")
		return
	}
	rawJSON, err := Decrypt(encrypted, iv)
	if err != nil {
		fmt.Printf("Failed to decrypt\n")
		return
	}
	var savedData map[string]interface{}
	err = json.Unmarshal(rawJSON, &savedData)
	if err != nil {
		fmt.Printf("Failed to unmarshal\n")
		return
	}
	floor := savedData["floor"].(float64)
	maxFloor := savedData["max_floor"].(float64)
	gold := savedData["gold"].(float64)
	party := savedData["party"].(map[string]interface{})

	s.Floor = int(floor)
	s.MaxFloor = int(maxFloor)
	s.Gold = int(gold)
	s.Party.Restore(party)
}

const (
	StateMoveMain                 = 1
	StateMoveStateSelectCharacter = 2
	StateMoveStatusShow           = 3
	StateMoveShopItemList         = 4
	StateMoveShopTarget           = 5

	StateBattleStart            = 100
	StateBattleSelectCommand    = 101
	StateBattlePlayerAttacked   = 102
	StateBattleGotXP            = 103
	StateBattleSelectMagic      = 104
	StateBattleSelectHealTarget = 105
	StateBattlePlayerKilled     = 106
	StateBattleAllKilled        = 107
	StateBattleEnemyAttacked    = 110
	StateBattleLevelUp          = 120
)
