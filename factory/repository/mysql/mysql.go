package mysql

import "fmt"

type Mysql struct {
}

func NewRepository() *Mysql {
	return &Mysql{}
}

func (m *Mysql) Find(id int) string {
	return "Hello Word from mysql"
}

func (m *Mysql) Save(data string) error {
	fmt.Println("Save data to mysql")
	return nil
}
