package command

import (
	"syscall/js"

	"github.com/fkmhrk/web-simple-rpg/command/move"
	"github.com/fkmhrk/web-simple-rpg/command/title"
	"github.com/fkmhrk/web-simple-rpg/model"
)

type CommandFunc func(state *model.GameState, args []js.Value) map[string]interface{}

var commands = make(map[string]CommandFunc)

func init() {
	commands["start"] = title.Start
	commands["menu"] = move.Menu
	commands["status"] = move.Status
	commands["back"] = Back
	commands["next_floor"] = move.NextFloor
	commands["select_character"] = SelectCharacter
}

func Exec(state *model.GameState, args []js.Value) map[string]interface{} {
	c := args[0].String()
	f, ok := commands[c]
	if ok {
		return f(state, args)
	}
	return map[string]interface{}{
		"error_code": "UNKNOWN_CODE",
	}
}

func Back(state *model.GameState, args []js.Value) map[string]interface{} {
	switch state.State {
	case model.StateMoveStateSelectCharacter:
		state.State = model.StateMoveMain
	case model.StateMoveStatusShow:
		state.State = model.StateMoveStateSelectCharacter
	}
	return map[string]interface{}{
		"next_page": "",
		"data":      state.ToJSON(),
	}
}

func SelectCharacter(state *model.GameState, args []js.Value) map[string]interface{} {
	switch state.State {
	case model.StateMoveStateSelectCharacter:
		return move.SelectStatusCharacter(state, args)
	}
	return map[string]interface{}{
		"error_code": "UNKNOWN_CODE",
	}
}
