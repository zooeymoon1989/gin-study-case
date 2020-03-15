package model

type TodoList struct {
	Id   int64  `gorm:"primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Done int8   `gorm:"column:done" json:"done"`
}

func (l TodoList) GetAll() ([]TodoList, error) {
	panic("Implement me!")
}

func (TodoList)TableName() string {
	return "todo_list"
}

type TodoListMethod interface {
	GetAll() ([]TodoList , error)
}