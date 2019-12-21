package battle

import (
	"syscall/js"

	"math/rand"

	"github.com/fkmhrk/web-simple-rpg/command/result"
	"github.com/fkmhrk/web-simple-rpg/model"
)

func Next(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State == model.StateBattleStart {
		return findNextCharacter(state, args)
	}
	if state.State == model.StateBattlePlayerAttacked {
		return checkEnemyKilled(state, args)
	}
	if state.State == model.StateBattleEnemyAttacked {
		return checkPlayerKilled(state, args)
	}
	if state.State == model.StateBattleGotXP {
		return checkLevelUp(state, args)
	}
	// todo restore save data
	return map[string]interface{}{
		"next_page": "",
		"data":      state.ToJSON(),
	}
}

func Fight(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateBattleSelectCommand {
		return result.ErrInvalidState
	}
	currentCharacter := state.Party.Characters[state.Index]
	damage := currentCharacter.STR - state.Enemy.DEF*4
	damage = disperse(damage, 40)
	state.Enemy.HP -= damage
	if state.Enemy.HP < 0 {
		state.Enemy.HP = 0
	}

	state.State = model.StateBattlePlayerAttacked

	data := state.ToJSON()
	data["battle"] = map[string]interface{}{
		"attack_type": 1,
		"damage":      damage,
	}

	return map[string]interface{}{
		"next_page": "",
		"data":      data,
	}
}

func Magic(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateBattleSelectCommand {
		return result.ErrInvalidState
	}
	state.State = model.StateBattleSelectMagic
	data := state.ToJSON()
	return map[string]interface{}{
		"next_page": "",
		"data":      data,
	}
}

func HealMagic(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateBattleSelectMagic {
		return result.ErrInvalidState
	}
	currentCharacter := state.Party.Characters[state.Index]
	if currentCharacter.MP < 1 {
		return map[string]interface{}{
			"next_page": "",
			"data":      state.ToJSON(),
		}
	}
	state.State = model.StateBattleSelectHealTarget
	data := state.ToJSON()
	return map[string]interface{}{
		"next_page": "",
		"data":      data,
	}
}

func Target(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateBattleSelectHealTarget {
		return result.ErrInvalidState
	}
	var target int = args[1].Int()
	currentCharacter := state.Party.Characters[state.Index]
	targetCharacter := state.Party.Characters[target]

	currentCharacter.MP -= 1
	value := currentCharacter.Level * 16
	targetCharacter.HP += value
	if targetCharacter.HP > targetCharacter.MaxHP {
		targetCharacter.HP = targetCharacter.MaxHP
	}
	return findNextCharacter(state, args)
}

func DoCureMagic(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateBattleSelectMagic {
		return result.ErrInvalidState
	}
	currentCharacter := state.Party.Characters[state.Index]
	if currentCharacter.MP < 3 {
		return map[string]interface{}{
			"next_page": "",
			"data":      state.ToJSON(),
		}
	}
	currentCharacter.MP -= 3
	value := currentCharacter.Level * 6
	for _, ch := range state.Party.Characters {
		ch.HP += value
		if ch.HP > ch.MaxHP {
			ch.HP = ch.MaxHP
		}
	}
	return findNextCharacter(state, args)
}

func StateUp(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.State != model.StateBattleLevelUp {
		return result.ErrInvalidState
	}
	targetCharacter := state.Party.Characters[state.Index]
	var stateType int = args[1].Int()
	switch stateType {
	case 0:
		targetCharacter.MaxHP += 5
	case 1:
		targetCharacter.MaxMP += 2
	case 2:
		targetCharacter.STR += 1
	case 3:
		targetCharacter.DEF += 1
	}
	return checkLevelUp(state, args)
}

func findNextCharacter(state *model.GameState, args []js.Value) map[string]interface{} {
	for i := state.Index + 1; i < len(state.Party.Characters); i++ {
		// alive?
		ch := state.Party.Characters[i]
		if ch.HP > 0 {
			state.Index = i
			state.State = model.StateBattleSelectCommand
			data := state.ToJSON()
			data["battle"] = map[string]interface{}{
				"index": i,
			}
			return map[string]interface{}{
				"next_page": "",
				"data":      data,
			}
		}
	}
	// enemy phase
	return doEnemyAction(state, args)
}

func checkEnemyKilled(state *model.GameState, args []js.Value) map[string]interface{} {
	if state.Enemy.HP == 0 {
		state.State = model.StateBattleGotXP
		state.Index = 0
		data := state.ToJSON()
		return map[string]interface{}{
			"next_page": "",
			"data":      data,
		}
	}
	return findNextCharacter(state, args)
}

func doEnemyAction(state *model.GameState, args []js.Value) map[string]interface{} {
	// normal attack
	var target int
	for {
		target = rand.Intn(len(state.Party.Characters))
		if state.Party.Characters[target].HP > 0 {
			break
		}
	}

	targetCharacter := state.Party.Characters[target]
	damage := state.Enemy.STR - targetCharacter.DEF*4
	damage = disperse(damage, 30)
	targetCharacter.HP -= damage
	if targetCharacter.HP < 0 {
		targetCharacter.HP = 0
	}
	state.Index = target
	state.State = model.StateBattleEnemyAttacked
	data := state.ToJSON()
	data["battle"] = map[string]interface{}{
		"index":  target,
		"damage": damage,
	}
	return map[string]interface{}{
		"next_page": "",
		"data":      data,
	}
}

func checkPlayerKilled(state *model.GameState, args []js.Value) map[string]interface{} {
	targetCharacter := state.Party.Characters[state.Index]
	if targetCharacter.HP == 0 {
		state.State = model.StateBattleGotXP
		data := state.ToJSON()
		return map[string]interface{}{
			"next_page": "",
			"data":      data,
		}
	}
	state.Index = -1
	return findNextCharacter(state, args)
}

func checkLevelUp(state *model.GameState, args []js.Value) map[string]interface{} {
	for i := state.Index; i < len(state.Party.Characters); i++ {
		targetCharacter := state.Party.Characters[i]
		if targetCharacter.HP == 0 {
			continue
		}
		targetCharacter.XP += state.Enemy.XP
		state.Index = i
		if targetCharacter.XP >= targetCharacter.Next &&
			targetCharacter.Level < 99 {
			// Level UP
			targetCharacter.Level++
			targetCharacter.Next += (targetCharacter.Level * 10)
			targetCharacter.MaxHP += 6
			targetCharacter.MaxMP += 2
			targetCharacter.STR += 1
			state.State = model.StateBattleLevelUp
			data := state.ToJSON()
			data["battle"] = map[string]interface{}{
				"index": i,
			}
			return map[string]interface{}{
				"next_page": "",
				"data":      data,
			}
		}
	}
	state.State = model.StateMoveMain
	return map[string]interface{}{
		"next_page": "move",
		"data":      state.ToJSON(),
	}
}
