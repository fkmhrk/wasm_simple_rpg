package result

import "github.com/fkmhrk/web-simple-rpg/model"

var (
	ErrInvalidState = map[string]interface{}{
		"error_code": "INVALID_STATE",
	}
)

func WithState(state *model.GameState) map[string]interface{} {
	return map[string]interface{}{
		"next_page": "",
		"data":      state.ToJSON(),
	}
}
