package entity

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todoItems []Todo

func GetTodos() []Todo {
	return todoItems
}

func GetTodoByID(id int) *Todo {
	for _, item := range todoItems {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func CreateTodo(todo Todo) {
	todoItems = append(todoItems, todo)
}

func UpdateTodo(id int, todo Todo) {
	for i, item := range todoItems {
		if item.ID == id {
			todoItems[i] = todo
			break
		}
	}
}

func DeleteTodo(id int) {
	for i, item := range todoItems {
		if item.ID == id {
			todoItems = append(todoItems[:i], todoItems[i+1:]...)
			return
		}
	}
}
