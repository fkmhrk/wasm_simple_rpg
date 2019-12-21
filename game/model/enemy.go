package model

type Enemy struct {
	Name    string
	HP      int
	STR     int
	DEF     int
	Special int
	XP      int
	Gold    int
	IsBoss  bool
}

func (e *Enemy) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"name": e.Name,
		"hp":   e.HP,
		"str":  e.STR,
		"def":  e.DEF,
		"sp":   e.Special,
		"xp":   e.XP,
		"gold": e.Gold,
	}
}

func NewNormalEnemy(floor int) *Enemy {
	switch floor {
	case 1:
		return makeEnemy("Enemy 0x01", 27, 10, 0, 0, 2, 3, false)
	default:
		return makeEnemy("Enemy 0x01", 27, 10, 0, 0, 2, 3, false)
	}
}

func NewBossEnemy(floor int) *Enemy {
	switch floor {
	case 1:
		return makeEnemy("Boss 0x11", 100, 35, 1, 0, 30, 25, true)
	default:
		return makeEnemy("Boss 0x11", 27, 10, 0, 0, 2, 3, true)
	}
}

func makeEnemy(name string, hp, str, def, special, xp, gold int, isBoss bool) *Enemy {
	return &Enemy{
		Name:    name,
		HP:      hp,
		STR:     str,
		DEF:     def,
		Special: special,
		XP:      xp,
		Gold:    gold,
		IsBoss:  isBoss,
	}
}
