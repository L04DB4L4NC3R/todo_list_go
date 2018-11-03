package model

type todo struct {
	Done bool
	Task string
	ID   string
}

type Todolist struct {
	Name string
	Todo []todo
}

func Newdata() Todolist {
	return Todolist{Name: "angad", Todo: getData()}
}

func getData() []todo {
	arr := []todo{}
	arr = append(arr, todo{Done: false, Task: "Make tea", ID: "1"})
	arr = append(arr, todo{Done: false, Task: "Eat a sandwich", ID: "2"})
	arr = append(arr, todo{Done: false, Task: "Feed dog", ID: "3"})
	arr = append(arr, todo{Done: true, Task: "Learn go", ID: "4"})
	return arr
}
