package model

type Party struct {
}

func (p *Party) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"characters": nil,
	}
}
