package main

import (
	"fmt"
	"pattern/factory/configuration"
	"pattern/factory/repository"
)

// Patterns factory for repositories DBS
func main() {

	config := &configuration.Configuration{
		Engine: "sqlserver",
		Host:   "localhost",
		Port:   "3306",
		User:   "root",
		Pass:   "root",
		DBName: "patrones",
	}

	repo, err := repository.NewRepository(config)
	if err != nil {
		panic(err)
	}
	err = repo.Save("Guardar")
	if err != nil {
		panic(err)
	}
	valor := repo.Find(2)

	fmt.Println(valor)

}
