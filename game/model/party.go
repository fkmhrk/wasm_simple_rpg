package model

type Party struct {
	Characters Characters
}

type Character struct {
	Name  string
	HP    int
	MaxHP int
	MP    int
	MaxMP int
	Level int
	XP    int
	Next  int
}

type Characters []*Character

func newParty() *Party {
	characters := make([]*Character, 0, 5)
	characters = append(characters, newCharacter("Fig"))
	characters = append(characters, newCharacter("Pri"))
	characters = append(characters, newCharacter("Thi"))
	characters = append(characters, newCharacter("Mag"))
	characters = append(characters, newCharacter("Ran"))
	return &Party{
		Characters: characters,
	}
}

func newCharacter(name string) *Character {
	return &Character{
		Name:  name,
		HP:    30,
		MaxHP: 30,
		MP:    3,
		MaxMP: 3,
		Level: 1,
		XP:    0,
		Next:  10,
	}
}

func (p *Party) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"characters": p.Characters.ToJSON(),
	}
}

func (c Characters) ToJSON() []interface{} {
	list := make([]interface{}, 0, len(c))
	for _, character := range c {
		list = append(list, character.ToJSON())
	}
	return list
}

func (c *Character) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"name":    c.Name,
		"hp":      c.HP,
		"max_hp":  c.MaxHP,
		"mp":      c.MP,
		"max_mp":  c.MaxMP,
		"level":   c.Level,
		"xp":      c.XP,
		"next_xp": c.Next,
	}
}
