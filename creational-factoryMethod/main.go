package main

import (
	"fmt"
	. "jmantillap/factorymethod/handlers"
	. "jmantillap/factorymethod/models"
)

func main() {
	factory, _ := NewFactory()
	ak47, _ := factory.GetFusil("ak47")
	musket, _ := factory.GetFusil("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IFusil) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
