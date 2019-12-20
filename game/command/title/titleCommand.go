package title

import (
	"syscall/js"

	"github.com/fkmhrk/web-simple-rpg/model"
)

func Start(state *model.GameState, args []js.Value) map[string]interface{} {
	var data string = args[1].String()
	if len(data) == 0 {
		state.State = 1
		return map[string]interface{}{
			"next_page": "move",
			"data":      state.ToJSON(),
		}
	}
	// todo restore save data
	return map[string]interface{}{
		"next_page": "move",
		"data":      state.ToJSON(),
	}
}
