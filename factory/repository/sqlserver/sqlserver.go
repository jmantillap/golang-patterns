package sqlserver

import "fmt"

type SqlServer struct {
}

func NewRepository() *SqlServer {
	return &SqlServer{}
}

func (m *SqlServer) Find(id int) string {
	return "Hello Word from SqlServer"
}

func (m *SqlServer) Save(data string) error {
	fmt.Println("Save data to SqlServer")
	return nil
}
