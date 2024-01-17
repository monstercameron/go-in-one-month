package objects

// Unique identifier for the todo item
// Description of the todo item
// Indicates whether the todo item is checked or not
type Todo struct {
	Id          int
	Description string
	Checked     bool
}

// Toggle toggles the checked status of the Todo.
func (t *Todo) Toggle() {
	t.Checked = !t.Checked
}

// Update updates the description of the Todo.
// It takes a string parameter 'description' and assigns it to the Todo's Description field.
func (t *Todo) Update(description string) {
	t.Description = description
}


type Todos struct {
	Todos []Todo
}

// NewTodos is a package-level function that returns a new instance of Todos.
func NewTodos() *Todos {
	return &Todos{
		Todos: make([]Todo, 0),
	}
}

func (t *Todos) Add(todo Todo) {
	t.Todos = append(t.Todos, todo)
}

func (t *Todos) Remove(todo Todo) {
	for i, v := range t.Todos {
		if v.Id == todo.Id {
			t.Todos = append(t.Todos[:i], t.Todos[i+1:]...)
			break
		}
	}
}
