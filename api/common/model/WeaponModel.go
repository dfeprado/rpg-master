package model

type Weapon struct {
	Item Item `json:"item"`
	// Ponto de parada. Vai diminuindo conforme o jogador usa
	PA          uint    `json:"pa"`
	Damage      string  `json:"damange"`
	TwoHanded   bool    `json:"twoHanded"`
	Distance    float32 `json:"distance"`
	Effect      string  `json:"effect"`
	Size        string  `json:"size"`
	Enhancement string  `json:"enchancement"`
}
