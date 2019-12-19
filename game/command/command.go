package command

import (
	"syscall/js"

	"github.com/fkmhrk/web-simple-rpg/command/title"
	"github.com/fkmhrk/web-simple-rpg/model"
)

type CommandFunc func(state *model.GameState, args []js.Value) map[string]interface{}

var commands = make(map[string]CommandFunc)

func init() {
	commands["start"] = title.Start
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
