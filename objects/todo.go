pakage main

type Todo struct {
	Id int
	Description string
	Checked bool
	static todos []Todo = make([]Todo, 0)
}

func (t *Todo) Toggle() {
	t.Checked = !t.Checked
}

func (t *Todo) Update(description string) {
	t.Description = description
}

func (t *Todo) Delete() {
	for i, todo := range todos {
		if todo.Id == t.Id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}
}

func (t *Todo) Save() {
	todos = append(todos, *t)
}

func (t *Todo) All() []Todo {
	return todos
}

func (t *Todo) Find(id int) Todo {
	for _, todo := range todos {
		if todo.Id == id {
			return todo
		}
	}
	return Todo{}
}

func (t *Todo) Clear() {
	todos = make([]Todo, 0)
}

func (t *Todo) Count() int {
	return len(todos)
}

func (t *Todo) Complete() int {
	count := 0
	for _, todo := range todos {
		if todo.Checked {
			count++
		}
	}
	return count
}