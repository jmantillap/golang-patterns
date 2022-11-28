package repository

import (
	"fmt"
	"pattern/factory/configuration"
	"pattern/factory/repository/mysql"
	"pattern/factory/repository/sqlserver"
)

type Repository interface {
	Find(id int) string
	Save(data string) error
}

func NewRepository(config *configuration.Configuration) (Repository, error) {

	var repo Repository
	var err error

	switch config.Engine {
	case "mysql":
		repo = mysql.NewRepository()
	case "sqlserver":
		repo = sqlserver.NewRepository()
	default:
		//err = errors.New("not supported engine")
		err = fmt.Errorf("unsupported engine: %s", config.Engine)
	}

	return repo, err
}
