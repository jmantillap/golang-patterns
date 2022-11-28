package handlers

import (
	"fmt"
	. "jmantillap/factorymethod/models"
)

type Factory struct {
}

func NewFactory() (*Factory, error) {
	return &Factory{}, nil
}

func (f *Factory) GetFusil(gunType string) (IFusil, error) {

	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")

}
