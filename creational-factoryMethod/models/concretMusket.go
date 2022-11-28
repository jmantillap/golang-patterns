package models

type Musket struct {
	Fusil
}

func NewMusket() IFusil {
	return &Musket{
		Fusil: Fusil{
			name:  "Musket gun",
			power: 4,
		},
	}
}
