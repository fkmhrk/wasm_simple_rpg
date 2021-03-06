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
	STR   int
	DEF   int
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
		STR:   3,
		DEF:   0,
	}
}

func (p *Party) Copy() *Party {
	characters := make([]*Character, 0, len(p.Characters))
	for _, c := range p.Characters {
		characters = append(characters, c.Copy())
	}
	return &Party{
		Characters: characters,
	}
}

func (p *Party) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"characters": p.Characters.ToJSON(),
	}
}

func (p *Party) Restore(data map[string]interface{}) {
	characters := data["characters"].([]interface{})
	list := make([]*Character, 0, len(characters))
	for _, characterData := range characters {
		ch := &Character{}
		ch.Restore(characterData)
		list = append(list, ch)
	}
	p.Characters = list
}

func (c Characters) ToJSON() []interface{} {
	list := make([]interface{}, 0, len(c))
	for _, character := range c {
		list = append(list, character.ToJSON())
	}
	return list
}

func (c *Character) AddHP(value int) {
	c.HP += value
	if c.HP < 0 {
		c.HP = 0
	} else if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
}

func (c *Character) Copy() *Character {
	return &Character{
		Name:  c.Name,
		HP:    c.HP,
		MaxHP: c.MaxHP,
		MP:    c.MP,
		MaxMP: c.MaxMP,
		Level: c.Level,
		XP:    c.XP,
		Next:  c.Next,
		STR:   c.STR,
		DEF:   c.DEF,
	}
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
		"str":     c.STR,
		"def":     c.DEF,
	}
}

func (c *Character) Restore(data interface{}) {
	m := data.(map[string]interface{})

	name := m["name"].(string)
	hp := m["hp"].(float64)
	maxHP := m["max_hp"].(float64)
	mp := m["mp"].(float64)
	maxMP := m["max_mp"].(float64)
	level := m["level"].(float64)
	xp := m["xp"].(float64)
	next := m["next_xp"].(float64)
	str := m["str"].(float64)
	def := m["def"].(float64)

	c.Name = name
	c.HP = int(hp)
	c.MaxHP = int(maxHP)
	c.MP = int(mp)
	c.MaxMP = int(maxMP)
	c.Level = int(level)
	c.XP = int(xp)
	c.Next = int(next)
	c.STR = int(str)
	c.DEF = int(def)

}
