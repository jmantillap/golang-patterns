package models

type Fusil struct {
	name  string
	power int
}

func (f *Fusil) GetName() string {
	return f.name
}

func (f *Fusil) GetPower() int {
	return f.power
}

func (f *Fusil) SetPower(power int) {
	f.power = power
}

func (f *Fusil) SetName(name string) {
	f.name = name
}
