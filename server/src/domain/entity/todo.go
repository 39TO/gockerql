package entity

type Todo struct {
	Id     string
	Title  string
	Done   bool
	UserId int
}
type Todos []*Todo
