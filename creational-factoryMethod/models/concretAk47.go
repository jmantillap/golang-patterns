package models

type Ak47 struct {
	Fusil
}

func NewAk47() IFusil {
	return &Ak47{
		Fusil: Fusil{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
