package move

import (
	"syscall/js"

	"github.com/fkmhrk/web-simple-rpg/command/result"
	"github.com/fkmhrk/web-simple-rpg/model"
)

func Menu(state *model.GameState, args []js.Value) map[string]interface{} {
	var data int = args[1].Int()
	_ = data
	// todo restore save data
	return map[string]interface{}{
		"next_page": "",
		"data":      state.ToJSON(),
	}
}

func Shop(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateMoveMain {
		return result.ErrInvalidState
	}
	state.State = model.StateMoveShopItemList
	return result.WithState(state)
}

func Buy(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateMoveShopItemList {
		return result.ErrInvalidState
	}
	var index int = args[1].Int()
	if index == 0 {
		// herb
		if state.Gold < 2 {
			return result.WithState(state)
		}
		state.Index = 0
		state.State = model.StateMoveShopTarget
		return result.WithState(state)
	} else if index == 1 {
		// potion
		if state.Gold < 5 {
			return result.WithState(state)
		}
		state.Index = 1
		state.State = model.StateMoveShopTarget
		return result.WithState(state)
	}
	return result.WithState(state)
}

func ShopTarget(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateMoveShopTarget {
		return result.ErrInvalidState
	}
	var target int = args[1].Int()
	targetCharacter := state.Party.Characters[target]
	if state.Index == 0 {
		state.Gold -= 2
		targetCharacter.HP = targetCharacter.MaxHP
	} else if state.Index == 1 {
		state.Gold -= 5
		targetCharacter.MP = targetCharacter.MaxMP
	}
	state.State = model.StateMoveShopItemList
	return result.WithState(state)
}

func Status(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateMoveMain {
		return map[string]interface{}{
			"error_code": "INVALID_STATE",
		}
	}
	state.State = model.StateMoveStateSelectCharacter
	return map[string]interface{}{
		"next_page": "",
		"data":      state.ToJSON(),
	}
}

func NextFloor(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateMoveMain {
		return map[string]interface{}{
			"error_code": "INVALID_STATE",
		}
	}
	// check max floor
	if state.Floor == 10 {
		return map[string]interface{}{
			"next_page": "clear",
			"data":      state.ToJSON(),
		}
	}
	if state.Floor == state.MaxFloor {
		// boss battle
		// prepare
		state.PrepareForBattle()
		// select enemy
		state.Enemy = model.NewBossEnemy(state.Floor)

		state.State = model.StateBattleStart

		return map[string]interface{}{
			"next_page": "battle",
			"data":      state.ToJSON(),
		}
	}
	state.Floor++
	return map[string]interface{}{
		"next_page": "",
		"data":      state.ToJSON(),
	}
}

func Find(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateMoveMain {
		return map[string]interface{}{
			"error_code": "INVALID_STATE",
		}
	}
	// prepare
	state.PrepareForBattle()
	// select enemy
	state.Enemy = model.NewNormalEnemy(state.Floor)

	state.State = model.StateBattleStart

	return map[string]interface{}{
		"next_page": "battle",
		"data":      state.ToJSON(),
	}
}

func SelectStatusCharacter(state *model.GameState, args []js.Value) map[string]interface{} {
	var index int = args[1].Int()
	state.State = model.StateMoveStatusShow
	data := state.ToJSON()
	data["selected_index"] = index
	return map[string]interface{}{
		"next_page": "",
		"data":      data,
	}
}
