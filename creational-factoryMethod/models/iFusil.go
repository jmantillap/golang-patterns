package models

type IFusil interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

func NewFusil(fusilType string) (IFusil, error) {
	return nil, nil
}
