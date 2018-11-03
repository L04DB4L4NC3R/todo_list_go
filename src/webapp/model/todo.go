package model

type todo struct {
	Done bool
	Task string
}

type Todolist struct {
	Name string
	Todo []todo
}

func Newdata() Todolist {
	return Todolist{Name: "angad", Todo: []todo{}}
}
