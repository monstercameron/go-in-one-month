package objects

import (
	"strconv"
)

// Unique identifier for the todo
// Description of the todo
// Indicates whether the todo is checked or not
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

// Todos represents a collection of Todo items.
type Todos struct {
	Todos []*Todo
}

// NewTodos creates a new instance of Todos.
// It initializes the slice to hold pointers to Todo structs.
func NewTodos() *Todos {
	return &Todos{
		Todos: make([]*Todo, 0), // Initialize the slice to hold pointers to Todo structs
	}
}

// Add adds a new todo to the Todos slice.
// It appends the given todo to the existing Todos slice.
func (t *Todos) Add(todo *Todo) {
	t.Todos = append(t.Todos, todo)
}

// Remove removes the specified todo from the Todos slice.
// It searches for the todo with a matching Id and removes it from the slice.
// If no matching todo is found, the slice remains unchanged.
func (t *Todos) Remove(todo Todo) {
	for i, v := range t.Todos {
		if v.Id == todo.Id {
			t.Todos = append(t.Todos[:i], t.Todos[i+1:]...)
			break
		}
	}
}

// Get retrieves a Todo object from the Todos collection based on the provided id.
// If the id is not a valid integer or if no Todo with the matching id is found, it returns nil.
func (t *Todos) Get(id string) *Todo {
	num, err := strconv.Atoi(id)
	if err != nil {
		// handle the error appropriately
		return nil
	}

	for _, v := range t.Todos {
		if v.Id == num {
			return v
		}
	}
	return nil
}

// todos is a variable that holds a new instance of the Todos struct.
var TodoList = NewTodos()

// PopulateTodos populates the TodoList with some initial todo items.
// It adds three todo items to the TodoList:
// - Buy milk
// - Buy eggs
// - Buy bread
func PopulateTodos() {
	TodoList.Add(&Todo{Id: 1, Description: "Buy milk", Checked: false})
	TodoList.Add(&Todo{Id: 2, Description: "Buy eggs", Checked: true})
	TodoList.Add(&Todo{Id: 3, Description: "Buy bread", Checked: false})
}
