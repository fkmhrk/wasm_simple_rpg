package move

import (
	"syscall/js"

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
