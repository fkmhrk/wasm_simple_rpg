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
	case 2:
		return makeEnemy("Enemy 0x02", 90, 25, 0, 0, 6, 7, false)
	case 3:
		return makeEnemy("Enemy 0x03", 170, 40, 0, 0, 19, 11, false)
	case 4:
		return makeEnemy("Enemy 0x04", 225, 45, 1, 0, 37, 15, false)
	case 5:
		return makeEnemy("Enemy 0x05", 310, 60, 1, 0, 64, 19, false)
	case 6:
		return makeEnemy("Enemy 0x06", 380, 72, 1, 0, 80, 23, false)
	case 7:
		return makeEnemy("Enemy 0x07", 410, 86, 2, 0, 93, 27, false)
	case 8:
		return makeEnemy("Enemy 0x08", 440, 99, 2, 0, 102, 31, false)
	case 9:
		return makeEnemy("Enemy 0x09", 480, 110, 2, 0, 114, 35, false)
	case 10:
		return makeEnemy("Enemy 0x0A", 650, 130, 3, 0, 138, 44, false)
	default:
		return makeEnemy("Enemy 0x01", 27, 10, 0, 0, 2, 3, false)
	}
}

func NewBossEnemy(floor int) *Enemy {
	switch floor {
	case 1:
		return makeEnemy("Boss 0x11", 100, 35, 0, 0, 30, 25, true)
	case 2:
		return makeEnemy("Boss 0x12", 800, 40, 0, 0, 90, 40, true)
	case 3:
		return makeEnemy("Boss 0x13", 700, 80, 0, 0, 120, 55, true)
	case 4:
		return makeEnemy("Boss 0x14", 1000, 100, 1, 0, 150, 70, true)
	case 5:
		return makeEnemy("Boss 0x15", 1250, 130, 0, 0, 183, 95, true)
	case 6:
		return makeEnemy("Boss 0x16", 1400, 145, 1, 0, 214, 110, true)
	case 7:
		return makeEnemy("Boss 0x17", 1700, 170, 1, 0, 248, 125, true)
	case 8:
		return makeEnemy("Boss 0x18", 2000, 200, 2, 0, 272, 150, true)
	case 9:
		return makeEnemy("Boss 0x19", 2550, 220, 2, 0, 310, 180, true)
	case 10:
		return makeEnemy("Final Boss", 4500, 350, 2, 0, 0, 0, true)
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
